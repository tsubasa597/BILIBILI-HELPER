package task

import (
	"bili/config"
	"bili/utils"
)

// DailySliver2Coin 银瓜子换硬币信息
func (info *DailyInfo) DailySliver2Coin() {
	// 银瓜子兑换硬币汇率
	var exchangeRate float64 = 700
	response, err := utils.Get(config.ApiList.Sliver2CoinsStatus)
	if err != nil {
		config.Log.Fatal(err)
	}
	info.Slivers = response.Data["silver"].(float64)
	info.Coins = response.Data["coin"].(float64)
	if info.Slivers < exchangeRate {
		config.Log.Error("当前银瓜子余额为: %v,不足700,不进行兑换", info.Slivers)
	} else {
		response, err = utils.Get(config.ApiList.Sliver2Coins)
		if response.Code != 403 && err != nil {
			config.Log.Fatal(err)
		}
		if response.Code == 0 {
			config.Log.Info("银瓜子兑换成功")
			info.Coins++
			info.Slivers -= exchangeRate
			config.Log.Info("当前银瓜子余额: %v", info.Slivers)
			config.Log.Info("兑换银瓜子后硬币余额: %v", (info.Coins))
		} else {
			config.Log.Warning("银瓜子兑换硬币失败 原因是: " + response.Message)
		}
	}
}
