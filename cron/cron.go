// Package cron 完成固定操作
package cron

import (
	"sort"
	"sync"

	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"go.uber.org/atomic"
)

// Cron 管理所有任务，由 chan 传递数据
// Enties([]Entity) <= Entity <= Tasker
type Cron struct {
	Ch     chan interface{}
	state  *atomic.Int32
	enties Enties
	ids    *sync.Map // id (int64) => Tasker
	add    chan *Entry
	stop   chan struct{}
	wg     *sync.WaitGroup
}

// New 初始化
func New() Cron {
	return Cron{
		Ch:     make(chan interface{}, 1),
		state:  atomic.NewInt32(info.Stop),
		enties: make(Enties, 0),
		ids:    &sync.Map{},
		add:    make(chan *Entry, 1),
		stop:   make(chan struct{}),
		wg:     &sync.WaitGroup{},
	}
}

// Add 添加新任务
func (c *Cron) Add(id int64, t Tasker) {
	ti := time.Now()
	entry := &Entry{
		Task: t,
		id:   id,
		prev: ti,
		next: ti,
	}

	if _, ok := c.ids.Load(id); ok {
		return
	}

	c.ids.Store(id, t)

	if c.state.Load() == info.Running {
		c.add <- entry
		return
	}

	c.enties = append(c.enties, entry)
}

// Stop 停止监听
// channle 中所有信息被读取之后返回
func (c Cron) Stop() {
	c.state.Swap(info.Stop)
	c.stop <- struct{}{}
}

// Start 开始运行
// 调用 Stop 后请不要调用 Start，否则会触发 panic
func (c Cron) Start() {
	if c.state.CAS(info.Stop, info.Running) {
		go c.run(c.Ch)
	}
}

// StopByID 停止指定 id 任务的执行
func (c Cron) StopByID(id int64) {
	c.ids.Delete(id)
}

// Info 返回所有任务的信息
func (c Cron) Info() []info.Task {
	if c.state.Load() == info.Stop {
		return []info.Task{}
	}

	taskInfos := make([]info.Task, 0, len(c.enties))
	c.ids.Range(func(key, value interface{}) bool {
		taskInfos = append(taskInfos, value.(Tasker).Info())
		return true
	})

	return taskInfos
}

func (cron *Cron) run(ch chan interface{}) {
	var (
		effective time.Time
		now       = time.Now().Local()
	)

	for {
		sort.Sort(cron.enties)
		if len(cron.enties) > 0 {
			effective = cron.enties[0].next
		} else {
			effective = now.AddDate(15, 0, 0) // 等待添加任务
		}

		select {
		case <-cron.stop:
			cron.wg.Wait()

			close(ch)
			return
		case now = <-time.After(time.Until(effective)):
			for i, entry := range cron.enties {
				if entry.next != effective {
					break
				}

				if _, ok := cron.ids.Load(entry.id); !ok {
					cron.enties = append(cron.enties[:i], cron.enties[i+1:]...)
					continue
				}

				entry.prev = now
				entry.next = entry.Task.Next(now)

				cron.wg.Add(1 /** 确保协程退出 */)
				go entry.Task.Run(ch, cron.wg)
			}
		case task := <-cron.add:
			cron.enties = append(cron.enties, task)
		}
	}
}
