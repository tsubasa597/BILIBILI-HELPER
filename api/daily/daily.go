package daily

import (
	"math/rand"
	"net/url"
	"strconv"
	"strings"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/api/proto"
	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
)

const (
	// 银瓜子最低兑换要求
	_exchangeRate int64 = 700
)

// WatchVideo 视频模拟观看，观看时间在 [0, 90) 之间
func WatchVideo(ap api.API, bvid string) error {
	data := url.Values{
		"bvid":        []string{bvid},
		"played_time": []string{strconv.Itoa(rand.Intn(90))},
	}

	resp := &proto.BaseResponse{}
	if err := ap.Req.Posts(api.VideoHeartbeat, data, resp); err != nil {
		return ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	if resp.Code != ecode.Sucess {
		return ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: resp.Message,
		}
	}

	return nil
}

// ShareVideo 分享视频
func ShareVideo(ap api.API, bvid string) error {
	data := url.Values{
		"bvid": []string{bvid},
		"csrf": []string{ap.GetJwt()},
	}

	resp := &proto.BaseResponse{}
	if err := ap.Req.Posts(api.AvShare, data, resp); err != nil {
		return ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	if resp.Code != ecode.Sucess {
		return ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: resp.Message,
		}
	}

	return nil
}

// Sliver2CoinsStatus 获取银瓜子和硬币的数量
func Sliver2CoinsStatus(ap api.API) (*proto.Sliver2CoinsStatusResponse, error) {
	resp := &proto.Sliver2CoinsStatusResponse{}

	if err := ap.Req.Gets(api.Sliver2CoinsStatus, resp); err != nil {
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
func Sliver2Coins(ap api.API) error {
	status, err := Sliver2CoinsStatus(ap)
	if err != nil {
		return err
	}

	if status.Data.Silver < _exchangeRate {
		return ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: ecode.MsgExchange,
		}
	}

	data := url.Values{
		"csrf": []string{ap.GetJwt()},
	}

	resp := &proto.BaseResponse{}
	if err := ap.Req.Posts(api.Sliver2Coins, data, resp); err != nil {
		return ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	if resp.Code != ecode.Sucess {
		return ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: resp.Message,
		}
	}

	return nil
}

// GetRandomAV 随机获取一个视频的 av 号
func GetRandomAV(ap api.API) (string, error) {
	resp := &proto.RandomAvResponse{}
	if err := ap.Req.Gets(api.RandomAV, resp); err != nil {
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

	return "", ecode.APIErr{
		E:   ecode.ErrGetInfo,
		Msg: ecode.ErrLoad,
	}
}

// LiveCheckin 直播签到
func LiveCheckin(ap api.API) error {
	resp := &proto.BaseResponse{}
	if err := ap.Req.Gets(api.LiveCheckin, resp); err != nil {
		return ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	if resp.Code != ecode.Sucess {
		return ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: resp.Message,
		}
	}

	return nil
}
