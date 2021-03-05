package task

// Tasker 任务
type Tasker interface {
	// LiveCheckin 直播签到
	LiveCheckin()
	// UserCheck 用户检查
	UserCheck()
	// Sliver2CoinsStatus 银瓜子换硬币状态
	Sliver2CoinsStatus()
	// Sliver2Coins 银瓜子换硬币
	Sliver2Coins()
	// VideoWatch 观看视频
	VideoWatch(string)
}
