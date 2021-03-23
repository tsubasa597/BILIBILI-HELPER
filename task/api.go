package task

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

var (
	apiList map[string]string = map[string]string{
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
	}
)

// GiveGift 直播赠送礼物
func (info *Info) GiveGift(param []string) {
	sendGift("510", "7706705", conf.Cookie.UserID, conf.Cookie.BiliJct)
}

// liveGetRecommend 随机获取一个直播间的 room_id
func liveGetRecommend() float64 {
	res, err := Get(apiList["LiveRecommend"])
	if err != nil {
		fmt.Println(err)
	}
	return res.Data["list"].([]interface{})[0].(map[string]interface{})["roomid"].(float64)
}

// liveGetRoomUID 通过直播间 roomid 获取主播 uid
func liveGetRoomUID(roomID string) float64 {
	res, err := Get(apiList["LiveGetRoomUID"] + "?room_id=" + roomID)
	if err != nil {
		fmt.Println(err)
	}
	return res.Data["room_info"].(map[string]interface{})["uid"].(float64)
}

// getRoomInfoOld 根据 uid 获取其 roomid
func getRoomInfoOld(uid string) float64 {
	res, err := Get(apiList["RoomInfoOld"] + "?mid=" + uid)
	if err != nil {
		fmt.Println(err)
	}
	return res.Data["roomid"].(float64)
}

// getGiftBagList 获取背包礼物
func getGiftBagList() []interface{} {
	res, err := Get(apiList["GiftBagList"])
	if err != nil {
		log.Fatalln(err)
	}
	return res.Data["list"].([]interface{})
}

// sendGift 送出礼物
func sendGift(roomID string, uid string, userId string, biliJct string) {
	giftBags := getGiftBagList()
	if len(giftBags) <= 0 {
		fmt.Println("背包里没有礼物")
	} else {
		gift := giftBags[0].(map[string]interface{})
		postBody := []byte("biz_id=" + roomID +
			"&ruid=" + uid +
			"&bag_id=" + fmt.Sprintf("%f", gift["bag_id"].(float64)) +
			"&gift_id=" + fmt.Sprintf("%f", gift["gift_id"].(float64)) +
			"&gift_num=" + fmt.Sprintf("%f", gift["gift_num"].(float64)) +
			"&uid=" + userId +
			"&csrf=" + biliJct +
			"&send_ruid=" + "0" +
			"&storm_beat_id=" + "0" +
			"&price=" + "0" +
			"&platform=" + "pc" +
			"&biz_code=" + "live")
		res, err := Post(apiList["GiftSend"], postBody)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}
}

func userCheck(logInfo chan []interface{}, params map[string]string) bool {
	response, err := Get(apiList["Login"])
	if err != nil {
		logInfo <- []interface{}{"Fatal", err}
	}
	if response.Code == 0 && response.Data["isLogin"].(bool) {
		logInfo <- []interface{}{"Info", "Cookies有效，登录成功"}
		return true
	} else {
		logInfo <- []interface{}{"Fatal", "Cookies可能失效了,请仔细检查Github Secrets中DEDEUSERID SESSDATA BILI_JCT三项的值是否正确、过期"}
		return false
	}
	//info.Coins = response.Data["money"].(float64)
	//info.Level = response.Data["level_info"].(map[string]interface{})["current_level"].(float64)
	//info.NextLevelExp = response.Data["level_info"].(map[string]interface{})["next_exp"].(float64) - response.Data["level_info"].(map[string]interface{})["current_exp"].(float64)
}

func watchVideo(logInfo chan []interface{}, params map[string]string) {
	postBody := []byte("bvid=" + params["bvid"] + "&played_time=" + strconv.Itoa(rand.Intn(90)))
	response, err := Post(apiList["VideoHeartbeat"], postBody)
	if err != nil {
		logInfo <- []interface{}{"Fatal", err}
	}
	if response.Code == 0 {
		logInfo <- []interface{}{"Info", "视频播放成功"}
	} else {
		logInfo <- []interface{}{"Warn", "视频播放失败,原因: " + response.Message}
	}
}

func shareVideo(logInfo chan []interface{}, params map[string]string) {
	postBody := []byte("bvid=" + params["bvid"] + "&csrf=" + conf.Cookie.BiliJct)
	response, err := Post(apiList["AvShare"], postBody)
	if err != nil && response.Code != 0 {
		logInfo <- []interface{}{"Fatal", err}
	}
	if response.Code == 0 {
		logInfo <- []interface{}{"Info", "视频分享成功"}
	} else {
		logInfo <- []interface{}{"Warn", "视频分享失败,原因: " + response.Message}
	}
}

func sliver2Coins(logInfo chan []interface{}, params map[string]string) {
	// 银瓜子兑换硬币汇率
	var exchangeRate float64 = 700
	response, err := Get(apiList["Sliver2CoinsStatus"])
	if err != nil {
		logInfo <- []interface{}{"Fatal", err}
	}
	slivers := response.Data["silver"].(float64)
	coins := response.Data["coin"].(float64)
	if slivers < exchangeRate {
		logInfo <- []interface{}{"Error", "当前银瓜子余额为: ", slivers, "，不足700,不进行兑换"}
	} else {
		response, err = Get(apiList["Sliver2Coins"])
		if response.Code != 403 && err != nil {
			logInfo <- []interface{}{"Fatal", err}
		}
		if response.Code == 0 {
			logInfo <- []interface{}{"Info", "银瓜子兑换成功"}
			coins++
			slivers -= exchangeRate
			logInfo <- []interface{}{"Info", "当前银瓜子余额: ", slivers}
			logInfo <- []interface{}{"Info", "兑换银瓜子后硬币余额: ", coins}
		} else {
			logInfo <- []interface{}{"Warn", "银瓜子兑换硬币失败 原因是: " + response.Message}
		}
	}
}

func checkLive(logInfo chan []interface{}, params map[string]string) {
	response, err := Get(apiList["LiveCheckin"])
	if err != nil {
		logInfo <- []interface{}{"Fatal", err}

	}
	if response.Code == 0 {
		logInfo <- []interface{}{"Info", "直播签到成功，本次签到获得" + response.Data["text"].(string) + "," + response.Data["specialText"].(string)}
	} else {
		logInfo <- []interface{}{"Warn", "直播签到失败: " + response.Message}
	}
}
