package task

import (
	"sort"
	"sync"
	"time"
)

// Entry 保存下一次和这次的运行时间
type Entry struct {
	Task Tasker
	prev time.Time
	next time.Time
}

type Enties struct {
	enties []*Entry
	add    chan *Entry
	stop   chan struct{}
	wg     *sync.WaitGroup
}

var (
	_ sort.Interface = (*Enties)(nil)
)

func newEnties() Enties {
	return Enties{
		enties: make([]*Entry, 0),
		add:    make(chan *Entry, 1),
		stop:   make(chan struct{}),
		wg:     &sync.WaitGroup{},
	}
}

func (e Enties) Len() int {
	return len(e.enties)
}

func (e Enties) Less(i, j int) bool {
	return e.enties[i].prev.After(e.enties[j].prev)
}

func (e Enties) Swap(i, j int) {
	e.enties[i], e.enties[j] = e.enties[j], e.enties[i]
}

func (e *Enties) run(ch chan interface{}) {
	var (
		effective time.Time
		now       = time.Now().Local()
	)

	for {
		sort.Sort(e)
		if len(e.enties) > 0 {
			effective = e.enties[0].next
		} else {
			effective = now.AddDate(15, 0, 0) // 等待添加任务
		}

		select {
		case <-e.stop:
			e.wg.Wait()

			close(ch)
			return
		case now = <-time.After(time.Until(effective)):
			for _, entry := range e.enties {
				if entry.next != effective {
					break
				}

				entry.prev = now
				entry.next = entry.Task.Next(now)

				e.wg.Add(1 /** 确保协程退出 */)
				go entry.Task.Run(ch, e.wg)
			}
		case task := <-e.add:
			e.enties = append(e.enties, task)
		}
	}
}
