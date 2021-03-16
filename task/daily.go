package task

import (
	"bili/config"
	"sync"
)

// DailyInfo 任务信息
type DailyInfo struct {
	Level        float64
	NextLevelExp float64
	IsLogin      bool
	Slivers      float64
	Coins        float64
	// Err          []error
}

// New 启动日常任务
func New() {
	status := &DailyInfo{}
	status.UserCheck()
	var wg sync.WaitGroup
	if status.IsLogin {
		if config.Conf.Status.IsLiveCheckin {
			Task(status.DailyLiveCheckin).Run(&wg, "")
		}
		if config.Conf.Status.IsSliver2Coins {
			Task(status.DailySliver2Coin).Run(&wg, "")
		}
		if config.Conf.Status.IsVideoWatch {
			Task(status.DailyVideo).Run(&wg, "BV1NT4y137Jc")
		}
		if config.Conf.Status.IsVideoShare {
			Task(status.DailyVideoShare).Run(&wg, "BV1UN411Q7k3")
		}
	}
	wg.Wait()
}
