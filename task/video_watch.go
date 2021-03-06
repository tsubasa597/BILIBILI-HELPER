package task

import (
	"bili/apiquery"
	"bili/utils"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
)

// VideoWatch 观看视频
func (rs *JSONResponse) VideoWatch(bvid string) {
	postBody := []byte("bvid=" + bvid + "&played_time=" + strconv.Itoa(rand.Intn(90)))
	res, err := utils.Post(apiquery.ApiList.VideoHeartbeat, postBody)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(res, &rs)
}

// VideoShare 分享视频
func (rs *JSONResponse) VideoShare(bvid string) {
	postBody := []byte("bvid=" + bvid + "&csrf=" + utils.Http.Verify.BiliJct)
	res, err := utils.Post(apiquery.ApiList.AvShare, postBody)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(res, &rs)
}

// DailyVideo 观看视频
func (info *Status) DailyVideo(ts Tasker) {
	ts.VideoWatch("BV1NT4y137Jc")
	response := ts.GetJSONResponse()
	if response.Code == 0 {
		info.IsVideoWatch = true
		fmt.Println("视频播放成功")
	} else {
		fmt.Println("视频播放失败,原因: " + response.Message)
	}
	if !info.IsVideoShare {
		ts.VideoShare("BV1NT4y137Jc")
		response = ts.GetJSONResponse()
		if response.Code == 0 {
			info.IsVideoShare = true
			fmt.Println("视频分享成功")
		} else {
			fmt.Println("视频分享失败，原因: " + response.Message)
		}
	}
}
