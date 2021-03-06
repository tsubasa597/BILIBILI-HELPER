package login

// Cookie 用于登录的必要参数
type Cookie struct {
	UserID   string `yaml:"userId"`
	SessData string `yaml:"sessData"`
	BiliJct  string `yaml:"biliJct"`
}

// GetVerify 将参数进行组合
func (cookie Cookie) GetVerify() string {
	// TODO 修改组合方式
	return "bili_jct=" + cookie.BiliJct + ";SESSDATA=" + cookie.SessData + ";DedeUserID=" + cookie.UserID + ";"
}
