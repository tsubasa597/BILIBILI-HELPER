package cron

import (
	"sort"
	"time"
)

// Entry 保存下一次和这次的运行时间
type Entry struct {
	Task Tasker
	id   int64
	prev time.Time
	next time.Time
}

// Enties 保存所有 Entity 用于排序
type Enties []*Entry

var (
	_ sort.Interface = (*Enties)(nil)
)

func (e Enties) Len() int {
	return len(e)
}

func (e Enties) Less(i, j int) bool {
	return e[i].prev.After(e[j].prev)
}

func (e Enties) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
