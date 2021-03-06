package apiquery

// APIList Api 列表
type APIList struct {
	// 上报观看记录
	VideoHeartbeat string
	// 直播签到
	LiveCheckin string
	// 登录状态
	Login string
	// 银瓜子兑换状态
	Sliver2CoinsStatus string
	// 硬币换银瓜子
	Sliver2Coins string
	// 视频分享
	AvShare string
}

var (
	ApiList APIList = APIList{}
)

func init() {
	ApiList.LiveCheckin = "https://api.live.bilibili.com/xlive/web-ucenter/v1/sign/DoSign"
	ApiList.VideoHeartbeat = "https://api.bilibili.com/x/click-interface/web/heartbeat"
	ApiList.Login = "https://api.bilibili.com/x/web-interface/nav"
	ApiList.Sliver2CoinsStatus = "https://api.live.bilibili.com/pay/v1/Exchange/getStatus"
	ApiList.Sliver2Coins = "https://api.live.bilibili.com/pay/v1/Exchange/silver2coin"
	ApiList.AvShare = "https://api.bilibili.com/x/web-interface/share/add"
}
