package cron

import (
	"sync"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/api/live"
	"github.com/tsubasa597/BILIBILI-HELPER/api/proto"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

const (
	_waitLive = 60 /* 未开播时，监控直播间间隔时间 */
)

type Live struct {
	Name string
	Uid  int64

	room      *live.Room
	enterData *proto.EnterLive_Data
	status    atomic.Int32
	timeCell  time.Duration
	api       api.API
	log       *zap.Logger
}

var (
	_ Tasker = (*Live)(nil)
)

func NewLive(ap api.API, uid int64, log *zap.Logger) (*Live, error) {
	return &Live{
		Uid:      uid,
		status:   *atomic.NewInt32(int32(proto.LiveStatus_NoLiving)),
		timeCell: time.Duration(_waitLive * time.Second),
		api:      ap,
		log:      log,
	}, nil
}

func (l *Live) Run(ch chan<- interface{}, wg *sync.WaitGroup) {
	if l.status.Load() != int32(proto.LiveStatus_NoLiving) {
		ch <- l.room.In(l.enterData)
		return
	}

	info, err := live.Status(l.Uid)
	if err != nil {
		l.log.Error(err.Error())
		return
	} else if info.Status == proto.LiveStatus_NoLiving {
		return
	}

	l.room, err = live.NewRoom(l.api, info)
	if err != nil {
		l.log.Error(err.Error())
		return
	}

	l.status.Swap(int32(info.Status))
	data, err := l.room.Enter()
	if err != nil {
		l.log.Error(err.Error())
		return
	}

	l.enterData = data
	l.timeCell = time.Duration(data.HeartbeatInterval)
}

func (live Live) Next(t time.Time) time.Time {
	return t.Add(live.timeCell)
}
func (live Live) Info() Info {
	return Info{
		Name: live.Name,
	}
}
