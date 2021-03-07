package task

import (
	"bili/config"
	"bili/utils"
)

// Tasker 任务
type Tasker interface {
	// LiveCheckin 直播签到
	liveCheckin()
	// UserCheck 用户检查
	userCheck()
	// Sliver2CoinsStatus 银瓜子换硬币状态
	sliver2CoinsStatus()
	// Sliver2Coins 银瓜子换硬币
	sliver2Coins()
	// VideoWatch 观看视频
	videoWatch(string)
	// videoShare 分享视频
	videoShare(string)
}

// Response 返回 json 的结构
type Response utils.JSON

func init() {
	status := &Status{}
	status.UserCheck()
	if status.IsLogin {
		if config.Conf.Status.IsLiveCheckin {
			status.DailyLiveCheckin()
		}
		if config.Conf.Status.IsSliver2Coins {
			status.DailySliver2Coin()
		}
		if config.Conf.Status.IsVideoWatch {
			status.DailyVideo()
		}
		if config.Conf.Status.IsVideoShare {
			status.DailyVideoShare()
		}
	}
}
