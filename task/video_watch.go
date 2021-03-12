package task

import (
	"bili/apiquery"
	"bili/config"
	"bili/utils"
	"log"
	"math/rand"
	"strconv"
)

// DailyVideo 观看视频
func (info *DailyInfo) DailyVideo(bvid string) {
	postBody := []byte("bvid=" + bvid + "&played_time=" + strconv.Itoa(rand.Intn(90)))
	response, err := utils.Post(apiquery.ApiList.VideoHeartbeat, postBody)
	if err != nil {
		log.Fatal(err)
	}
	if response.Code == 0 {
		log.Println("视频播放成功")
	} else {
		log.Println("视频播放失败,原因: " + response.Message)
	}
}

// DailyVideoShare 分享视频
func (info *DailyInfo) DailyVideoShare(bvid string) {
	postBody := []byte("bvid=" + bvid + "&csrf=" + config.Conf.Cookie.BiliJct)
	response, err := utils.Post(apiquery.ApiList.AvShare, postBody)
	if err != nil {
		log.Fatal(err)
	}
	if response.Code == 0 {
		log.Println("视频分享成功")
	} else {
		log.Println("视频分享失败,原因: " + response.Message)
	}
}
