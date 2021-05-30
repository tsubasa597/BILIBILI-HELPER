package api

import (
	"github.com/spf13/viper"
)

// Cookie 用于登录的必要参数
type cookie struct {
	UserID   string
	SessData string
	BiliJct  string
}

func newCookie(path string) cookie {
	vip := viper.New()
	vip.SetConfigFile(path)
	vip.ReadInConfig()

	c := cookie{
		BiliJct:  vip.GetString("Bili.biliJct"),
		SessData: vip.GetString("Bili.sessData"),
		UserID:   vip.GetString("Bili.userId"),
	}
	return c
}
