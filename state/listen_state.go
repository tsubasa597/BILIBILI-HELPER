package state

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/BILIBILI-HELPER/e"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
)

// ListenStater 监听对象集合接口
type ListenStater interface {
	GetState() State
	Get(int64) (Listener, error)
	Put(int64, Listener) error
	GetAll() []Info
	StopOne(int64) error
	PauseOne(int64) bool
	Pause(duration int) bool
	StartOne(int64) bool
	Start() bool
	Do(int64, func(int64, Listener, *logrus.Entry) []info.Interface) ([]info.Interface, error)
}

var _ ListenStater = (*DeListenState)(nil)

// DeListenState 监听对象集合
type DeListenState struct {
	Ctx   context.Context
	state State
	infos map[int64]Listener
	log   *logrus.Entry
	mutex *sync.RWMutex
}

// NewDeListenState 初始化
func NewDeListenState(ctx context.Context, log *logrus.Entry) *DeListenState {
	return &DeListenState{
		Ctx:   ctx,
		state: Runing,
		infos: make(map[int64]Listener),
		log:   log,
		mutex: &sync.RWMutex{},
	}
}

// GetState 获取状态
func (d *DeListenState) GetState() State {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	return d.state
}

// Get 获取指定 uid 的监听信息
func (d *DeListenState) Get(id int64) (Listener, error) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	info, ok := d.infos[id]
	if !ok {
		return nil, fmt.Errorf(e.ErrNotListen)
	}

	return info, nil
}

// Put 添加监听信息
func (d *DeListenState) Put(id int64, up Listener) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if _, ok := d.infos[id]; ok {
		return fmt.Errorf(e.ErrRepeatListen)
	}

	d.infos[id] = up
	return nil
}

// GetAll 获取所有监听的信息
func (d *DeListenState) GetAll() (ups []Info) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	for _, info := range d.infos {
		ups = append(ups, info.GetInfo())
	}

	return
}

// StopOne 停止监听
func (d *DeListenState) StopOne(id int64) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	info, ok := d.infos[id]
	if !ok {
		return fmt.Errorf(e.ErrNotListen)
	}

	info.Stop()
	delete(d.infos, id)
	return nil
}

// Do 获取信息
func (d *DeListenState) Do(id int64, f func(int64, Listener, *logrus.Entry) []info.Interface) ([]info.Interface, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	switch d.state {
	case Pause:
		return nil, fmt.Errorf(e.ErrPause)
	case Stop:
		return nil, fmt.Errorf(e.ErrorStop)
	}

	inf, ok := d.infos[id]
	if !ok {
		return nil, fmt.Errorf(e.ErrNotListen)
	}

	switch inf.GetState() {
	case Pause:
		return nil, fmt.Errorf(e.ErrPause)
	case Stop:
		delete(d.infos, id)
		return nil, fmt.Errorf(e.ErrorStop)
	}

	infos := f(id, inf, d.log)
	if len(infos) > 0 {
		inf.Update(infos[0].GetTime())
	}

	return infos, nil
}

// PauseOne 暂停
func (d *DeListenState) PauseOne(id int64) bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	inf, ok := d.infos[id]
	if !ok {
		d.log.Error(e.ErrNotListen)
		return false
	}

	return inf.Pause()
}

// Pause 暂停所有，指定时间之后恢复
func (d *DeListenState) Pause(duration int) bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.state = Pause
	for _, inf := range d.infos {
		if !inf.Pause() {
			d.log.Error(e.ErrorStop)
		}
	}

	go func(ticker *time.Ticker) {
		for {
			select {
			case <-d.Ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				d.state = Runing
				for _, inf := range d.infos {
					if !inf.Start() {
						d.log.Error(e.ErrorStop)
					}
				}
			}
		}
	}(time.NewTicker(time.Minute * time.Duration(duration)))

	d.log.Debug("暂停监听")
	return true
}

// StartOne 从暂停状态开始
func (d *DeListenState) StartOne(id int64) bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	inf, ok := d.infos[id]
	if !ok {
		d.log.Error(e.ErrNotListen)
		return false
	}

	return inf.Start()
}

// Start 从暂停状态开始
func (d *DeListenState) Start() bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.state = Runing
	for _, inf := range d.infos {
		if !inf.Start() {
			d.log.Error(e.ErrorStop)
		}
	}

	return true
}
