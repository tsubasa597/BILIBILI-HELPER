package task

import (
	"bili/apiquery"

	"bili/utils"
	"encoding/json"
	"fmt"
)

// Sliver2CoinsStatus 银瓜子换硬币状态
func (rs *Response) Sliver2CoinsStatus() {
	res, err := utils.Get(apiquery.ApiList.Sliver2CoinsStatus)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(res, &rs)
}

// Sliver2Coins 银瓜子换硬币
func (rs *Response) Sliver2Coins() {
	res, err := utils.Get(apiquery.ApiList.Sliver2Coins)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(res, &rs)
}
