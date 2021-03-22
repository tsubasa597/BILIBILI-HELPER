package task

import (
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/conf"
	"github.com/tsubasa597/BILIBILI-HELPER/utils"
)

// Tasker 任务
type Tasker interface {
	Run()
}

// Task task 类型的函数
type Task func(v ...string)

// Run task 类型的函数调用
func (t Task) Run(wg *sync.WaitGroup, v ...string) {
	defer wg.Done()
	t(v...)
}

// Start 启动任务
func Start() {
	var wg sync.WaitGroup
	task := New()
	for _, i := range task.tasks {
		// 防止请求过快出错
		time.Sleep(time.Second)
		wg.Add(1)
		go i.Run(&wg, task.params...)
	}

	wg.Wait()
	task.done <- 1
}

// UserCheck 用户检查
func (info *Daily) UserCheck() {
	response, err := utils.Get(api.ApiList.Login)
	if err != nil {
		info.logInfo <- []interface{}{"Fatal", err}
	}
	if response.Code == 0 && response.Data["isLogin"].(bool) {
		info.IsLogin = true
		info.logInfo <- []interface{}{"Info", "Cookies有效，登录成功"}
	} else {
		info.IsLogin = false
		info.logInfo <- []interface{}{"Fatal", "Cookies可能失效了,请仔细检查Github Secrets中DEDEUSERID SESSDATA BILI_JCT三项的值是否正确、过期"}
	}
	info.Coins = response.Data["money"].(float64)
	info.Level = response.Data["level_info"].(map[string]interface{})["current_level"].(float64)
	info.NextLevelExp = response.Data["level_info"].(map[string]interface{})["next_exp"].(float64) - response.Data["level_info"].(map[string]interface{})["current_exp"].(float64)
}

// DailyVideo 观看视频
func (info *Daily) DailyVideo(param ...string) {
	postBody := []byte("bvid=" + param[0] + "&played_time=" + strconv.Itoa(rand.Intn(90)))
	response, err := utils.Post(api.ApiList.VideoHeartbeat, postBody)
	if err != nil {
		info.logInfo <- []interface{}{"Fatal", err}
	}
	if response.Code == 0 {
		info.logInfo <- []interface{}{"Info", "视频播放成功"}
	} else {
		info.logInfo <- []interface{}{"Warn", "视频播放失败,原因: " + response.Message}
	}
}

// DailyVideoShare 分享视频
func (info *Daily) DailyVideoShare(param ...string) {
	postBody := []byte("bvid=" + param[0] + "&csrf=" + info.conf.Cookie.BiliJct)
	response, err := utils.Post(api.ApiList.AvShare, postBody)
	if err != nil && response.Code != 0 {
		info.logInfo <- []interface{}{"Fatal", err}
	}
	if response.Code == 0 {
		info.logInfo <- []interface{}{"Info", "视频分享成功"}
	} else {
		info.logInfo <- []interface{}{"Warn", "视频分享失败,原因: " + response.Message}
	}
}

// DailySliver2Coin 银瓜子换硬币信息
func (info *Daily) DailySliver2Coin(param ...string) {
	// 银瓜子兑换硬币汇率
	var exchangeRate float64 = 700
	response, err := utils.Get(api.ApiList.Sliver2CoinsStatus)
	if err != nil {
		info.logInfo <- []interface{}{"Fatal", err}
	}
	info.Slivers = response.Data["silver"].(float64)
	info.Coins = response.Data["coin"].(float64)
	if info.Slivers < exchangeRate {
		info.logInfo <- []interface{}{"Error", "当前银瓜子余额为: ", info.Slivers, "，不足700,不进行兑换"}
	} else {
		response, err = utils.Get(api.ApiList.Sliver2Coins)
		if response.Code != 403 && err != nil {
			conf.Log.Fatal(err)
		}
		if response.Code == 0 {
			info.logInfo <- []interface{}{"Info", "银瓜子兑换成功: "}
			info.Coins++
			info.Slivers -= exchangeRate
			info.logInfo <- []interface{}{"Info", "当前银瓜子余额: ", info.Slivers}
			info.logInfo <- []interface{}{"Info", "兑换银瓜子后硬币余额: ", info.Coins}
		} else {
			info.logInfo <- []interface{}{"Warn", "银瓜子兑换硬币失败 原因是: " + response.Message}
		}
	}
}

// DailyLiveCheckin 直播签到信息
func (info *Daily) DailyLiveCheckin(param ...string) {
	response, err := utils.Get(api.ApiList.LiveCheckin)
	if err != nil {
		info.logInfo <- []interface{}{"Fatal", err}
	}
	if response.Code == 0 {
		info.logInfo <- []interface{}{"Info", "直播签到成功，本次签到获得" + response.Data["text"].(string) + "," + response.Data["specialText"].(string)}
	} else {
		info.logInfo <- []interface{}{"Warn", "直播签到失败: " + response.Message}
	}

}
