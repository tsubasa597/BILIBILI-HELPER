package task

// Info 任务信息
type Info interface {
	// LiveCheckin 直播签到
	LiveCheckinInfo()
	// UserCheck 用户检查
	UserCheckInfo()
	// Sliver2CoinsStatus 银瓜子换硬币状态
	// Sliver2Coins 银瓜子换硬币
	Sliver2CoinsInfo()
	// VideoWatch 观看视频
	// VideoWatchInfo(string)
}

// Status 任务信息
type Status struct {
	Level                float64
	NextLevelExp         float64
	IsLogin              bool
	IsVideoWatch         bool
	IsLiveCheckin        bool
	IsSliver2CoinsStatus bool
	Slivers              float64
	Coins                float64
}
