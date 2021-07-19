package listen

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
)

type Live struct {
	ups sync.Map
}

func getLiverStatus(uid int64, log *logrus.Entry) (info info.Live) {
	rep, err := api.GetUserInfo(uid)
	if err != nil {
		log.Errorln(err)
		return
	}
	info.Name = rep.Data.Name
	info.Time = int32(time.Now().Unix())
	info.LiveRoomURL = rep.Data.LiveRoom.Url
	info.LiveTitle = rep.Data.LiveRoom.Title
	if rep.Data.LiveRoom.RoomStatus == 1 {
		info.LiveStatus = true
	}
	return
}

var _ Listener = (*Live)(nil)

func (live *Live) Listen(uid int64, _ api.API, log *logrus.Entry) []info.Infoer {
	status := getLiverStatus(uid, log)
	return []info.Infoer{&status}
}

func (live *Live) StopListenUP(uid int64) error {
	if _, ok := live.ups.Load(uid); ok {
		live.ups.Delete(uid)
		return nil
	} else {
		return fmt.Errorf(ErrNotListen)
	}
}

func (live *Live) GetList() (ups [][]string) {
	live.ups.Range(func(key, value interface{}) bool {
		ups = append(ups, []string{value.(*UpRoutine).Name, fmt.Sprint(value.(*UpRoutine).Time)})
		return true
	})

	return ups
}

func (live *Live) Add(uid, _ int64, _ api.API, ctx context.Context, cancel context.CancelFunc) error {
	if _, ok := live.ups.Load(uid); ok {
		return fmt.Errorf(ErrRepeatListen)
	}

	var name string

	if s, err := api.GetUserName(uid); err == nil {
		name = s
	}

	live.ups.Store(uid, &UpRoutine{
		Ctx:    ctx,
		Cancel: cancel,
		Name:   name,
	})
	return nil
}

func NewLive() *Live {
	return &Live{
		ups: sync.Map{},
	}
}
