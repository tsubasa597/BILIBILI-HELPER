package task

import (
	"bili/apiquery"

	"bili/utils"
	"encoding/json"
	"fmt"
)

// Sliver2CoinsStatus 银瓜子换硬币状态
func (rs *JSONResponse) Sliver2CoinsStatus() {
	res, err := utils.Get(apiquery.ApiList.Sliver2CoinsStatus)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(res, &rs)
}

// Sliver2Coins 银瓜子换硬币
func (rs *JSONResponse) Sliver2Coins() {
	res, err := utils.Get(apiquery.ApiList.Sliver2Coins)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(res, &rs)
}

// Sliver2CoinsInfo 银瓜子换硬币信息
func (info *Status) Sliver2CoinsInfo(ts Tasker) {
	// 银瓜子兑换硬币汇率
	var exchangeRate float64 = 700
	ts.Sliver2CoinsStatus()
	response := ts.GetJSONResponse()
	info.Slivers = response.Data["silver"].(float64)
	info.Coins = response.Data["coin"].(float64)
	if info.Slivers < exchangeRate {
		fmt.Printf("当前银瓜子余额为: %v,不足700,不进行兑换", info.Slivers)
	} else {
		ts.Sliver2Coins()
		response = ts.GetJSONResponse()
		if response.Code == 0 {
			fmt.Println("银瓜子兑换成功")
			fmt.Printf("当前银瓜子余额: %v", (info.Slivers - exchangeRate))
			fmt.Printf("兑换银瓜子后硬币余额: %v", (info.Coins + 1))
			info.IsSliver2CoinsStatus = true
		} else {
			fmt.Println("银瓜子兑换硬币失败 原因是: " + response.Message)
		}
	}
}
