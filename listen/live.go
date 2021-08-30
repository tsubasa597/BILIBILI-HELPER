package listen

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
)

// Live 所有 uid 的状态
type Live struct {
	ups Infoer
	log *logrus.Entry
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
func (live *Live) Listen(uid int64) []info.Infoer {
	status := getLiverStatus(uid, live.log)
	return []info.Infoer{&status}
}

// StopListenUP 停止监听
func (live *Live) StopListenUP(uid int64) error {
	return live.ups.StopOne(uid)
}

// GetList 所有监听的状态
func (live *Live) GetList() []*UpRoutine {
	return live.ups.GetAll()
}

// Add 添加监听
func (live *Live) Add(ctx context.Context, cancel context.CancelFunc, uid int64, _ int32) error {
	if _, err := live.ups.Get(uid); err != nil {
		return err
	}

	var name string

	if s, err := api.GetUserName(uid); err == nil {
		name = s
	}

	live.ups.Put(uid, NewUpRoutine(ctx, cancel, StateRuning, 0, name))
	return nil
}

// NewLive 初始化
func NewLive(log *logrus.Entry) *Live {
	return &Live{
		ups: NewDefaultInfos(),
		log: log,
	}
}
