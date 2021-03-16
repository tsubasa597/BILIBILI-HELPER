package task

import (
	"bili/config"
	"bili/utils"
	"math/rand"
	"strconv"
	"sync"
)

// DailyVideo 观看视频
func (info *DailyInfo) DailyVideo(wg *sync.WaitGroup, param ...string) {
	defer wg.Done()
	postBody := []byte("bvid=" + param[0] + "&played_time=" + strconv.Itoa(rand.Intn(90)))
	response, err := utils.Post(config.ApiList.VideoHeartbeat, postBody)
	if err != nil {
		config.Log.Fatal(err)
	}
	if response.Code == 0 {
		config.Log.Info("视频播放成功")
	} else {
		config.Log.Warning("视频播放失败,原因: " + response.Message)
	}
}

// DailyVideoShare 分享视频
func (info *DailyInfo) DailyVideoShare(wg *sync.WaitGroup, param ...string) {
	defer wg.Done()
	postBody := []byte("bvid=" + param[0] + "&csrf=" + config.Conf.Cookie.BiliJct)
	response, err := utils.Post(config.ApiList.AvShare, postBody)
	if err != nil && response.Code != 0 {
		config.Log.Fatal(err)
	}
	if response.Code == 0 {
		config.Log.Info("视频分享成功")
	} else {
		config.Log.Warnln("视频分享失败,原因: " + response.Message)
	}
}
