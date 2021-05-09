package global

import (
	"fmt"

	"github.com/spf13/viper"
)

// Cookie 用于登录的必要参数
type Cookie struct {
	UserID   string
	SessData string
	BiliJct  string
}

// GetVerify 将参数进行组合
func (cookie Cookie) GetVerify() string {
	return fmt.Sprintf("bili_jct=%s;SESSDATA=%s;DedeUserID=%s;", cookie.BiliJct, cookie.SessData, cookie.UserID)
}

func NewConfig(path string) Cookie {
	vip := viper.New()
	vip.SetConfigFile(path)
	vip.ReadInConfig()

	c := Cookie{
		BiliJct:  vip.GetString("Bili.biliJct"),
		SessData: vip.GetString("Bili.sessData"),
		UserID:   vip.GetString("Bili.userId"),
	}
	return c
}
