package api

import (
	"github.com/spf13/viper"
	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
)

// Cookie 用于登录的必要参数
type cookie struct {
	userID    string
	sessData  string
	biliJct   string
	liveBuvid string
}

const (
	_jct    = "Bili.biliJct"
	_sess   = "Bili.sessData"
	_userid = "Bili.userId"
	_buvid  = "Bili.buvid"
)

// NewCookie 读取指定文件中的 cookie 值
func NewCookie(path string) (*cookie, error) {
	vip := viper.New()
	vip.SetConfigFile(path)

	if err := vip.ReadInConfig(); err != nil {
		return nil, ecode.APIErr{
			E:   ecode.ErrLoad,
			Msg: err.Error(),
		}
	}

	if vip.GetString(_jct) == "" || vip.GetString(_sess) == "" ||
		vip.GetString(_userid) == "" || vip.GetString(_buvid) == "" {
		return nil, ecode.APIErr{
			E:   ecode.ErrLoad,
			Msg: ecode.MsgNoCookie,
		}
	}

	c := &cookie{
		biliJct:   vip.GetString(_jct),
		sessData:  vip.GetString(_sess),
		userID:    vip.GetString(_userid),
		liveBuvid: vip.GetString(_buvid),
	}
	return c, nil
}
