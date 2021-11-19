package info

// Live 监听的直播信息
type Live struct {
	Info
	LiveStatus  bool
	LiveRoomURL string
	LiveTitle   string
}
