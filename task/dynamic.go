package task

import (
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/state"
)

type Dynamic struct {
	UID      int64
	Time     int64
	timeCell time.Duration
	state    state.State
}

var (
	_ Tasker = (*Dynamic)(nil)
)

func (d *Dynamic) Run(ch chan<- interface{}) {
	if d.state != state.Runing {
		return
	}

	dynamics := api.GetAllDynamics(d.UID, d.Time)
	if len(dynamics) > 0 {
		d.Time = dynamics[0].Time
	}

	ch <- dynamics
}

func (d Dynamic) Next(t time.Time) time.Time {
	return t.Add(time.Minute * d.timeCell)
}

func NewDynamic(uid, ti int64, t time.Duration) *Dynamic {
	return &Dynamic{
		UID:      uid,
		Time:     ti,
		timeCell: t,
		state:    state.Runing,
	}
}
