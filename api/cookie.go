package api

import (
	"github.com/spf13/viper"
	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
)

// Cookie 用于登录的必要参数
type Cookie struct {
	UserID   string
	SessData string
	BiliJct  string
}

const (
	_jct    = "Bili.biliJct"
	_sess   = "Bili.sessData"
	_userid = "Bili.userId"
)

// NewCookie 读取指定文件中的 cookie 值
func NewCookie(path string) (*Cookie, error) {
	vip := viper.New()
	vip.SetConfigFile(path)

	if err := vip.ReadInConfig(); err != nil {
		return nil, ecode.APIErr{
			E:   ecode.ErrLoad,
			Msg: err.Error(),
		}
	}

	if vip.GetString(_jct) == "" || vip.GetString(_sess) == "" || vip.GetString(_userid) == "" {
		return nil, ecode.APIErr{
			E:   ecode.ErrLoad,
			Msg: ecode.MsgNoCookie,
		}
	}

	c := &Cookie{
		BiliJct:  vip.GetString(_jct),
		SessData: vip.GetString(_sess),
		UserID:   vip.GetString(_userid),
	}
	return c, nil
}
