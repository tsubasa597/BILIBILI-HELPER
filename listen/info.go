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
	StopOne(int64) error
	GetAll() []*UpRoutine
	Pause(int64)
	Start(int64)
}

var _ Infoer = (*DefaultInfos)(nil)

// DefaultInfos 默认监听对象集合
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

// StopOne 停止监听已监听用户
func (d *DefaultInfos) StopOne(uid int64) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	up, ok := d.infos[uid]
	if !ok {
		return fmt.Errorf(errNotListen)
	}

	up.Stop()
	delete(d.infos, uid)

	return nil
}

// GetAll 获取所有正在监听的信息
func (d *DefaultInfos) GetAll() (ups []*UpRoutine) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	for _, info := range d.infos {
		if info.GetState() != StateRuning {
			continue
		}

		ups = append(ups, info)
	}

	return
}

// Start 从暂停状态恢复
func (d DefaultInfos) Start(uid int64) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	if up, ok := d.infos[uid]; ok {
		up.Start()
	}
}

// Pause 暂停监听
func (d DefaultInfos) Pause(uid int64) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	if up, ok := d.infos[uid]; ok {
		up.Pause()
	}
}

// UpRoutine 监听信息
type UpRoutine struct {
	Name   string // Name 用户姓名
	Ctx    context.Context
	Time   int32
	state  *int32
	cancel context.CancelFunc
}

// NewUpRoutine 初始化
func NewUpRoutine(ctx context.Context, cancel context.CancelFunc, state int32, t int32, name string) *UpRoutine {
	return &UpRoutine{
		Ctx:    ctx,
		Name:   name,
		Time:   t,
		cancel: cancel,
		state:  &state,
	}
}

// GetState 获取当前状态
func (up UpRoutine) GetState() int32 {
	return atomic.LoadInt32(up.state)
}

// Pause 暂停
func (up UpRoutine) Pause() {
	atomic.SwapInt32(up.state, StatePause)
}

// Start 开始
func (up UpRoutine) Start() {
	atomic.SwapInt32(up.state, StateRuning)
}

// Stop 停止
func (up UpRoutine) Stop() {
	up.cancel()
}

const (
	// StatePause 暂停状态
	StatePause int32 = iota
	// StateRuning 正在运行状态
	StateRuning
)
