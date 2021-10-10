// Package task 完成固定操作
package task

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/state"
)

// Tasker 任务接口
type Tasker interface {
	Run(chan<- interface{})
	Next(time.Time) time.Time
}

// Entry 保存下一次和这次的运行时间
type Entry struct {
	Task Tasker
	prev time.Time
	next time.Time
}

// Corn 管理所有任务，由 chan 传递数据
type Corn struct {
	tasks   *sync.Map
	running int32
	once    *sync.Once
	Ch      chan interface{}
}

// New 初始化
func New() Corn {
	return Corn{
		tasks: &sync.Map{},
		once:  &sync.Once{},
		Ch:    make(chan interface{}),
	}
}

// Add 添加新任务
func (c Corn) Add(t Tasker) {
	if _, ok := c.tasks.Load(t); ok {
		return
	}

	ti := time.Now()
	c.tasks.Store(t, &Entry{
		prev: ti,
		next: t.Next(ti),
		Task: t,
	})
}

// Stop 停止
func (c Corn) Stop() {
	atomic.SwapInt32(&c.running, int32(state.Stop))
	close(c.Ch)
}

// Start 开始运行
func (c Corn) Start() {
	atomic.SwapInt32(&c.running, int32(state.Runing))

	c.once.Do(func() {
		go c.run()
	})
}

func (c Corn) run() {
	for atomic.LoadInt32(&c.running) == int32(state.Runing) {
		c.tasks.Range(func(key, value interface{}) bool {
			t := time.Now()
			entry := value.(*Entry)
			if entry.next.Before(t) {
				go entry.Task.Run(c.Ch)
				entry.prev = t
				entry.next = entry.Task.Next(t)
			}

			time.Sleep(time.Second) // 防止请求太快
			return true
		})
	}
}
