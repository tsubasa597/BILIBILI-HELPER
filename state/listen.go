package state

import (
	"context"
)

// Listener 监听信息接口
type Listener interface {
	GetInfo() Info
	GetState() State
	Update(interface{})
	Pause() bool
	Start() bool
	Stop()
}

var _ Listener = (*UpRoutine)(nil)

// UpRoutine 监听信息  !!非线程安全!!
type UpRoutine struct {
	Name   string // Name 用户姓名
	Ctx    context.Context
	Time   int32
	state  State
	cancel context.CancelFunc
}

// NewUpRoutine 初始化
func NewUpRoutine(ctx context.Context, cancel context.CancelFunc, t int32, name string) *UpRoutine {
	return &UpRoutine{
		Ctx:    ctx,
		Name:   name,
		Time:   t,
		cancel: cancel,
		state:  Runing,
	}
}

// GetInfo 获取监听信息
func (up UpRoutine) GetInfo() Info {
	return Info{
		Name: up.Name,
		Time: up.Time,
		Ctx:  up.Ctx,
	}
}

// GetState 获取当前状态
func (up UpRoutine) GetState() State {
	return up.state
}

// Pause 从运行状态暂停
func (up UpRoutine) Pause() bool {
	switch up.state {
	case Pause:
		return true
	case Runing:
		up.state = Pause
		return true
	default:
		return false
	}
}

// Start 从暂停状态开始
func (up UpRoutine) Start() bool {
	switch up.state {
	case Pause:
		up.state = Runing
		return true
	case Runing:
		return true
	default:
		return false
	}
}

// Stop 停止
func (up UpRoutine) Stop() {
	up.cancel()
}

// Update 更新时间
func (up *UpRoutine) Update(t interface{}) {
	up.Time = t.(int32)
}
