// Package task 完成固定操作
package task

import (
	"sync"

	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"github.com/tsubasa597/BILIBILI-HELPER/state"
	"go.uber.org/atomic"
)

// Tasker 任务接口
type Tasker interface {
	Run(chan<- interface{}, *sync.WaitGroup)
	Next(time.Time) time.Time
	Info() info.Task
}

// Corn 管理所有任务，由 chan 传递数据
type Corn struct {
	Ch    chan interface{}
	tasks Enties
	ids   *sync.Map
	state *atomic.Int32
}

// New 初始化
func New() Corn {
	return Corn{
		Ch:    make(chan interface{}, 1),
		tasks: newEnties(),
		ids:   &sync.Map{},
		state: atomic.NewInt32(state.Stop),
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

	if _, ok := c.ids.Load(id); ok {
		return
	}

	c.ids.Store(id, t)

	if c.state.Load() == state.Running {
		c.tasks.add <- entry
		return
	}

	c.tasks.enties = append(c.tasks.enties, entry)
}

// Stop 停止监听
// channle 中所有信息被读取之后返回
func (c Corn) Stop() {
	c.state.Swap(state.Stop)
	c.tasks.stop <- struct{}{}
}

// Start 开始运行
// 调用 Stop 后请不要调用 Start，否则会触发 panic
func (c Corn) Start() {
	if c.state.CAS(state.Stop, state.Running) {
		go c.tasks.run(c.Ch)
	}
}

// Info 返回所有任务的信息
func (c Corn) Info() []info.Task {
	if c.state.Load() == state.Stop {
		return []info.Task{}
	}

	taskInfos := make([]info.Task, 0, len(c.tasks.enties))
	c.ids.Range(func(key, value interface{}) bool {
		taskInfos = append(taskInfos, value.(Tasker).Info())
		return true
	})

	return taskInfos
}
