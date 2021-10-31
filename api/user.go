package api

import (
	fmt "fmt"

	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/requests"
)

// GetUserInfo 用户详情
func GetUserInfo(uid int64) (*XSpaceAccInfoResponse, error) {
	resp := &XSpaceAccInfoResponse{}
	if err := requests.Gets(fmt.Sprintf("%s?mid=%d", spaceAccInfo, uid), resp); err != nil {
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

// UserCheck 用户登录验证
func (api API) UserCheck() (*BaseResponse, error) {
	resp := &BaseResponse{}
	if err := api.Req.Gets(userLogin, resp); err != nil {
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

// GetUserName 获取用户姓名
func GetUserName(uid int64) (string, error) {
	info, err := GetUserInfo(uid)
	if err != nil {
		return "", ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	return info.Data.Name, nil
}
