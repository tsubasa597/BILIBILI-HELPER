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
	Ctx    context.Context
	cancel context.CancelFunc
	tasks  *sync.Map
	once   *sync.Once
	Ch     chan interface{}
}

// New 初始化
func New(ctx context.Context) Corn {
	ctx, cancel := context.WithCancel(ctx)
	return Corn{
		Ctx:    ctx,
		cancel: cancel,
		tasks:  &sync.Map{},
		once:   &sync.Once{},
		Ch:     make(chan interface{}),
	}
}

// Add 添加新任务 指针！
func (c Corn) Add(id int64, t Tasker) {
	if task, ok := c.tasks.Load(id); ok && task.(*Entry).Task.State() == state.Runing {
		return
	}

	c.tasks.Delete(id)

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

func (c Corn) run() {
	defer close(c.Ch)

	for {
		select {
		case <-c.Ctx.Done():
			return
		default:
			c.tasks.Range(func(key, value interface{}) bool {
				if value.(*Entry).Task.State() == state.Stop {
					c.tasks.Delete(key)
				}

				t := time.Now()
				entry := value.(*Entry)
				if entry.next.Before(t) {
					entry.Task.Run(c.Ch)
					entry.prev = t
					entry.next = entry.Task.Next(t)
				}

				return true
			})
		}
	}
}
