package cron

import (
	"fmt"
	"sync"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api/dynamic"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Dynamic 动态信息
type Dynamic struct {
	UID  int64
	Name string

	// 最新动态的更新时间
	time int64

	timeCell time.Duration
	log      *zap.Logger
	mutex    *sync.Mutex
}

const (
	_errDKey = "UID 信息"
)

var (
	_ Tasker = (*Dynamic)(nil)
)

// NewDynamic 初始化
func NewDynamic(uid, ti int64, timeCell time.Duration, log *zap.Logger) *Dynamic {
	if log == nil {
		log = zap.NewExample()
	}

	return &Dynamic{
		UID:      uid,
		time:     ti,
		timeCell: timeCell,
		log:      log,
		mutex:    &sync.Mutex{},
	}
}

// Run 监听最新动态
func (d *Dynamic) Run() interface{} {
	d.mutex.Lock()

	dynamics, errs := dynamic.GetAll(d.UID, d.time)
	if len(dynamics) > 0 {
		d.Name = dynamics[0].Name
		d.time = dynamics[0].Time
	}

	for _, err := range errs {
		d.log.Error(err.Error(), zapcore.Field{
			Key:    _errDKey,
			Type:   zapcore.StringType,
			String: fmt.Sprintf("UID: %d", d.UID),
		})
	}

	d.mutex.Unlock()

	return dynamics
}

// Next 下次运行时间
func (d Dynamic) Next(t time.Time) time.Time {
	return t.Add(d.timeCell)
}

// Info 返回任务的信息
func (d Dynamic) Info() Info {
	return Info{
		ID:       d.UID,
		Name:     d.Name,
		Time:     d.time,
		TimeCell: d.timeCell,
	}
}
