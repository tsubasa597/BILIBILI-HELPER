package task

import (
	"bili/apiquery"
	"bili/utils"
	"encoding/json"
	"fmt"
)

// UserCheck 用户检查
func (rs *Response) UserCheck() {
	res, err := utils.Get(apiquery.ApiList.Login)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(res, &rs)
	if rs.Code == 0 && rs.Data["isLogin"].(bool) {
		fmt.Println("登录成功")
	} else {
		fmt.Println("登录失败")
	}
}
