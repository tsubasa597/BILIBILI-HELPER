package task

import (
	"bili/config"
)

// DailyInfo 任务信息
type DailyInfo struct {
	conf         config.Config
	task         []Task
	Level        float64
	NextLevelExp float64
	IsLogin      bool
	Slivers      float64
	Coins        float64
	// Err          []error
}

// New 启动日常任务
func New() (status *DailyInfo) {
	status.UserCheck()
	if status.IsLogin {
		if status.conf.Status.IsLiveCheckin {
			status.task = append(status.task, Task(status.DailyLiveCheckin))
		}
		if status.conf.Status.IsSliver2Coins {
			status.task = append(status.task, Task(status.DailySliver2Coin))
		}
		if status.conf.Status.IsVideoWatch {
			status.task = append(status.task, Task(status.DailyVideo))
		}
		if status.conf.Status.IsVideoShare {
			status.task = append(status.task, Task(status.DailyVideoShare))
		}
	}
	return status
}
