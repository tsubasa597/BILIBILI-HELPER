package daily

import "github.com/tsubasa597/BILIBILI-HELPER/api"

type Daily struct {
	api api.API
}

func NewDaily(api api.API) Daily {
	return Daily{
		api: api,
	}
}

func (d Daily) Run() (res string) {
	if err, ok := d.UserCheck(); ok {
		res += "WatchVideo: " + d.WatchVideo("BV1NT4y137Jc") + "\n"
		res += "ShareVideo: " + d.ShareVideo("BV1NT4y137Jc") + "\n"
		res += "Sliver2Coins: " + d.Sliver2Coins() + "\n"
		res += "LiveCheckin: " + d.LiveCheckin()
	} else {
		res += "UserCheck: " + err
	}
	return
}

func (d Daily) UserCheck() (string, bool) {
	resp, err := d.api.UserCheck()
	if err != nil {
		d.api.Entry.Debugln(err)
		return err.Error(), false
	}

	if resp.Code == 0 {
		d.api.Entry.Debugln("登录成功")
		return "登录成功", true
	}

	d.api.Entry.Debugln(resp.Message)
	return resp.Message, false
}

func (d Daily) WatchVideo(bvid string) string {
	resp, err := d.api.WatchVideo(bvid)
	if err != nil && resp.Code != 0 {
		d.api.Entry.Debugln(err)
		return err.Error()
	}

	if resp.Code == 0 {
		d.api.Entry.Debugln("播放成功")
		return "播放成功"
	}

	d.api.Entry.Debugln(resp.Message)
	return resp.Message
}

func (d Daily) Sliver2Coins() string {
	const exchangeRate int64 = 700
	status, err := d.api.Sliver2CoinsStatus()
	if err != nil {
		d.api.Entry.Debugln(err)
		return err.Error()
	}

	if status.Data.Silver < exchangeRate {
		d.api.Entry.Debugln("当前银瓜子余额不足700,不进行兑换")
		return "当前银瓜子余额不足700,不进行兑换"
	}

	resp, err := d.api.Sliver2Coins()

	if resp.Code == 0 {
		d.api.Entry.Debugln("兑换成功")
		return "兑换成功"
	}

	if resp.Code == 403 {
		return resp.Message
	}

	if err != nil {
		d.api.Entry.Debugln(err)
		return err.Error()
	}

	d.api.Entry.Debugln(resp.Message)
	return resp.Message
}

func (d Daily) ShareVideo(bvid string) string {
	resp, err := d.api.ShareVideo(bvid)
	if err != nil && resp.Code != 0 {
		d.api.Entry.Debugln(err)
		return err.Error()
	}

	if resp.Code == 0 {
		d.api.Entry.Debugln("分享成功")
		return "分享成功"
	}

	d.api.Entry.Debugln(resp.Message)
	return resp.Message
}

func (d Daily) LiveCheckin() string {
	resp, err := d.api.LiveCheckin()
	if err != nil {
		d.api.Entry.Debugln(err)
		return err.Error()
	}

	if resp.Code == 0 {
		d.api.Entry.Debugln("签到成功")
		return "签到成功"
	}

	d.api.Entry.Debugln("重复签到")
	return "重复签到"
}
