package api

import (
	"encoding/json"
	"math/rand"
	"net/url"
	"strconv"
)

// UserCheck 用户登录验证
func (api API) UserCheck() (*BaseResponse, error) {
	rep, err := api.r.Get(UserLogin)
	if err != nil {
		return nil, err
	}

	resp := &BaseResponse{}
	err = json.Unmarshal(rep, &resp)

	return resp, err
}

// WatchVideo 视频模拟观看，观看时间在 [0, 90) 之间
func (api API) WatchVideo(bvid string) (*BaseResponse, error) {
	data := url.Values{
		"bvid":        []string{bvid},
		"played_time": []string{strconv.Itoa(rand.Intn(90))},
	}

	rep, err := api.r.Post(VideoHeartbeat, data)
	if err != nil {
		return nil, err
	}

	resp := &BaseResponse{}
	err = json.Unmarshal(rep, &resp)

	return resp, err
}

// ShareVideo 分享视频
func (api API) ShareVideo(bvid string) (*BaseResponse, error) {
	data := url.Values{
		"bvid": []string{bvid},
		"csrf": []string{api.conf.BiliJct},
	}

	rep, err := api.r.Post(AvShare, data)
	if err != nil {
		return nil, err
	}

	resp := &BaseResponse{}
	err = json.Unmarshal(rep, &resp)

	return resp, err
}

// Sliver2CoinsStatus 获取银瓜子和硬币的数量
func (api API) Sliver2CoinsStatus() (*Sliver2CoinsStatusResponse, error) {
	rep, err := api.r.Get(Sliver2CoinsStatus)
	if err != nil {
		return nil, err
	}

	resp := &Sliver2CoinsStatusResponse{}
	err = json.Unmarshal(rep, &resp)

	return resp, err
}

// Sliver2Coins 将银瓜子兑换为硬币
func (api API) Sliver2Coins() (*BaseResponse, error) {
	rep, err := api.r.Get(Sliver2Coins)
	if err != nil {
		return nil, err
	}

	resp := &BaseResponse{}
	err = json.Unmarshal(rep, &resp)

	return resp, err
}

// LiveCheckin 直播签到
func (api API) LiveCheckin() (*BaseResponse, error) {
	rep, err := api.r.Get(LiveCheckin)
	if err != nil {
		return nil, err
	}

	resp := &BaseResponse{}
	err = json.Unmarshal(rep, &resp)

	return resp, err
}

type Daily struct {
	api API
}

func NewDaily(api API) Daily {
	return Daily{
		api: api,
	}
}

func (d Daily) Run() (res string) {
	if err, ok := d.userCheck(); ok {
		res += "WatchVideo: " + d.watchVideo("BV1NT4y137Jc") + "\n"
		res += "ShareVideo: " + d.shareVideo("BV1NT4y137Jc") + "\n"
		res += "Sliver2Coins: " + d.sliver2Coins() + "\n"
		res += "LiveCheckin: " + d.liveCheckin()
	} else {
		res += "UserCheck: " + err
	}
	return
}

func (d Daily) userCheck() (string, bool) {
	resp, err := d.api.UserCheck()
	if err != nil {
		d.api.log.Debugln(err)
		return err.Error(), false
	}

	if resp.Code == 0 {
		d.api.log.Debugln("登录成功")
		return "登录成功", true
	}

	d.api.log.Debugln(resp.Message)
	return resp.Message, false
}

func (d Daily) watchVideo(bvid string) string {
	resp, err := d.api.WatchVideo(bvid)
	if err != nil && resp.Code != 0 {
		d.api.log.Debugln(err)
		return err.Error()
	}

	if resp.Code == 0 {
		d.api.log.Debugln("播放成功")
		return "播放成功"
	}

	d.api.log.Debugln(resp.Message)
	return resp.Message
}

func (d Daily) sliver2Coins() string {
	const exchangeRate int64 = 700
	status, err := d.api.Sliver2CoinsStatus()
	if err != nil {
		d.api.log.Debugln(err)
		return err.Error()
	}

	if status.Data.Silver < exchangeRate {
		d.api.log.Debugln("当前银瓜子余额不足700,不进行兑换")
		return "当前银瓜子余额不足700,不进行兑换"
	}

	resp, err := d.api.Sliver2Coins()

	if resp.Code == 0 {
		d.api.log.Debugln("兑换成功")
		return "兑换成功"
	}

	if resp.Code == 403 {
		return resp.Message
	}

	if err != nil {
		d.api.log.Debugln(err)
		return err.Error()
	}

	d.api.log.Debugln(resp.Message)
	return resp.Message
}

func (d Daily) shareVideo(bvid string) string {
	resp, err := d.api.ShareVideo(bvid)
	if err != nil && resp.Code != 0 {
		d.api.log.Debugln(err)
		return err.Error()
	}

	if resp.Code == 0 {
		d.api.log.Debugln("分享成功")
		return "分享成功"
	}

	d.api.log.Debugln(resp.Message)
	return resp.Message
}

func (d Daily) liveCheckin() string {
	resp, err := d.api.LiveCheckin()
	if err != nil {
		d.api.log.Debugln(err)
		return err.Error()
	}

	if resp.Code == 0 {
		d.api.log.Debugln("签到成功")
		return "签到成功"
	}

	d.api.log.Debugln("重复签到")
	return "重复签到"
}
