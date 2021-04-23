package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"

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

// GetDynamicMessage 获取目标 uid 的第一条记录
func (api API) GetDynamicMessage(hostUID int64) (s string, e error) {
	dynamicSvrSpaceHistoryResponse, err := api.GetDynamicSrvSpaceHistory(hostUID)
	if err != nil {
		return "", err
	}

	if dynamicSvrSpaceHistoryResponse.Code != 0 {
		return "", RequestErr{}
	}

	return GetOriginCard(dynamicSvrSpaceHistoryResponse.Data.Cards[0])
}

// GetOriginCard 获取 Card 的源动态
func GetOriginCard(c *Card) (string, error) {
	switch c.Desc.Type {
	case 0:
		return "", DynamicUnknownErr{}
	case 1:
		dynamic := &CardWithOrig{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			return "", err
		}

		return GetOriginCard(&Card{
			Desc: &Card_Desc{
				Type: dynamic.Item.OrigType,
			},
			Card: dynamic.Origin,
		})
	case 2:
		dynamic := &CardWithImage{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			return "", err
		}

		return dynamic.Item.Description, nil
	case 4:
		dynamic := &CardTextOnly{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			return "", err
		}

		return dynamic.Item.Content, nil
	case 8:
		dynamic := &CardWithVideo{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			return "", err
		}

		return dynamic.Title, nil
	case 64:
		dynamic := &CardWithPost{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			return "", err
		}

		return dynamic.String(), nil
	case 256:
		dynamic := &CardWithMusic{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			return "", err
		}

		return dynamic.String(), nil
	case 512:
		dynamic := &CardWithAnime{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			return "", err
		}

		return dynamic.String(), nil
	case 2048:
		dynamic := &CardWithSketch{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			return "", err
		}

		return dynamic.Sketch.DescText, nil
	case 4200:
		dynamic := &CardWithLive{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			return "", err
		}

		return dynamic.String(), nil
	case 4308:
		dynamic := &CardWithLiveV2{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			return "", err
		}

		return dynamic.String(), nil
	}

	return "", LoadErr{}
}

// GetDynamicSrvSpaceHistory 获取目的 uid 的所有动态
func (api API) GetDynamicSrvSpaceHistory(hostUID int64) (*DynamicSvrSpaceHistoryResponse, error) {
	rep, err := api.r.Get(fmt.Sprintf("%s?host_uid=%d", DynamicSrvSpaceHistory, hostUID))
	if err != nil {
		return nil, err
	}

	resp := &DynamicSvrSpaceHistoryResponse{}

	err = json.Unmarshal(rep, &resp)

	return resp, err
}

// UserCheck 用户登录验证
func (api API) UserCheck() (*Response, error) {
	rep, err := api.r.Get(UserLogin)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(rep, &resp)

	return resp, err
}

// WatchVideo 视频模拟观看，观看时间在 [0, 90) 之间
func (api API) WatchVideo(bvid string) (*Response, error) {
	data := url.Values{
		"bvid":        []string{bvid},
		"played_time": []string{strconv.Itoa(rand.Intn(90))},
	}

	rep, err := api.r.Post(VideoHeartbeat, data)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(rep, &resp)

	return resp, err
}

// ShareVideo 分享视频
func (api API) ShareVideo(bvid string) (*Response, error) {
	data := url.Values{
		"bvid": []string{bvid},
		"csrf": []string{api.conf.BiliJct},
	}
	rep, err := api.r.Post(AvShare, data)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(rep, &resp)

	return resp, err
}

// Sliver2CoinsStatus 获取银瓜子和硬币的数量
func (api API) Sliver2CoinsStatus() (*Sliver2CoinsStatusResponse, error) {
	rep, err := api.r.Get(Sliver2CoinsStatus)
	if err != nil {
		return nil, err
	}

	resp := &Sliver2CoinsStatusResponse{}
	err = json.Unmarshal(rep, &resp)

	return resp, err
}

// Sliver2Coins 将银瓜子兑换为硬币
func (api API) Sliver2Coins() (*Response, error) {
	rep, err := api.r.Get(Sliver2Coins)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(rep, &resp)

	return resp, err
}

// LiveCheckin 直播签到
func (api API) LiveCheckin() (*Response, error) {
	rep, err := api.r.Get(LiveCheckin)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(rep, &resp)

	return resp, err
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
