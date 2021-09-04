package listen

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

// Infoer 监听对象接口
type Infoer interface {
	Get(int64) (*UpRoutine, error)
	Put(int64, *UpRoutine) error
	GetAll() []*UpRoutine
	StopOne(uid int64) error
}

var _ Infoer = (*DefaultInfos)(nil)

// DefaultInfos 监听对象集合
type DefaultInfos struct {
	infos map[int64]*UpRoutine
	mutex *sync.RWMutex
}

// NewDefaultInfos 初始化
func NewDefaultInfos() *DefaultInfos {
	return &DefaultInfos{
		infos: make(map[int64]*UpRoutine),
		mutex: &sync.RWMutex{},
	}
}

// Get 获取指定 uid 的监听信息
func (d *DefaultInfos) Get(uid int64) (*UpRoutine, error) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	info, ok := d.infos[uid]
	if !ok {
		return nil, fmt.Errorf(errNotListen)
	}

	return info, nil
}

// Put 添加监听信息
func (d *DefaultInfos) Put(uid int64, up *UpRoutine) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if _, ok := d.infos[uid]; ok {
		return fmt.Errorf(errRepeatListen)
	}

	d.infos[uid] = up
	return nil
}

// GetAll 获取所有监听的信息
func (d *DefaultInfos) GetAll() (ups []*UpRoutine) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	for _, info := range d.infos {
		ups = append(ups, info)
	}

	return
}

// StopOne 停止监听
func (d *DefaultInfos) StopOne(uid int64) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	info, ok := d.infos[uid]
	if !ok {
		return fmt.Errorf(errNotListen)
	}
	info.Stop()
	delete(d.infos, uid)
	return nil
}

// State 运行状态
type State int32

const (
	// StatueStoped 停止状态
	StatueStoped State = iota
	// StatePause 暂停状态
	StatePause
	// StateRuning 正在运行状态
	StateRuning
)

// UpRoutine 监听信息  非线程安全
type UpRoutine struct {
	Name   string // Name 用户姓名
	Ctx    context.Context
	Time   int32
	state  *State
	cancel context.CancelFunc
}

// NewUpRoutine 初始化
func NewUpRoutine(ctx context.Context, cancel context.CancelFunc, state State, t int32, name string) *UpRoutine {
	return &UpRoutine{
		Ctx:    ctx,
		Name:   name,
		Time:   t,
		cancel: cancel,
		state:  &state,
	}
}

// GetState 获取当前状态
func (up UpRoutine) GetState() State {
	return State(atomic.LoadInt32((*int32)(up.state)))
}

// Pause 从运行状态暂停
func (up UpRoutine) Pause() bool {
	switch *up.state {
	case StatePause:
		return true
	case StateRuning:
		return atomic.CompareAndSwapInt32((*int32)(up.state), int32(StateRuning), int32(StatePause))
	default:
		return false
	}
}

// Start 从暂停状态开始
func (up UpRoutine) Start() bool {
	switch *up.state {
	case StatePause:
		return atomic.CompareAndSwapInt32((*int32)(up.state), int32(StatePause), int32(StateRuning))
	case StateRuning:
		return true
	default:
		return false
	}
}

// Stop 停止
func (up UpRoutine) Stop() bool {
	switch *up.state {
	case StatueStoped:
		return false
	default:
		up.cancel()
		atomic.StoreInt32((*int32)(up.state), int32(StatueStoped))
		return true
	}
}
