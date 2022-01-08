package info

import "time"

// Task 任务的信息
type Task struct {
	// ID uid 或 rid
	ID int64
	// Type 动态类型
	Type DynamicType
	// TimeCell 时间间隔
	TimeCell time.Duration
	// Time 动态时间
	Time int64
}
