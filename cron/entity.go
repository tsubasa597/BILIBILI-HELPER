package cron

import (
	"container/heap"
	"time"
)

// Entity 保存下一次和这次的运行时间
type Entity struct {
	ID   int64
	Task Tasker

	// Prev 待执行时间
	Prev time.Time

	// next 下一次执行时间
	next time.Time
}

// Done 调整执行时间
func (entity *Entity) Done() {
	entity.Prev = entity.next
	entity.next = entity.Task.Next(entity.Prev)
}

// Entities 保存所有 Entity 用于排序
type Entities struct {
	data []*Entity
	num  int
}

var (
	_ heap.Interface = (*Entities)(nil)
)

func (e Entities) Len() int {
	return e.num
}

func (e Entities) Less(i, j int) bool {
	return e.data[i].Prev.Before(e.data[j].Prev)
}

func (e Entities) Swap(i, j int) {
	e.data[i], e.data[j] = e.data[j], e.data[i]
}

func (e *Entities) Pop() interface{} {
	e.num--
	return e.data[e.Len()]
}

func (e *Entities) Push(v interface{}) {
	entity := v.(*Entity)
	entity.next = entity.Task.Next(entity.Prev)

	e.num++
	e.data = append(e.data, entity)
}

// Remove 从队列中删除
func (e *Entities) Remove() {
	e.data = append(e.data[:e.Len()], e.data[e.Len()+1:]...)
}

// Reset 重新排列队列
func (e *Entities) Reset() time.Time {
	if e.num == 0 {
		e.num = len(e.data)
	}

	heap.Init(e)

	if e.Len() == 0 {
		return time.Now().AddDate(15, 0, 0) // 等待添加任务
	}

	return e.data[e.Len()-1].Prev
}
