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

// Run 获取动态
func (d *Dynamic) Run(ch chan<- interface{}) {
	if d.state == state.Stop {
		return
	}

	dynamics := api.GetAllDynamics(d.UID, d.Time)
	if len(dynamics) > 0 {
		d.Time = dynamics[0].Time
	}

	ch <- dynamics
}

// State 获取运行状态
func (d Dynamic) State() state.State {
	return d.state
}

// Next 下次运行时间
func (d Dynamic) Next(t time.Time) time.Time {
	return t.Add(time.Minute * d.timeCell)
}

// NewDynamic 初始化
func NewDynamic(uid, ti int64, t time.Duration) *Dynamic {
	return &Dynamic{
		UID:      uid,
		Time:     ti,
		timeCell: t,
		state:    state.Runing,
	}
}
