package info

import "time"

type Task struct {
	// ID uid 或 rid
	ID int64
	// Type 动态类型
	Type Type
	// TimeCell 时间间隔
	TimeCell time.Duration
	// Time 动态时间
	Time int64
}
