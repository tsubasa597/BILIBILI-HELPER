package info

type Live struct {
	Name        string
	Time        int32
	LiveStatus  bool
	LiveRoomURL string
	LiveTitle   string
}

var _ Infoer = (*Live)(nil)

func (Live) Type() Type {
	return TLive
}

func (live Live) GetData() []interface{} {
	return []interface{}{live.LiveStatus, live.LiveRoomURL, live.LiveTitle}
}

func (live *Live) GetInstance() interface{} {
	return live
}
