package api

import (
	"net/http"

	"github.com/tsubasa597/BILIBILI-HELPER/global"
)

const (
	BaseHost     = "https://api.bilibili.com"
	BaseLiveHost = "https://api.live.bilibili.com"
	BaseVCHost   = "https://api.vc.bilibili.com"
	VideoView    = "https://www.bilibili.com/video"
	DynamicView  = "https://t.bilibili.com"

	RoomInit               = BaseLiveHost + "/room/v1/Room/room_init"
	SpaceAccInfo           = BaseHost + "/x/space/acc/info"
	DynamicSrvSpaceHistory = BaseVCHost + "/dynamic_svr/v1/dynamic_svr/space_history"
	GetRoomInfoOld         = BaseLiveHost + "/room/v1/Room/getRoomInfoOld"
	DynamicSrvDynamicNew   = BaseVCHost + "/dynamic_svr/v1/dynamic_svr/dynamic_new"
	RelationModify         = BaseHost + "/x/relation/modify"
	RelationFeedList       = BaseLiveHost + "/relation/v1/feed/feed_list"
	GetAttentionList       = BaseVCHost + "/feed/v1/feed/get_attention_list"
	UserLogin              = BaseHost + "/x/web-interface/nav"
	VideoHeartbeat         = BaseHost + "/x/click-interface/web/heartbeat"
	AvShare                = BaseHost + "/x/web-interface/share/add"
	Sliver2CoinsStatus     = BaseLiveHost + "/pay/v1/Exchange/getStatus"
	Sliver2Coins           = BaseLiveHost + "/pay/v1/Exchange/silver2coin"
	LiveCheckin            = BaseLiveHost + "/xlive/web-ucenter/v1/sign/DoSign"
)

type API struct {
	conf global.Cookie
	r    *global.Requests
}

type Response struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	TTL     int                    `json:"ttl"`
	Data    map[string]interface{} `json:"data"`
}

func New(c global.Cookie) *API {
	r := global.NewRequests()
	r.SetHeader(http.Header{
		"Connection":   []string{"keep-alive"},
		"User-Agent":   []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.70"},
		"Cookie":       []string{c.GetVerify()},
		"Content-Type": []string{"application/x-www-form-urlencoded"},
	})
	return &API{
		r:    r,
		conf: c,
	}
}
