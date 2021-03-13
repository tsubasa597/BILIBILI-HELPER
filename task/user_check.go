package task

import (
	"bili/config"
	"bili/utils"
)

// UserCheck 用户检查
func (info *DailyInfo) UserCheck() {
	response, err := utils.Get(config.ApiList.Login)
	if err != nil {
		config.Log.Fatal(err)
	}
	if response.Code == 0 && response.Data["isLogin"].(bool) {
		info.IsLogin = true
		config.Log.Info("Cookies有效，登录成功")
	} else {
		info.IsLogin = false
		config.Log.Fatal("Cookies可能失效了,请仔细检查Github Secrets中DEDEUSERID SESSDATA BILI_JCT三项的值是否正确、过期")
	}
	info.Coins = response.Data["money"].(float64)
	info.Level = response.Data["level_info"].(map[string]interface{})["current_level"].(float64)
	info.NextLevelExp = response.Data["level_info"].(map[string]interface{})["next_exp"].(float64) - response.Data["level_info"].(map[string]interface{})["current_exp"].(float64)
}
