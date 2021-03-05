package task

import (
	"bili/apiquery"
	"bili/utils"
	"encoding/json"
	"fmt"
)

// VideoWatch 观看视频
func (rs *Response) VideoWatch(bvid string) {
	// TODO 修改 played_time
	postBody := []byte("bvid=" + bvid + "&played_time=21")
	res, err := utils.Post(apiquery.ApiList.VideoHeartbeat, postBody)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(res, &rs)
}
