package info

// Live 监听的直播信息
type Live struct {
	Info
	LiveStatus  bool
	LiveRoomURL string
	LiveTitle   string
}

var _ Interface = (*Live)(nil)

// GetInstance 将 Infoer 接口转为 interface{}
func (live *Live) GetInstance() interface{} {
	return live
}
