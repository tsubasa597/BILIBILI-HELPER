// Package listen 固定时间间隔监听 bilibili 用户的动态或直播情况
package listen

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"github.com/tsubasa597/BILIBILI-HELPER/state"
)

// Listener 所需要监听的信息的接口
type Listener interface {
	GetState() state.State
	ListenInfo(int64) ([]info.Interface, error)
	StopListenUP(int64) error
	GetList() []state.Info
	Add(context.Context, context.CancelFunc, int64, int32) error
}

// Listen 管理监听状态
type Listen struct {
	listener Listener
	Ctx      context.Context
	cancel   context.CancelFunc
	log      *logrus.Entry
}

// listen 以固定时间间隔进行监听
func (listen *Listen) listen(ctx context.Context, uid int64, ticker *time.Ticker, ch chan<- []info.Interface) {
	listen.log.Debugf("Start : %T %d", listen.listener, uid)
	for {
		select {
		case <-ctx.Done():
			listen.log.Debugf("Stop : %T %d", listen.listener, uid)

			close(ch)
			ticker.Stop()

			return
		case <-ticker.C:
			listen.log.Debugf("Get Info From : %T %d", listen.listener, uid)

			infos, err := listen.listener.ListenInfo(uid)
			if err != nil {
				listen.log.Debug(err)
				continue
			}

			ch <- infos
		}
	}
}

// StopUP 停止监听
func (listen *Listen) StopUP(uid int64) error {
	return listen.listener.StopListenUP(uid)
}

// Stop 释放资源
func (listen *Listen) Stop() {
	listen.cancel()
}

// GetList 返回监听状态
func (listen *Listen) GetList() []state.Info {
	return listen.listener.GetList()
}

// Add 添加 id 进行监听
func (listen *Listen) Add(id int64, t int32, duration time.Duration) (context.Context, chan []info.Interface, error) {
	ctx, cl := context.WithCancel(listen.Ctx)
	if err := listen.listener.Add(ctx, cl, id, t); err != nil {
		return nil, nil, err
	}
	ch := make(chan []info.Interface, 1)

	go listen.listen(ctx, id, time.NewTicker(duration), ch)

	return ctx, ch, nil
}

// GetState 获取状态
func (listen Listen) GetState() state.State {
	return listen.listener.GetState()
}

// New 初始化监控
func New(ctx context.Context, linster Listener, entry *logrus.Entry) *Listen {
	if entry == nil {
		entry = logrus.NewEntry(logrus.New())
	}

	ctx, cancel := context.WithCancel(ctx)
	return &Listen{
		Ctx:      ctx,
		listener: linster,
		cancel:   cancel,
		log:      entry,
	}
}
