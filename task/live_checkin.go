package task

import (
	"bili/apiquery"
	"bili/utils"
	"log"
)

// DailyLiveCheckin 直播签到信息
func (info *DailyInfo) DailyLiveCheckin() {
	response, err := utils.Get(apiquery.ApiList.LiveCheckin)
	if err != nil {
		log.Fatal(err)
	}
	if response.Code == 0 {
		log.Println("直播签到成功，本次签到获得" + response.Data["text"].(string) + "," + response.Data["specialText"].(string))

	} else {
		log.Println("直播签到失败: " + response.Message)
	}

}
