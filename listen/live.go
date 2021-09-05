package listen

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"github.com/tsubasa597/BILIBILI-HELPER/state"
)

// Live 所有 uid 的状态
type Live struct {
	ups state.ListenStater
}

func getLiverStatus(uid int64, state state.Listener, log *logrus.Entry) (infos []info.Interface) {
	rep, err := api.GetUserInfo(uid)
	if err != nil {
		log.Errorln(err)
		return
	}

	live := info.Live{}
	live.Name = rep.Data.Name
	live.Time = int32(time.Now().Unix())
	live.LiveRoomURL = rep.Data.LiveRoom.Url
	live.LiveTitle = rep.Data.LiveRoom.Title
	if rep.Data.LiveRoom.RoomStatus == 1 {
		live.LiveStatus = true
	}

	return append(infos, &live)
}

var _ Listener = (*Live)(nil)

// ListenInfo 监听直播信息
func (live *Live) ListenInfo(uid int64) ([]info.Interface, error) {
	return live.ups.Do(uid, getLiverStatus)
}

// StopListenUP 停止监听
func (live *Live) StopListenUP(uid int64) error {
	return live.ups.StopOne(uid)
}

// GetList 所有监听的状态
func (live *Live) GetList() []state.Info {
	return live.ups.GetAll()
}

// GetState 获取状态
func (live *Live) GetState() info.State {
	return live.ups.GetState()
}

// Add 添加监听
func (live *Live) Add(ctx context.Context, cancel context.CancelFunc, uid int64, _ int32) error {
	var name string

	if s, err := api.GetUserName(uid); err == nil {
		name = s
	}

	return live.ups.Put(uid, state.NewUpRoutine(ctx, cancel, 0, name))
}

// NewLive 初始化
func NewLive(log *logrus.Entry) *Live {
	return &Live{
		ups: state.NewDeListenState(log),
	}
}
