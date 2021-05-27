package info

type Live struct {
	Info
	LiveStatus  bool
	LiveRoomURL string
	LiveTitle   string
}

var _ Infoer = (*Live)(nil)

func (Live) Type() Type {
	return TLive
}

func (live Live) GetData() []interface{} {
	return []interface{} {live.LiveStatus, live.LiveRoomURL, live.LiveTitle}
}