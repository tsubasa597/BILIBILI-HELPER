package task

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api/dynamic"
	"github.com/tsubasa597/BILIBILI-HELPER/state"
)

type Dynamic struct {
	UID      int64
	Time     int64
	offect   int64
	timeCell time.Duration
	state    state.State
}

var (
	_ Tasker = (*Dynamic)(nil)
)

// NewDynamic 初始化
func NewDynamic(uid, ti int64, t time.Duration) *Dynamic {
	return &Dynamic{
		UID:      uid,
		Time:     ti,
		timeCell: t,
		state:    state.Stop,
	}
}

// Run 获取动态
func (d *Dynamic) Run(ch chan<- interface{}, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	if atomic.LoadInt32((*int32)(&d.state)) != int32(state.Stop) {
		return
	}

	atomic.SwapInt32((*int32)(&d.state), int32(state.Runing))
	defer func() {
		atomic.SwapInt32((*int32)(&d.state), int32(state.Stop))
	}()

	dynamics, _ := dynamic.GetDynamics(d.UID, d.offect)
	if len(dynamics) > 0 {
		d.Time = dynamics[0].Time
		d.offect = dynamics[len(dynamics)-1].Offect
	}

	ch <- dynamics
}

// Next 下次运行时间
func (d Dynamic) Next(t time.Time) time.Time {
	return t.Add(time.Minute * d.timeCell)
}
