package cron

import (
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api/comment"
)

// Info 任务的信息
type Info struct {
	// ID uid 或 rid
	ID int64
	// Type 动态类型
	Type comment.Type
	// TimeCell 时间间隔
	TimeCell time.Duration
	// Time 动态时间
	Time int64
	// Name 昵称
	Name string
}

// Tasker 任务接口
type Tasker interface {
	// Run(chan<- interface{}, *sync.WaitGroup)
	Run() interface{}
	Next(time.Time) time.Time
	Info() Info
}
