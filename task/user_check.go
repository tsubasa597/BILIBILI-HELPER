package task

import (
	"bili/apiquery"
	"bili/utils"
	"encoding/json"
	"fmt"
)

// UserCheck 用户检查
func (rs *JSONResponse) UserCheck() {
	res, err := utils.Get(apiquery.ApiList.Login)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(res, &rs)
}

// UserCheckInfo 用户检查
func (info *Status) UserCheckInfo(ts Tasker) {
	ts.UserCheck()
	var response *JSONResponse = ts.GetJSONResponse()
	if response == nil {
		fmt.Println("用户信息请求失败，如果是412错误，请在config.json中更换UA，412问题仅影响用户信息确认，不影响任务")
	} else {
		if response.Code == 0 && response.Data["isLogin"].(bool) {
			fmt.Println("Cookies有效，登录成功")
		} else {
			fmt.Println("Cookies可能失效了,请仔细检查Github Secrets中DEDEUSERID SESSDATA BILI_JCT三项的值是否正确、过期")
		}
	}
	info.IsLogin = true
	info.Coins = response.Data["money"].(float64)
	info.Level = response.Data["level_info"].(map[string]interface{})["current_level"].(float64)
	info.NextLevelExp = response.Data["level_info"].(map[string]interface{})["next_exp"].(float64) - response.Data["level_info"].(map[string]interface{})["current_exp"].(float64)
}
