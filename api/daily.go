package api

import (
	"math/rand"
	"net/url"
	"strconv"
	"strings"

	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
)

// WatchVideo 视频模拟观看，观看时间在 [0, 90) 之间
func (api API) WatchVideo(bvid string) (*BaseResponse, error) {
	data := url.Values{
		"bvid":        []string{bvid},
		"played_time": []string{strconv.Itoa(rand.Intn(90))},
	}

	resp := &BaseResponse{}
	if err := api.Req.Posts(videoHeartbeat, data, resp); err != nil {
		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	if resp.Code != ecode.Sucess {
		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: resp.Message,
		}
	}

	return resp, nil
}

// ShareVideo 分享视频
func (api API) ShareVideo(bvid string) (*BaseResponse, error) {
	data := url.Values{
		"bvid": []string{bvid},
		"csrf": []string{api.cookie.BiliJct},
	}

	resp := &BaseResponse{}
	if err := api.Req.Posts(avShare, data, resp); err != nil {
		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	if resp.Code != ecode.Sucess {
		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: resp.Message,
		}
	}

	return resp, nil
}

// Sliver2CoinsStatus 获取银瓜子和硬币的数量
func (api API) Sliver2CoinsStatus() (*Sliver2CoinsStatusResponse, error) {
	resp := &Sliver2CoinsStatusResponse{}

	if err := api.Req.Gets(sliver2CoinsStatus, resp); err != nil {
		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	if resp.Code != ecode.Sucess {
		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: resp.Message,
		}
	}

	return resp, nil
}

// Sliver2Coins 将银瓜子兑换为硬币
func (api API) Sliver2Coins() (*BaseResponse, error) {
	resp := &BaseResponse{}
	if err := api.Req.Gets(sliver2Coins, resp); err != nil {
		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	if resp.Code != ecode.Sucess {
		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: resp.Message,
		}
	}

	return resp, nil
}

// GetRandomAV 随机获取一个视频的 av 号
func (api API) GetRandomAV() (string, error) {
	resp := &RandomAvResponse{}
	if err := api.Req.Gets(randomAV, resp); err != nil {
		return "", ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	if resp.Code != ecode.Sucess {
		return "", ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: resp.Message,
		}
	}

	parms := strings.Split(resp.Data.Url, "/")
	if strings.HasPrefix(parms[len(parms)-1], "BV") {
		return parms[len(parms)-1], nil
	}

	return "", nil
}
