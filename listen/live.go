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

// Live 所有 uid 的状态 sync.Map -> map[int64]UpRoutine
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

// Listen 监听直播信息
func (live *Live) Listen(uid int64, _ api.API, log *logrus.Entry) []info.Infoer {
	status := getLiverStatus(uid, log)
	return []info.Infoer{&status}
}

// StopListenUP 停止监听
func (live *Live) StopListenUP(uid int64) error {
	if val, ok := live.ups.Load(uid); ok {
		live.ups.Delete(uid)
		val.(*UpRoutine).Cancel()
		return nil
	}

	return fmt.Errorf(errNotListen)
}

// GetList 所有监听的状态
func (live *Live) GetList() (ups [][2]string) {
	live.ups.Range(func(key, value interface{}) bool {
		ups = append(ups, [2]string{value.(*UpRoutine).Name, fmt.Sprint(value.(*UpRoutine).Time)})
		return true
	})

	return ups
}

// Add 添加监听
func (live *Live) Add(ctx context.Context, cancel context.CancelFunc, uid int64, _ int32, _ api.API) error {
	if _, ok := live.ups.Load(uid); ok {
		return fmt.Errorf(errRepeatListen)
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

// NewLive 初始化
func NewLive() *Live {
	return &Live{
		ups: sync.Map{},
	}
}
