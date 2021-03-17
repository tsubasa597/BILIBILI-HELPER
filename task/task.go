package task

// Tasker 任务
type Tasker interface {
	Run()
}

// Task task 类型的函数
type Task func(v ...string)

// Run task 类型的函数调用
func (t Task) Run(v ...string) {
	t(v...)
}

// Start 启动任务
func Start() {
	// todo := context.TODO()
	// context.WithDeadline(todo)
	for _, i := range New().task {
		go i.Run()
	}
}
