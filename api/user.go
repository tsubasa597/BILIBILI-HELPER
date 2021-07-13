package api

import (
	fmt "fmt"

	"github.com/tsubasa597/requests"
)

func GetUserInfo(uid int64) (*XSpaceAccInfoResponse, error) {
	resp := &XSpaceAccInfoResponse{}
	err := requests.Gets(fmt.Sprintf("%s?mid=%d", spaceAccInfo, uid), resp)

	return resp, err
}

// UserCheck 用户登录验证
func (api API) UserCheck() (*BaseResponse, error) {
	resp := &BaseResponse{}
	err := api.Req.Gets(userLogin, resp)

	return resp, err
}

func GetUserName(uid int64) (string, error) {
	info, err := GetUserInfo(uid)
	if err != nil {
		return "", err
	}

	return info.Data.Name, nil
}
