// Package task 完成固定操作
package task

import (
	"context"
	"sync"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/state"
)

// Tasker 任务接口
type Tasker interface {
	Run(chan<- interface{})
	Next(time.Time) time.Time
	State() state.State
}

// Entry 保存下一次和这次的运行时间
type Entry struct {
	Task Tasker
	prev time.Time
	next time.Time
}

// Corn 管理所有任务，由 chan 传递数据
type Corn struct {
	Ch     chan interface{}
	Ctx    context.Context
	cancel context.CancelFunc
	tasks  *sync.Map
	once   *sync.Once
	state  state.State
}

// New 初始化
func New(ctx context.Context) Corn {
	ctx, cancel := context.WithCancel(ctx)
	return Corn{
		Ch:     make(chan interface{}),
		Ctx:    ctx,
		cancel: cancel,
		tasks:  &sync.Map{},
		once:   &sync.Once{},
		state:  state.Runing,
	}
}

// Add 添加新任务
func (c Corn) Add(id int64, t Tasker) {
	if task, ok := c.tasks.Load(id); ok && task.(*Entry).Task.State() != state.Stop {
		return
	}

	ti := time.Now()
	c.tasks.Store(id, &Entry{
		prev: ti,
		next: ti,
		Task: t,
	})
}

// Stop 停止
func (c Corn) Stop() {
	c.cancel()
}

// Start 开始运行
func (c Corn) Start() {
	c.once.Do(func() {
		go c.run()
	})
}

func (c *Corn) run() {
	defer close(c.Ch)

	var (
		nextTime time.Time = time.Now().Local()
		now      time.Time
	)

	for {
		if state.Stop == c.state {
			return
		}

		c.tasks.Range(func(key, value interface{}) bool {
			select {
			case <-c.Ctx.Done():
				c.state = state.Stop
				return false
			case now = <-time.After(time.Until(nextTime)):
				if value.(*Entry).Task.State() == state.Stop {
					c.tasks.Delete(key)
					return true
				}

				entry := value.(*Entry)
				if entry.next.Before(now) {
					entry.Task.Run(c.Ch)

					next := entry.Task.Next(now)
					if next.Before(nextTime) {
						nextTime = next
					}

					entry.next = next
					entry.prev = now
				}
				return true
			}
		})
	}
}
