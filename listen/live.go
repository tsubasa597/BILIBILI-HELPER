package listen

import (
	"context"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
)

func (l *Listen) GetLiverStatus(uid int64) (info api.Info) {
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

func (l *Listen) LiveListen(uid int64) (context.Context, chan api.Info, error) {
	return l.AddListen(uid, l.GetLiverStatus)
}
