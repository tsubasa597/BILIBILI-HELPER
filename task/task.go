package task

import (
	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/global"
)

type Bili struct {
	api.API
}

func New(c global.Cookie) *Bili {
	return &Bili{
		API: *api.New(c),
	}
}

func Run(b *Bili) (res string) {
	if err, ok := b.userCheck(); ok {
		res += "WatchVideo: " + b.watchVideo("BV1NT4y137Jc") + "\n"
		res += "ShareVideo: " + b.shareVideo("BV1NT4y137Jc") + "\n"
		res += "Sliver2Coins: " + b.sliver2Coins() + "\n"
		res += "LiveCheckin: " + b.liveCheckin()
	} else {
		res += "UserCheck: " + err
	}
	return
}

func (b *Bili) userCheck() (string, bool) {
	resp, err := b.UserCheck()
	if err != nil {
		return err.Error(), false
	}

	if resp.Code == 0 {
		return "登录成功", true
	}

	return resp.Message, false
}

func (b *Bili) watchVideo(bvid string) string {
	resp, err := b.WatchVideo(bvid)
	if err != nil {
		return err.Error()
	}

	if resp.Code == 0 {
		return "播放成功"
	}

	return resp.Message
}

func (b *Bili) sliver2Coins() string {
	const exchangeRate int64 = 700
	status, err := b.Sliver2CoinsStatus()
	if err != nil {
		return err.Error()
	}

	if status.Data.Silver < exchangeRate {
		return "当前银瓜子余额不足700,不进行兑换"
	}

	resp, err := b.Sliver2Coins()

	if resp.Code == 0 {
		return "兑换成功"
	}

	if resp.Code == 403 {
		return resp.Message
	}

	if err != nil {
		return err.Error()
	}

	return resp.Message
}

func (b *Bili) shareVideo(bvid string) string {
	resp, err := b.ShareVideo(bvid)
	if err != nil {
		return err.Error()
	}

	if resp.Code == 0 {
		return "分享成功"
	}

	return resp.Message
}

func (b *Bili) liveCheckin() string {
	resp, err := b.LiveCheckin()
	if err != nil {
		return err.Error()
	}

	if resp.Code == 0 {
		return "签到成功"
	}

	return "重复签到"
}
