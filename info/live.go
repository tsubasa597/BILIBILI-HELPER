package info

// Live 监听的直播信息
type Live struct {
	Name        string
	Time        int32
	LiveStatus  bool
	LiveRoomURL string
	LiveTitle   string
}

var _ Infoer = (*Live)(nil)

// Type 判断监听类型，转换 interface{}
func (Live) Type() Type {
	return TLive
}

// GetInstance 将 Infoer 接口转为 interface{}
func (live *Live) GetInstance() interface{} {
	return live
}
