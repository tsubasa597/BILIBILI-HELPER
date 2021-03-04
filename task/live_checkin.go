package task

import (
	"bili/apiquery"
	"bili/utils"
	"fmt"
)

// LiveCheckin 直播签到
func LiveCheckin() {
	res, err := utils.Get(apiquery.ApiList.LiveCheckin)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
