package info

import (
	"github.com/spf13/viper"
)

// Cookie 用于登录的必要参数
type Cookie struct {
	UserID   string
	SessData string
	BiliJct  string
}

func NewCookie(path string) (*Cookie, error) {
	vip := viper.New()
	vip.SetConfigFile(path)

	if err := vip.ReadInConfig(); err != nil {
		return nil, err
	}

	c := &Cookie{
		BiliJct:  vip.GetString("Bili.biliJct"),
		SessData: vip.GetString("Bili.sessData"),
		UserID:   vip.GetString("Bili.userId"),
	}
	return c, nil
}
