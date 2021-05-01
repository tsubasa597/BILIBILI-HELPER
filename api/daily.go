package api

import (
	"encoding/json"
	"math/rand"
	"net/url"
	"strconv"
)

// UserCheck 用户登录验证
func (api API) UserCheck() (*Response, error) {
	rep, err := api.r.Get(UserLogin)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(rep, &resp)

	return resp, err
}

// WatchVideo 视频模拟观看，观看时间在 [0, 90) 之间
func (api API) WatchVideo(bvid string) (*Response, error) {
	data := url.Values{
		"bvid":        []string{bvid},
		"played_time": []string{strconv.Itoa(rand.Intn(90))},
	}

	rep, err := api.r.Post(VideoHeartbeat, data)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(rep, &resp)

	return resp, err
}

// ShareVideo 分享视频
func (api API) ShareVideo(bvid string) (*Response, error) {
	data := url.Values{
		"bvid": []string{bvid},
		"csrf": []string{api.conf.BiliJct},
	}
	rep, err := api.r.Post(AvShare, data)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
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
func (api API) Sliver2Coins() (*Response, error) {
	rep, err := api.r.Get(Sliver2Coins)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(rep, &resp)

	return resp, err
}

// LiveCheckin 直播签到
func (api API) LiveCheckin() (*Response, error) {
	rep, err := api.r.Get(LiveCheckin)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(rep, &resp)

	return resp, err
}
