package task

import (
	"bili/apiquery"
	"bili/utils"
	"log"
)

// UserCheck 用户检查
func (info *DailyInfo) UserCheck() {
	response, err := utils.Get(apiquery.ApiList.Login)
	if err != nil {
		log.Fatal(err)
	}
	if response.Code == 0 && response.Data["isLogin"].(bool) {
		info.IsLogin = true
		log.Println("Cookies有效，登录成功")
	} else {
		info.IsLogin = false
		log.Println("Cookies可能失效了,请仔细检查Github Secrets中DEDEUSERID SESSDATA BILI_JCT三项的值是否正确、过期")
		return
	}
	info.Coins = response.Data["money"].(float64)
	info.Level = response.Data["level_info"].(map[string]interface{})["current_level"].(float64)
	info.NextLevelExp = response.Data["level_info"].(map[string]interface{})["next_exp"].(float64) - response.Data["level_info"].(map[string]interface{})["current_exp"].(float64)
}
