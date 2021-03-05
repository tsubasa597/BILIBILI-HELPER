package task

import (
	"bili/apiquery"
	"bili/utils"
	"encoding/json"
	"fmt"
)

// LiveCheckin 直播签到
func (rs *Response) LiveCheckin() {
	res, err := utils.Get(apiquery.ApiList.LiveCheckin)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(res, &rs)
}
