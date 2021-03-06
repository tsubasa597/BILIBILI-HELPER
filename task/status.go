package task

// Daily 任务信息
type Daily interface {
	// LiveCheckin 直播签到
	DailyLiveCheckin()
	// UserCheck 用户检查
	UserCheck()
	// Sliver2Coins 银瓜子换硬币
	DailySliver2Coins()
	// DailyVideo 观看视频
	DailyVideo(string)
}

// Status 任务信息
type Status struct {
	Level                float64
	NextLevelExp         float64
	IsLogin              bool
	IsVideoWatch         bool
	IsVideoShare         bool
	IsLiveCheckin        bool
	IsSliver2CoinsStatus bool
	Slivers              float64
	Coins                float64
}
