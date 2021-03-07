package task

import (
	"bili/apiquery"
	"bili/utils"
	"encoding/json"
	"fmt"
)

// LiveCheckin 直播签到
func (rs *Response) liveCheckin() {
	res, err := utils.Get(apiquery.ApiList.LiveCheckin)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(res, &rs)
}

// DailyLiveCheckin 直播签到信息
func (info *Status) DailyLiveCheckin() {
	info.rs.liveCheckin()
	response := info.rs
	if response.Code == 0 {
		fmt.Println("直播签到成功，本次签到获得" + response.Data["text"].(string) + "," + response.Data["specialText"].(string))
	} else {
		fmt.Println("直播签到失败: " + response.Message)
	}

}
