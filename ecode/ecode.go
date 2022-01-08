package ecode

type APIErr struct {
	E   string
	Msg string
}

const (
	// Sucess 接口请求成功 code 值
	Sucess = iota
)

const (
	SucessLogin  = "登录成功"
	SucessPlay   = "播放成功"
	SucessShare  = "分享成功"
	SucessSignIn = "直播签到成功"
)

const (
	ErrGetInfo       = "请求发生错误"
	ErrLoad          = "解析错误"
	ErrPause         = "正在暂停"
	ErrorStop        = "已经停止"
	ErrNoDynamic     = "该用户没有更多动态"
	ErrNoComment     = "没有更多评论"
	ErrUnknowDynamic = "未知动态"
	ErrExchange      = "当前银瓜子余额不足700,不进行兑换"
	ErrNoBvID        = "Bvid 为空，跳过"
	ErrNoCookie      = "cookie 文件缺少参数"
)

var (
	_ error = (*APIErr)(nil)
)

func (err APIErr) Error() string {
	if err.Msg != "" {
		return err.E + ": " + err.Msg
	}

	return err.E
}
