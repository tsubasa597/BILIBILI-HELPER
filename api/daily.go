package api

import (
	"math/rand"
	"net/url"
	"strconv"
	"strings"
)

// WatchVideo 视频模拟观看，观看时间在 [0, 90) 之间
func (api API) WatchVideo(bvid string) (*BaseResponse, error) {
	data := url.Values{
		"bvid":        []string{bvid},
		"played_time": []string{strconv.Itoa(rand.Intn(90))},
	}

	resp := &BaseResponse{}
	err := api.Req.Posts(videoHeartbeat, data, resp)

	return resp, err
}

// ShareVideo 分享视频
func (api API) ShareVideo(bvid string) (*BaseResponse, error) {
	data := url.Values{
		"bvid": []string{bvid},
		"csrf": []string{api.cookie.BiliJct},
	}

	resp := &BaseResponse{}
	err := api.Req.Posts(avShare, data, resp)

	return resp, err
}

// Sliver2CoinsStatus 获取银瓜子和硬币的数量
func (api API) Sliver2CoinsStatus() (*Sliver2CoinsStatusResponse, error) {
	resp := &Sliver2CoinsStatusResponse{}
	err := api.Req.Gets(sliver2CoinsStatus, resp)

	return resp, err
}

// Sliver2Coins 将银瓜子兑换为硬币
func (api API) Sliver2Coins() (*BaseResponse, error) {
	resp := &BaseResponse{}
	err := api.Req.Gets(sliver2Coins, resp)

	return resp, err
}

func (api API) GetRandomAV() (string, error) {
	resp := &RandomAvResponse{}
	err := api.Req.Gets(randomAV, resp)
	if err != nil {
		return "", err
	}
	parms := strings.Split(resp.Data.Url, "/")
	if strings.HasPrefix(parms[len(parms)-1], "BV") {
		return parms[len(parms)-1], nil
	}
	return "", nil
}
