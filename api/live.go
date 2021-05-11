package api

import (
	"context"
	"time"
)

func (l *Listen) GetLiverStatus(uid int64) (info Info) {
	rep, err := l.api.GetUserInfo(uid)
	if err != nil {
		info.Err = err
		return
	}
	info.Name = rep.Data.Name
	info.T = int32(time.Now().Unix())
	info.LiveRoomURL = rep.Data.LiveRoom.Url
	info.LiveTitle = rep.Data.LiveRoom.Title
	if rep.Data.LiveRoom.RoomStatus == 1 {
		info.LiveStatus = true
	}
	return
}

func (l *Listen) LiveListen(uid int64, ticker *time.Ticker) (context.Context, chan Info, error) {
	return l.AddListen(uid, l.GetLiverStatus)
}
