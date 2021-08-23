package listen

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
	log "github.com/tsubasa597/BILIBILI-HELPER/log"
)

const (
	NewListen = iota + 1
)

type Listener interface {
	Listen(int64, api.API, *logrus.Entry) []info.Infoer
	StopListenUP(int64) error
	GetList() [][]string
	Add(int64, int32, api.API, context.Context, context.CancelFunc) error
}

type Listen struct {
	Listener
	ctx    context.Context
	cancel context.CancelFunc
	api    *api.API
	log    *logrus.Entry
}

type UpRoutine struct {
	Name   string
	Cancel context.CancelFunc
	Ctx    context.Context
	Time   int32
}

func (listen *Listen) listen(ctx context.Context, uid int64, ticker *time.Ticker, ch chan<- []info.Infoer) {
	listen.log.Debugf("Start : %T %d", listen.Listener, uid)
	for {
		select {
		case <-ctx.Done():
			listen.log.Debugf("Stop : %T %d", listen.Listener, uid)
			ticker.Stop()
			return
		case <-ticker.C:
			ch <- listen.Listener.Listen(uid, *listen.api, listen.log)
		}
	}
}

func (listen *Listen) StopUP(uid int64, listener Listener) error {
	return listen.Listener.StopListenUP(uid)
}

// Stop 释放资源
func (listen *Listen) Stop() {
	listen.cancel()
}

func (listen *Listen) GetList() [][]string {
	return listen.Listener.GetList()
}

func (listen *Listen) Add(uid int64, t int32, duration time.Duration) (context.Context, chan []info.Infoer, error) {
	ct, cl := context.WithCancel(listen.ctx)
	if err := listen.Listener.Add(uid, t, *listen.api, ct, cl); err != nil {
		return nil, nil, err
	}
	ch := make(chan []info.Infoer, 1)

	go listen.listen(ct, uid, time.NewTicker(duration), ch)

	return ct, ch, nil
}

func New(linster Listener, api *api.API, entry *logrus.Entry) (*Listen, context.Context) {
	if entry == nil {
		entry = logrus.NewEntry(log.NewLog())
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Listen{
		Listener: linster,
		ctx:      ctx,
		cancel:   cancel,
		api:      api,
		log:      entry,
	}, ctx
}
