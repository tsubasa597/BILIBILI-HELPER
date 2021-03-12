package task

import "bili/config"

// Tasker 任务
type Tasker interface {
	Run()
}

// Task task 类型的函数
type Task func(string)

// Run task 类型的函数调用
func (t Task) Run(s string) {
	t(s)
}

// New 启动日常任务
func New() {
	status := &DailyInfo{}
	status.UserCheck()
	if status.IsLogin {
		if config.Conf.Status.IsLiveCheckin {
			status.DailyLiveCheckin()
		}
		if config.Conf.Status.IsSliver2Coins {
			status.DailySliver2Coin()
		}
		if config.Conf.Status.IsVideoWatch {
			Task(status.DailyVideo).Run("BV1NT4y137Jc")
		}
		if config.Conf.Status.IsVideoShare {
			Task(status.DailyVideoShare).Run("BV1NT4y137Jc")
		}
	}
}
