// Package task 完成固定操作
package task

import (
	"context"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/state"
)

// Tasker 任务接口
type Tasker interface {
	Run(chan<- interface{}, *sync.WaitGroup)
	Next(time.Time) time.Time
}

// Corn 管理所有任务，由 chan 传递数据
type Corn struct {
	Ch    chan interface{}
	wg    *sync.WaitGroup
	tasks Enties
	ids   map[int64]struct{}
	state state.State
	add   chan *Entry
	stop  chan struct{}
}

// New 初始化
func New(ctx context.Context) Corn {
	return Corn{
		Ch:    make(chan interface{}, 1),
		wg:    &sync.WaitGroup{},
		tasks: make(Enties, 0),
		ids:   make(map[int64]struct{}),
		state: state.Stop,
		add:   make(chan *Entry, 1),
		stop:  make(chan struct{}),
	}
}

// Add 添加新任务
func (c *Corn) Add(id int64, t Tasker) {
	ti := time.Now()
	entry := &Entry{
		prev: ti,
		next: ti,
		Task: t,
	}

	if _, ok := c.ids[id]; ok {
		return
	}

	c.ids[id] = struct{}{}

	if c.state == state.Runing {
		c.add <- entry
		return
	}

	c.tasks = append(c.tasks, entry)
}

// Stop 停止
func (c Corn) Stop() {
	atomic.SwapInt32((*int32)(&c.state), int32(state.Stop))
	c.stop <- struct{}{}
}

// Start 开始运行
func (c Corn) Start() {
	ok := atomic.CompareAndSwapInt32((*int32)(&c.state), int32(state.Stop), int32(state.Runing))
	if ok {
		go c.run()
	}
}

func (c *Corn) run() {
	var (
		effective time.Time
		now       = time.Now().Local()
	)

	for {
		sort.Sort(c.tasks)
		if len(c.tasks) > 0 {
			effective = c.tasks[0].next
		} else {
			effective = now.AddDate(15, 0, 0) // 等待添加任务
		}

		select {
		case <-c.stop:
			c.wg.Wait()

			close(c.Ch)
			return
		case now = <-time.After(time.Until(effective)):
			for _, entry := range c.tasks {
				if entry.next != effective {
					break
				}

				entry.prev = now
				entry.next = entry.Task.Next(now)

				c.wg.Add(1 /** 确保协程退出 */)
				go entry.Task.Run(c.Ch, c.wg)
			}
		case task := <-c.add:
			task.next = task.Task.Next(now)
			c.tasks = append(c.tasks, task)
		}
	}
}
