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
	Site         = "bilibili"
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

func (api API) GetDynamicMessage(hostUID int64) (s string, e error) {
	dynamicSvrSpaceHistoryResponse, err := api.GetDynamicSrvSpaceHistory(hostUID)
	if err != nil {
		return "", err
	}
	if dynamicSvrSpaceHistoryResponse.Code != 0 {
		return "", RequestErr{}
	}
	var dRToString func(int32, string)
	dRToString = func(i int32, c string) {
		switch i {
		case 1:

			dynamic := &CardWithOrig{}
			err = json.Unmarshal([]byte(c), dynamic)
			if err != nil {
				e = err
				return
			}
			dRToString(int32(dynamic.Item.OrigType), dynamic.Origin)

		case 2:
			dynamic := &CardWithImage{}
			e = json.Unmarshal([]byte(c), dynamic)
			s = dynamic.Item.Description
		case 4:
			dynamic := &CardTextOnly{}
			e = json.Unmarshal([]byte(c), dynamic)
			s = dynamic.Item.Content
		case 8:
			dynamic := &CardWithVideo{}
			e = json.Unmarshal([]byte(c), dynamic)
			s = dynamic.Desc
		case 64:
			dynamic := &CardWithPost{}
			e = json.Unmarshal([]byte(c), dynamic)
			s = dynamic.String()
		case 256:
			dynamic := &CardWithMusic{}
			e = json.Unmarshal([]byte(c), dynamic)
			s = dynamic.String()
		case 512:
			dynamic := &CardWithAnime{}
			e = json.Unmarshal([]byte(c), dynamic)
			s = dynamic.String()
		case 2048:
			dynamic := &CardWithSketch{}
			e = json.Unmarshal([]byte(c), dynamic)
			s = dynamic.Sketch.DescText
		case 4200:
			dynamic := &CardWithLive{}
			e = json.Unmarshal([]byte(c), dynamic)
			s = dynamic.String()
		case 4308:
			dynamic := &CardWithLiveV2{}
			e = json.Unmarshal([]byte(c), dynamic)
			s = dynamic.String()
		}
	}
	dRToString(int32(dynamicSvrSpaceHistoryResponse.Data.Cards[0].Desc.Type), dynamicSvrSpaceHistoryResponse.Data.Cards[0].Card)
	return
}

func (api API) GetDynamicSrvSpaceHistory(hostUID int64) (*DynamicSvrSpaceHistoryResponse, error) {
	rep, err := api.r.Get(fmt.Sprintf("%s?host_uid=%d", DynamicSrvSpaceHistory, hostUID))
	if err != nil {
		return nil, err
	}

	resp := &DynamicSvrSpaceHistoryResponse{}

	err = json.Unmarshal(rep, &resp)

	return resp, err
}

func (api API) UserCheck() (*Response, error) {
	rep, err := api.r.Get(UserLogin)
	if err != nil {
		return nil, err
	}
	resp := &Response{}
	err = json.Unmarshal(rep, &resp)
	return resp, err
}

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

func (api API) Sliver2CoinsStatus() (*Sliver2CoinsStatusResponse, error) {
	rep, err := api.r.Get(Sliver2CoinsStatus)
	if err != nil {
		return nil, err
	}
	resp := &Sliver2CoinsStatusResponse{}
	err = json.Unmarshal(rep, &resp)
	return resp, err
}

func (api API) Sliver2Coins() (*Response, error) {
	rep, err := api.r.Get(Sliver2Coins)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(rep, &resp)

	return resp, err
}

func (api API) LiveCheckin() (*Response, error) {
	rep, err := api.r.Get(LiveCheckin)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(rep, &resp)

	return resp, err
}

/**
// GiveGift 直播赠送礼物
func (api *API) GiveGift(param []string) {
	// info.api.sendGift("510", "7706705", info.api.Cookie.UserID, conf.Cookie.BiliJct)
}

// liveGetRecommend 随机获取一个直播间的 room_id
func (api API) liveGetRecommend() (float64, error) {
	res, err := api.r.Get(api.UrlList["LiveRecommend"])
	if err != nil {
		fmt.Println(err)
	}

	return res.Data["list"].([]interface{})[0].(map[string]interface{})["roomid"].(float64)
}

// liveGetRoomUID 通过直播间 roomid 获取主播 uid
func (api API) liveGetRoomUID(roomID string) float64 {
	res, err := api.requests.Get(api.UrlList["LiveGetRoomUID"] + "?room_id=" + roomID)
	if err != nil {
		fmt.Println(err)
	}
	return res.Data["room_info"].(map[string]interface{})["uid"].(float64)
}

// getRoomInfoOld 根据 uid 获取其 roomid
func (api API) getRoomInfoOld(uid string) float64 {
	res, err := api.requests.Get(api.UrlList["RoomInfoOld"] + "?mid=" + uid)
	if err != nil {
		fmt.Println(err)
	}
	return res.Data["roomid"].(float64)
}

// getGiftBagList 获取背包礼物
func (api API) getGiftBagList() []interface{} {
	res, err := api.requests.Get(api.UrlList["GiftBagList"])
	if err != nil {
		log.Fatalln(err)
	}
	return res.Data["list"].([]interface{})
}

// sendGift 送出礼物
func (api API) sendGift(roomID string, uid string) {
	giftBags := api.getGiftBagList()
	if len(giftBags) <= 0 {
		fmt.Println("背包里没有礼物")
	} else {
		gift := giftBags[0].(map[string]interface{})
		postBody := []byte("biz_id=" + roomID +
			"&ruid=" + uid +
			"&bag_id=" + fmt.Sprintf("%f", gift["bag_id"].(float64)) +
			"&gift_id=" + fmt.Sprintf("%f", gift["gift_id"].(float64)) +
			"&gift_num=" + fmt.Sprintf("%f", gift["gift_num"].(float64)) +
			"&uid=" + api.conf.Cookie.UserID +
			"&csrf=" + api.conf.Cookie.BiliJct +
			"&send_ruid=" + "0" +
			"&storm_beat_id=" + "0" +
			"&price=" + "0" +
			"&platform=" + "pc" +
			"&biz_code=" + "live")
		res, err := api.requests.Post(api.UrlList["GiftSend"], postBody)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}
}

func (api API) userCheck(params map[string]string) bool {
	response, err := api.requests.Get(api.UrlList["Login"])
	if err != nil {
		api.logInfo <- []interface{}{"Fatal", err}
	}
	if response.Code == 0 && response.Data["isLogin"].(bool) {
		api.logInfo <- []interface{}{"Info", "Cookies有效，登录成功"}
		return true
	} else {
		api.logInfo <- []interface{}{"Fatal", "Cookies可能失效了,请仔细检查Github Secrets中DEDEUSERID SESSDATA BILI_JCT三项的值是否正确、过期"}
		return false
	}
	//info.Coins = response.Data["money"].(float64)
	//info.Level = response.Data["level_info"].(map[string]interface{})["current_level"].(float64)
	//info.NextLevelExp = response.Data["level_info"].(map[string]interface{})["next_exp"].(float64) - response.Data["level_info"].(map[string]interface{})["current_exp"].(float64)
}

func (api API) watchVideo(params map[string]string) {
	postBody := []byte("bvid=" + params["bvid"] + "&played_time=" + strconv.Itoa(rand.Intn(90)))
	response, err := api.requests.Post(api.UrlList["VideoHeartbeat"], postBody)
	if err != nil {
		api.logInfo <- []interface{}{"Fatal", err}
	}
	if response.Code == 0 {
		api.logInfo <- []interface{}{"Info", "视频播放成功"}
	} else {
		api.logInfo <- []interface{}{"Warn", "视频播放失败,原因: " + response.Message}
	}
}

func (api API) shareVideo(params map[string]string) {
	// fmt.Println(params, api.conf)
	postBody := []byte("bvid=" + params["bvid"] + "&csrf=" + api.conf.Cookie.BiliJct)
	response, err := api.requests.Post(api.UrlList["AvShare"], postBody)
	if err != nil && response.Code != 0 {
		api.logInfo <- []interface{}{"Fatal", err}
	}
	if response.Code == 0 {
		api.logInfo <- []interface{}{"Info", "视频分享成功"}
	} else {
		api.logInfo <- []interface{}{"Warn", "视频分享失败,原因: " + response.Message}
	}
}

func (api API) sliver2Coins(params map[string]string) {
	// 银瓜子兑换硬币汇率
	var exchangeRate float64 = 700
	response, err := api.requests.Get(api.UrlList["Sliver2CoinsStatus"])
	if err != nil {
		api.logInfo <- []interface{}{"Fatal", err}
	}
	slivers := response.Data["silver"].(float64)
	coins := response.Data["coin"].(float64)
	if slivers < exchangeRate {
		api.logInfo <- []interface{}{"Error", "当前银瓜子余额为: ", slivers, "，不足700,不进行兑换"}
	} else {
		response, err = api.requests.Get(api.UrlList["Sliver2Coins"])
		if response.Code != 403 && err != nil {
			api.logInfo <- []interface{}{"Fatal", err}
		}
		if response.Code == 0 {
			api.logInfo <- []interface{}{"Info", "银瓜子兑换成功"}
			coins++
			slivers -= exchangeRate
			api.logInfo <- []interface{}{"Info", "当前银瓜子余额: ", slivers}
			api.logInfo <- []interface{}{"Info", "兑换银瓜子后硬币余额: ", coins}
		} else {
			api.logInfo <- []interface{}{"Warn", "银瓜子兑换硬币失败 原因是: " + response.Message}
		}
	}
}

func (api API) checkLive(params map[string]string) {
	response, err := api.requests.Get(api.UrlList["LiveCheckin"])
	if err != nil {
		api.logInfo <- []interface{}{"Fatal", err}

	}
	if response.Code == 0 {
		api.logInfo <- []interface{}{"Info", "直播签到成功，本次签到获得" + response.Data["text"].(string) + "," + response.Data["specialText"].(string)}
	} else {
		api.logInfo <- []interface{}{"Warn", "直播签到失败: " + response.Message}
	}
}
*/
func getUrlList() map[string]string {
	return map[string]string{
		"LiveCheckin":        "https://api.live.bilibili.com/xlive/web-ucenter/v1/sign/DoSign",
		"VideoHeartbeat":     "https://api.bilibili.com/x/click-interface/web/heartbeat",
		"Login":              "https://api.bilibili.com/x/web-interface/nav",
		"Sliver2CoinsStatus": "https://api.live.bilibili.com/pay/v1/Exchange/getStatus",
		"Sliver2Coins":       "https://api.live.bilibili.com/pay/v1/Exchange/silver2coin",
		"AvShare":            "https://api.bilibili.com/x/web-interface/share/add",
		"LiveRecommend":      "https://api.live.bilibili.com/relation/v1/AppWeb/getRecommendList",
		"LiveGetRoomUID":     "https://api.live.bilibili.com/xlive/web-room/v1/index/getInfoByRoom",
		"RoomInfoOld":        "http://api.live.bilibili.com/room/v1/Room/getRoomInfoOld",
		"GiftBagList":        "https://api.live.bilibili.com/xlive/web-room/v1/gift/bag_list",
		"GiftSend":           "https://api.live.bilibili.com/gift/v2/live/bag_send",
		"Dynamic":            "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/space_history",
	}
}

func New(c global.Cookie) *API {
	r := global.New()
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
