package task

import (
	"bili/apiquery"
	"bili/utils"
	"encoding/json"
	"fmt"
)

// LiveCheckin 直播签到
func (rs *JSONResponse) LiveCheckin() {
	res, err := utils.Get(apiquery.ApiList.LiveCheckin)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(res, &rs)
}

// LiveCheckinInfo 直播签到信息
func (info *Status) LiveCheckinInfo(ts Tasker) {
	ts.LiveCheckin()
	fmt.Println(ts)
}
