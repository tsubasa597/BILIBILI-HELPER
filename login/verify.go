package login

// UserInfo 用于登录的必要参数
type UserInfo struct {
	UserID   string
	SessData string
	BiliJct  string
}

// GetVerify 将参数进行组合
func (usif UserInfo) GetVerify() string {
	return "bili_jct=" + usif.BiliJct + ";SESSDATA=" + usif.SessData + ";DedeUserID=" + usif.UserID + ";"
}
