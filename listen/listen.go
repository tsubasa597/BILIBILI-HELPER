// Package listen 固定时间间隔监听 bilibili 用户的动态或直播情况
package listen

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
)

const (
	//NewListen 新监听动态的时间，防止返回大量数据
	NewListen = iota + 1
)

// Listener 所需要监听的信息的接口
type Listener interface {
	Listen(int64) []info.Infoer
	StopListenUP(int64) error
	GetList() []*UpRoutine
	Add(context.Context, context.CancelFunc, int64, int32) error
}

// Listen 管理监听状态
type Listen struct {
	listener Listener
	ctx      context.Context
	cancel   context.CancelFunc
	api      *api.API
	log      *logrus.Entry
}

// listen 以固定时间间隔进行监听
func (listen *Listen) listen(ctx context.Context, uid int64, ticker *time.Ticker, ch chan<- []info.Infoer) {
	listen.log.Debugf("Start : %T %d", listen.listener, uid)
	for {
		select {
		case <-ctx.Done():
			listen.log.Debugf("Stop : %T %d", listen.listener, uid)
			ticker.Stop()
			return
		case <-ticker.C:
			ch <- listen.listener.Listen(uid)
		}
	}
}

// StopUP 停止监听
func (listen *Listen) StopUP(uid int64, listener Listener) error {
	return listen.listener.StopListenUP(uid)
}

// Stop 释放资源
func (listen *Listen) Stop() {
	listen.cancel()
}

// GetList 返回监听状态
func (listen *Listen) GetList() []*UpRoutine {
	return listen.listener.GetList()
}

// Add 添加 uid 进行监听
func (listen *Listen) Add(uid int64, t int32, duration time.Duration) (context.Context, chan []info.Infoer, error) {
	ctx, cl := context.WithCancel(listen.ctx)
	if err := listen.listener.Add(ctx, cl, uid, t); err != nil {
		return nil, nil, err
	}
	ch := make(chan []info.Infoer, 1)

	go listen.listen(ctx, uid, time.NewTicker(duration), ch)

	return ctx, ch, nil
}

// New 初始化监控
func New(linster Listener, api *api.API, entry *logrus.Entry) (*Listen, context.Context) {
	if entry == nil {
		entry = logrus.NewEntry(logrus.New())
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Listen{
		listener: linster,
		ctx:      ctx,
		cancel:   cancel,
		api:      api,
		log:      entry,
	}, ctx
}
