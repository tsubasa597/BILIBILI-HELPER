package daily

import "github.com/tsubasa597/BILIBILI-HELPER/api"

func Run(api api.API) (res string) {
	if err, ok := UserCheck(api); ok {
		res += "WatchVideo: " + WatchVideo(api, "BV1NT4y137Jc") + "\n"
		res += "ShareVideo: " + ShareVideo(api, "BV1NT4y137Jc") + "\n"
		res += "Sliver2Coins: " + Sliver2Coins(api) + "\n"
		res += "LiveCheckin: " + LiveCheckin(api)
	} else {
		res += "UserCheck: " + err
	}
	return
}

func UserCheck(api api.API) (string, bool) {
	resp, err := api.UserCheck()
	if err != nil {
		api.Entry.Debugln(err)
		return err.Error(), false
	}

	if resp.Code == 0 {
		api.Entry.Debugln("登录成功")
		return "登录成功", true
	}

	api.Entry.Debugln(resp.Message)
	return resp.Message, false
}

func WatchVideo(api api.API, bvid string) string {
	resp, err := api.WatchVideo(bvid)
	if err != nil && resp.Code != 0 {
		api.Entry.Debugln(err)
		return err.Error()
	}

	if resp.Code == 0 {
		api.Entry.Debugln("播放成功")
		return "播放成功"
	}

	api.Entry.Debugln(resp.Message)
	return resp.Message
}

func Sliver2Coins(api api.API) string {
	const exchangeRate int64 = 700
	status, err := api.Sliver2CoinsStatus()
	if err != nil {
		api.Entry.Debugln(err)
		return err.Error()
	}

	if status.Data.Silver < exchangeRate {
		api.Entry.Debugln("当前银瓜子余额不足700,不进行兑换")
		return "当前银瓜子余额不足700,不进行兑换"
	}

	resp, err := api.Sliver2Coins()

	if resp.Code == 0 {
		api.Entry.Debugln("兑换成功")
		return "兑换成功"
	}

	if resp.Code == 403 {
		return resp.Message
	}

	if err != nil {
		api.Entry.Debugln(err)
		return err.Error()
	}

	api.Entry.Debugln(resp.Message)
	return resp.Message
}

func ShareVideo(api api.API, bvid string) string {
	resp, err := api.ShareVideo(bvid)
	if err != nil && resp.Code != 0 {
		api.Entry.Debugln(err)
		return err.Error()
	}

	if resp.Code == 0 {
		api.Entry.Debugln("分享成功")
		return "分享成功"
	}

	api.Entry.Debugln(resp.Message)
	return resp.Message
}

func LiveCheckin(api api.API) string {
	resp, err := api.LiveCheckin()
	if err != nil {
		api.Entry.Debugln(err)
		return err.Error()
	}

	if resp.Code == 0 {
		api.Entry.Debugln("签到成功")
		return "签到成功"
	}

	api.Entry.Debugln("重复签到")
	return "重复签到"
}
