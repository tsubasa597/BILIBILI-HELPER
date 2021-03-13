package task

import "bili/api"

// Tasker 任务
type Tasker interface {
	Run()
}

// Task task 类型的函数
type Task func([]string)

// Run task 类型的函数调用
func (t Task) Run(s ...string) {
	t(s)
}

// Start 启动任务
func Start() {
	Task(api.GiveGift).Run("1", "2")
}
