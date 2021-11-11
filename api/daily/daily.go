package daily

import (
	"math/rand"
	"net/url"
	"strconv"
	"strings"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/api/proto"
	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
)

// WatchVideo 视频模拟观看，观看时间在 [0, 90) 之间
func WatchVideo(api api.API, bvid string) (*proto.BaseResponse, error) {
	data := url.Values{
		"bvid":        []string{bvid},
		"played_time": []string{strconv.Itoa(rand.Intn(90))},
	}

	resp := &proto.BaseResponse{}
	if err := api.Req.Posts(info.VideoHeartbeat, data, resp); err != nil {
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
func ShareVideo(api api.API, bvid string) (*proto.BaseResponse, error) {
	data := url.Values{
		"bvid": []string{bvid},
		"csrf": []string{api.Cookie.BiliJct},
	}

	resp := &proto.BaseResponse{}
	if err := api.Req.Posts(info.AvShare, data, resp); err != nil {
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
func Sliver2CoinsStatus(api api.API) (*proto.Sliver2CoinsStatusResponse, error) {
	resp := &proto.Sliver2CoinsStatusResponse{}

	if err := api.Req.Gets(info.Sliver2CoinsStatus, resp); err != nil {
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
func Sliver2Coins(api api.API) (*proto.BaseResponse, error) {
	data := url.Values{
		"csrf": []string{api.Cookie.BiliJct},
	}

	resp := &proto.BaseResponse{}
	if err := api.Req.Posts(info.Sliver2Coins, data, resp); err != nil {
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
func GetRandomAV(api api.API) (string, error) {
	resp := &proto.RandomAvResponse{}
	if err := api.Req.Gets(info.RandomAV, resp); err != nil {
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

// LiveCheckin 直播签到
func LiveCheckin(api api.API) (*proto.BaseResponse, error) {
	resp := &proto.BaseResponse{}
	if err := api.Req.Gets(info.LiveCheckin, resp); err != nil {
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
