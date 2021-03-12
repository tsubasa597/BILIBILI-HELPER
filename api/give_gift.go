package api

import (
	"bili/apiquery"
	"bili/config"
	"bili/utils"
	"fmt"
	"log"
)

// GiveGift 直播赠送礼物
func (js JSON) GiveGift() {
	js.sendGift("510", "7706705")
}

// liveGetRecommend 随机获取一个直播间的 room_id
func (js JSON) liveGetRecommend() float64 {
	res, err := utils.Get(apiquery.ApiList.LiveRecommend)
	if err != nil {
		fmt.Println(err)
	}
	return res.Data["list"].([]interface{})[0].(map[string]interface{})["roomid"].(float64)
}

// liveGetRoomUID 通过直播间 roomid 获取主播 uid
func (js JSON) liveGetRoomUID(roomID string) float64 {
	res, err := utils.Get(apiquery.ApiList.LiveGetRoomUID + "?room_id=" + roomID)
	if err != nil {
		fmt.Println(err)
	}
	return res.Data["room_info"].(map[string]interface{})["uid"].(float64)
}

// getRoomInfoOld 根据 uid 获取其 roomid
func (js JSON) getRoomInfoOld(uid string) float64 {
	res, err := utils.Get(apiquery.ApiList.RoomInfoOld + "?mid=" + uid)
	if err != nil {
		fmt.Println(err)
	}
	return res.Data["roomid"].(float64)
}

// getGiftBagList 获取背包礼物
func getGiftBagList() []interface{} {
	res, err := utils.Get(apiquery.ApiList.GiftBagList)
	if err != nil {
		log.Fatalln(err)
	}
	return res.Data["list"].([]interface{})
}

// sendGift 送出礼物
func (js JSON) sendGift(roomID string, uid string) {
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
			"&uid=" + config.Conf.Cookie.UserID +
			"&csrf=" + config.Conf.Cookie.BiliJct +
			"&send_ruid=" + "0" +
			"&storm_beat_id=" + "0" +
			"&price=" + "0" +
			"&platform=" + "pc" +
			"&biz_code=" + "live")
		res, err := utils.Post(apiquery.ApiList.GiftSend, postBody)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}
}
