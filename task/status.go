package task

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
