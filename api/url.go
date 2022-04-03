package api

const (
	baseHost     = "https://api.bilibili.com"
	baseLiveHost = "https://api.live.bilibili.com"
	baseVCHost   = "https://api.vc.bilibili.com"
	baseLive     = "https://live-trace.bilibili.com"

	// SpaceAccInfo 用户空间详细信息
	SpaceAccInfo = baseHost + "/x/space/acc/info"

	// DynamicSrvSpaceHistory 用户历史动态
	DynamicSrvSpaceHistory = baseVCHost + "/dynamic_svr/v1/dynamic_svr/space_history"

	// UserLogin 用户信息
	UserLogin = baseHost + "/x/web-interface/nav"

	// VideoHeartbeat 观看视频心跳
	VideoHeartbeat = baseHost + "/x/click-interface/web/heartbeat"

	// AvShare 分享视频
	AvShare = baseHost + "/x/web-interface/share/add"

	// Sliver2CoinsStatus 银瓜子状态
	Sliver2CoinsStatus = baseLiveHost + "/xlive/revenue/v1/wallet/getStatus"

	// Sliver2Coins 银瓜子换硬币
	Sliver2Coins = baseLiveHost + "/xlive/revenue/v1/wallet/silver2coin"

	// LiveCheckin 直播签到
	LiveCheckin = baseLiveHost + "/xlive/web-ucenter/v1/sign/DoSign"

	// RandomAV 视频全站默认排名
	RandomAV = baseHost + "/x/web-interface/search/default"

	// Reply 评论区
	Reply = baseHost + "/x/v2/reply"

	// EnterRoom 进入直播间
	EnterRoom = baseLive + "/xlive/data-interface/v1/x25Kn/E"

	// InRoom 直播间心跳
	InRoom = baseLive + "/xlive/data-interface/v1/x25Kn/X"

	// LiveInfo 直播间信息
	LiveInfo = baseLiveHost + "/room/v1/Room/get_info"

	// LiveInfo2 直播间信息
	LiveInfo2 = baseLiveHost + "xlive/web-room/v1/index/getInfoByRoom"
)
