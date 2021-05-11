package api

import (
	fmt "fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"
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

	errGetDynamic    = "请求发生错误"
	errNoDynamic     = "该用户没有动态"
	errUnknowDynamic = "未知动态"
	errNotListen     = "该用户未监听"
	errRepeatListen  = "重复监听"
	errLoad          = "解析错误"
)

type API struct {
	conf     cookie
	Requests requests
	entry    *logrus.Entry
}

type Info struct {
	T       int32
	Err     error
	Content string
	Card    interface{}
	Name    string
	Live
}

type Live struct {
	LiveStatus  bool
	LiveRoomURL string
	LiveTitle   string
}

func New(c cookie, enrty *logrus.Entry) API {
	r := newRequests()
	r.setHeader(http.Header{
		"Connection":   []string{"keep-alive"},
		"User-Agent":   []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.70"},
		"Cookie":       []string{c.getVerify()},
		"Content-Type": []string{"application/x-www-form-urlencoded"},
	})

	if enrty == nil {
		enrty = logrus.NewEntry(newLog())
	}

	return API{
		Requests: r,
		conf:     c,
		entry:    enrty,
	}
}

func (api API) GetUserInfo(uid int64) (*XSpaceAccInfoResponse, error) {
	resp := &XSpaceAccInfoResponse{}
	err := api.Requests.gets(fmt.Sprintf("%s?mid=%d", SpaceAccInfo, uid), resp)
	return resp, err
}

// UserCheck 用户登录验证
func (api API) UserCheck() (*BaseResponse, error) {
	resp := &BaseResponse{}
	err := api.Requests.gets(UserLogin, resp)

	return resp, err
}

// WatchVideo 视频模拟观看，观看时间在 [0, 90) 之间
func (api API) WatchVideo(bvid string) (*BaseResponse, error) {
	data := url.Values{
		"bvid":        []string{bvid},
		"played_time": []string{strconv.Itoa(rand.Intn(90))},
	}

	resp := &BaseResponse{}
	err := api.Requests.posts(VideoHeartbeat, data, resp)

	return resp, err
}

// ShareVideo 分享视频
func (api API) ShareVideo(bvid string) (*BaseResponse, error) {
	data := url.Values{
		"bvid": []string{bvid},
		"csrf": []string{api.conf.BiliJct},
	}

	resp := &BaseResponse{}
	err := api.Requests.posts(AvShare, data, resp)

	return resp, err
}

// Sliver2CoinsStatus 获取银瓜子和硬币的数量
func (api API) Sliver2CoinsStatus() (*Sliver2CoinsStatusResponse, error) {
	resp := &Sliver2CoinsStatusResponse{}
	err := api.Requests.gets(Sliver2CoinsStatus, resp)

	return resp, err
}

// Sliver2Coins 将银瓜子兑换为硬币
func (api API) Sliver2Coins() (*BaseResponse, error) {
	resp := &BaseResponse{}
	err := api.Requests.gets(Sliver2Coins, resp)

	return resp, err
}

// LiveCheckin 直播签到
func (api API) LiveCheckin() (*BaseResponse, error) {
	resp := &BaseResponse{}
	err := api.Requests.gets(LiveCheckin, resp)

	return resp, err
}

// GetDynamicSrvSpaceHistory 获取目的 uid 的所有动态
func (api API) GetDynamicSrvSpaceHistory(hostUID int64) (*DynamicSvrSpaceHistoryResponse, error) {
	resp := &DynamicSvrSpaceHistoryResponse{}
	err := api.Requests.gets(fmt.Sprintf("%s?host_uid=%d", DynamicSrvSpaceHistory, hostUID), resp)

	return resp, err
}

func (api API) LiverStatus(uid int64) (*GetRoomInfoOldResponse, error) {
	resp := &GetRoomInfoOldResponse{}
	err := api.Requests.gets(fmt.Sprintf("%s?mid=%d", GetRoomInfoOld, uid), resp)

	return resp, err
}
