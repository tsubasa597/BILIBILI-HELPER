package info

const (
	BaseHost     = "https://api.bilibili.com"
	BaseLiveHost = "https://api.live.bilibili.com"
	BaseVCHost   = "https://api.vc.bilibili.com"
	// videoView    = "https://www.bilibili.com/video"
	// dynamicView  = "https://t.bilibili.com"

	// roomInit               = baseLiveHost + "/room/v1/Room/room_init"
	SpaceAccInfo           = BaseHost + "/x/space/acc/info"
	DynamicSrvSpaceHistory = BaseVCHost + "/dynamic_svr/v1/dynamic_svr/space_history"
	GetRoomInfoOld         = BaseLiveHost + "/room/v1/Room/getRoomInfoOld"
	// dynamicSrvDynamicNew   = baseVCHost + "/dynamic_svr/v1/dynamic_svr/dynamic_new"
	// relationModify         = baseHost + "/x/relation/modify"
	// relationFeedList       = baseLiveHost + "/relation/v1/feed/feed_list"
	// getAttentionList       = baseVCHost + "/feed/v1/feed/get_attention_list"
	UserLogin          = BaseHost + "/x/web-interface/nav"
	VideoHeartbeat     = BaseHost + "/x/click-interface/web/heartbeat"
	AvShare            = BaseHost + "/x/web-interface/share/add"
	Sliver2CoinsStatus = BaseLiveHost + "/xlive/revenue/v1/wallet/getStatus"
	Sliver2Coins       = BaseLiveHost + "/xlive/revenue/v1/wallet/silver2coin"
	LiveCheckin        = BaseLiveHost + "/xlive/web-ucenter/v1/sign/DoSign"
	RandomAV           = BaseHost + "/x/web-interface/search/default"
	Reply              = BaseHost + "/x/v2/reply"
)
