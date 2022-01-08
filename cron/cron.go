// Package cron 完成固定操作
package cron

import (
	"container/heap"
	"sync"

	"time"

	"go.uber.org/atomic"
)

const (
	// Stop 停止状态
	Stop = iota
	// Pause 暂停状态
	Pause
	// Running 正在运行状态
	Running
)

// Cron 管理所有任务，由 chan 传递数据
// Enties([]Entity) <= Entity <= Tasker
type Cron struct {
	Ch       chan interface{}
	state    *atomic.Int32
	entities Entities
	ids      *sync.Map // id (int64) => Tasker
	add      chan *Entity
	stop     chan struct{}
	wg       *sync.WaitGroup
}

// New 初始化
func New() Cron {
	return Cron{
		Ch:       make(chan interface{}, 1),
		state:    atomic.NewInt32(Stop),
		entities: Entities{},
		ids:      &sync.Map{},
		add:      make(chan *Entity, 1),
		stop:     make(chan struct{}),
		wg:       &sync.WaitGroup{},
	}
}

// Add 添加新任务
func (c *Cron) Add(id int64, task Tasker, startTime time.Time) {
	entity := &Entity{
		ID:   id,
		Task: task,
		Prev: startTime,
	}

	if _, ok := c.ids.Load(entity.ID); ok {
		return
	}

	c.ids.Store(entity.ID, entity.Task)

	if c.state.Load() == Running {
		c.add <- entity
		return
	}

	heap.Push(&c.entities, entity)
}

// Stop 停止监听
// channle 中所有信息被读取之后返回
func (c Cron) Stop() {
	c.state.Swap(Stop)
	c.stop <- struct{}{}
}

// Start 开始运行
// 调用 Stop 后请不要调用 Start，否则会触发 panic
func (c Cron) Start() {
	if c.state.CAS(Stop, Running) {
		go c.run()
	}
}

// StopByID 停止指定 id 任务的执行
func (c Cron) StopByID(id int64) {
	c.ids.Delete(id)
}

// Info 返回所有任务的信息
func (c Cron) Info() []Info {
	if c.state.Load() == Stop {
		return []Info{}
	}

	taskInfos := make([]Info, 0)
	c.ids.Range(func(key, value interface{}) bool {
		taskInfos = append(taskInfos, value.(Tasker).Info())
		return true
	})

	return taskInfos
}

func (cron *Cron) run() {
	var (
		effective time.Time
	)

	for {
		effective = cron.entities.Reset()

		select {
		case <-cron.stop:
			cron.wg.Wait()

			close(cron.Ch)
			return
		case <-time.After(time.Until(effective)):
			for cron.entities.Len() > 0 {
				entity := heap.Pop(&cron.entities).(*Entity)

				if _, ok := cron.ids.Load(entity.ID); !ok {
					cron.entities.Remove()
					continue
				}

				if entity.Prev.After(effective) {
					break
				}

				cron.wg.Add(1)
				go entity.Task.Run(cron.Ch, cron.wg)

				entity.Done()
			}
		case entity := <-cron.add:
			heap.Push(&cron.entities, entity)
		}
	}
}
