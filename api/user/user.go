package user

import (
	fmt "fmt"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/api/proto"
	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/requests"
)

// Info 用户部分信息
type Info struct {
	UID  int64
	Name string
	Face string
	Sign string
}

// GetInfo 获取用户详情
func GetInfo(uid int64) (*proto.XSpaceAccInfoResponse, error) {
	resp := &proto.XSpaceAccInfoResponse{}
	if err := requests.Gets(fmt.Sprintf("%s?mid=%d", api.SpaceAccInfo, uid), resp); err != nil {
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

// Check 用户登录验证
func Check(ap api.API) error {
	resp := &proto.BaseResponse{}
	if err := ap.Req.Gets(api.UserLogin, resp); err != nil {
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

// GetName 获取用户姓名
func GetName(uid int64) (string, error) {
	info, err := GetInfo(uid)
	if err != nil {
		return "", ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	return info.Data.Name, nil
}
