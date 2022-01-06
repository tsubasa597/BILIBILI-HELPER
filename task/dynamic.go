package task

import (
	"sync"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api/dynamic"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
)

type Dynamic struct {
	UID      int64
	Time     int64 // 最新动态的更新时间
	offect   int64 // 翻页参数
	timeCell time.Duration
	rwMutex  *sync.RWMutex
}

var (
	_ Tasker = (*Dynamic)(nil)
)

// NewDynamic 初始化
// 时间间隔 timeCell 的单位为 **分钟**
func NewDynamic(uid, ti int64, timeCell time.Duration) *Dynamic {
	return &Dynamic{
		UID:      uid,
		Time:     ti,
		timeCell: timeCell * time.Minute,
		rwMutex:  &sync.RWMutex{},
	}
}

// Run 获取动态
func (d *Dynamic) Run(ch chan<- interface{}, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	d.rwMutex.Lock()
	defer d.rwMutex.Unlock()

	dynamics, _ := dynamic.GetDynamics(d.UID, d.offect)
	if len(dynamics) > 0 {
		d.Time = dynamics[0].Time
		d.offect = dynamics[len(dynamics)-1].Offect
	}

	ch <- dynamics
}

// Next 下次运行时间
func (d Dynamic) Next(t time.Time) time.Time {
	d.rwMutex.Lock()
	defer d.rwMutex.Unlock()

	return t.Add(d.timeCell)
}

// Info 返回任务的信息
func (d Dynamic) Info() info.Task {
	d.rwMutex.RLock()
	defer d.rwMutex.RUnlock()

	return info.Task{
		ID:       d.UID,
		Time:     d.Time,
		TimeCell: d.timeCell,
	}
}
