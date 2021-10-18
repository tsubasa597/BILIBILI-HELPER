package ecode

const (
	// Sucess 接口请求成功 code 值
	Sucess = iota

	SucessLogin  = "登录成功"
	SucessPlay   = "播放成功"
	SucessShare  = "分享成功"
	SucessSignIn = "直播签到成功"

	ErrNoDynamic     = "该用户没有更多动态"
	ErrNoComment     = "没有更多评论"
	ErrUnknowDynamic = "未知动态"
	ErrExchange      = "当前银瓜子余额不足700,不进行兑换"
	ErrNoBvID        = "Bvid 为空，跳过"
)
