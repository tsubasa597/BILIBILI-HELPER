package task

import (
	"bili/config"
	"bili/utils"
)

// DailyLiveCheckin 直播签到信息
func (info *DailyInfo) DailyLiveCheckin(param ...string) {
	response, err := utils.Get(config.ApiList.LiveCheckin)
	if err != nil {
		config.Log.Fatal(err)
	}
	if response.Code == 0 {
		config.Log.Info("直播签到成功，本次签到获得" + response.Data["text"].(string) + "," + response.Data["specialText"].(string))

	} else {
		config.Log.Warning("直播签到失败: " + response.Message)
	}

}
