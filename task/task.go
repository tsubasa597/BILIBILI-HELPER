package task

import (
	"sync"
)

// Tasker 任务
type Tasker interface {
	Run()
}

// Task task 类型的函数
type Task func(wg *sync.WaitGroup, v ...string)

// Run task 类型的函数调用
func (t Task) Run(wg *sync.WaitGroup, v ...string) {
	wg.Add(1)
	go t(wg, v...)
}

// Start 启动任务
func Start() {
	// Task(api.GiveGift).Run("1", "2")
}
