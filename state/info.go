package state

import "context"

// State 运行状态
type State int32

const (
	// Stop 停止状态
	Stop State = iota
	// Pause 暂停状态
	Pause
	// Runing 正在运行状态
	Runing
)

// Info 获取的状态信息
type Info struct {
	Name string
	Time int32
	Ctx  context.Context
}
