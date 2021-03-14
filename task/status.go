package task

import "bili/config"

// Dailyer 任务信息
type Dailyer interface {
	// LiveCheckin 直播签到
	DailyLiveCheckin()
	// UserCheck 用户检查
	UserCheck()
	// Sliver2Coins 银瓜子换硬币
	DailySliver2Coins()
	// DailyVideo 观看视频
	DailyVideo(string)
	// DailyVideo 分享视频
	DailyVideoShare()
}

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
	if status.IsLogin {
		if config.Conf.Status.IsLiveCheckin {
			status.DailyLiveCheckin()
		}
		if config.Conf.Status.IsSliver2Coins {
			status.DailySliver2Coin()
		}
		if config.Conf.Status.IsVideoWatch {
			status.DailyVideo("BV1NT4y137Jc")
		}
		if config.Conf.Status.IsVideoShare {
			status.DailyVideoShare("BV1UN411Q7k3")
		}
	}
}
