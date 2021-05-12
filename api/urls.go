package api

const (
	baseHost     = "https://api.bilibili.com"
	baseLiveHost = "https://api.live.bilibili.com"
	baseVCHost   = "https://api.vc.bilibili.com"
	videoView    = "https://www.bilibili.com/video"
	dynamicView  = "https://t.bilibili.com"

	roomInit               = baseLiveHost + "/room/v1/Room/room_init"
	spaceAccInfo           = baseHost + "/x/space/acc/info"
	dynamicSrvSpaceHistory = baseVCHost + "/dynamic_svr/v1/dynamic_svr/space_history"
	getRoomInfoOld         = baseLiveHost + "/room/v1/Room/getRoomInfoOld"
	dynamicSrvDynamicNew   = baseVCHost + "/dynamic_svr/v1/dynamic_svr/dynamic_new"
	relationModify         = baseHost + "/x/relation/modify"
	relationFeedList       = baseLiveHost + "/relation/v1/feed/feed_list"
	getAttentionList       = baseVCHost + "/feed/v1/feed/get_attention_list"
	userLogin              = baseHost + "/x/web-interface/nav"
	videoHeartbeat         = baseHost + "/x/click-interface/web/heartbeat"
	avShare                = baseHost + "/x/web-interface/share/add"
	sliver2CoinsStatus     = baseLiveHost + "/pay/v1/Exchange/getStatus"
	sliver2Coins           = baseLiveHost + "/pay/v1/Exchange/silver2coin"
	liveCheckin            = baseLiveHost + "/xlive/web-ucenter/v1/sign/DoSign"
)
