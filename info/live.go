package info

type Live struct {
	Info
	LiveStatus  bool
	LiveRoomURL string
	LiveTitle   string
}

func (live Live) GetData() []interface{} {
	return []interface{}{live.LiveRoomURL, live.LiveStatus, live.LiveTitle}
}
