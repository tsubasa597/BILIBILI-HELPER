package apiquery

// APIList Api 列表
type APIList struct {
	// 上报观看记录
	VideoHeartbeat string
	// 直播签到
	LiveCheckin string
}

var (
	ApiList APIList = APIList{}
)

func init() {
	ApiList.LiveCheckin = "https://api.live.bilibili.com/xlive/web-ucenter/v1/sign/DoSign"
	ApiList.VideoHeartbeat = "https://api.bilibili.com/x/click-interface/web/heartbeat"

}
