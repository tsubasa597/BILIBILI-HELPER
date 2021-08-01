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
	listenCtx    context.Context
	listenCancel context.CancelFunc
	api          api.API
	log          *logrus.Entry
}

type UpRoutine struct {
	Name   string
	Cancel context.CancelFunc
	Ctx    context.Context
	Time   int32
}

func (listen *Listen) listen(ctx context.Context, uid int64, listener Listener, tick <-chan time.Time, ch chan<- []info.Infoer) {
	listen.log.Debugf("Start : %T %d", listener, uid)
	for {
		select {
		case <-ctx.Done():
			listen.log.Debugf("Stop : %T %d", listener, uid)
			return
		case <-tick:
			ch <- listener.Listen(uid, listen.api, listen.log)
		}
	}
}

func StopUP(uid int64, listener Listener) error {
	listener.StopListenUP(uid)

	return nil
}

// Stop 释放资源
func (listen *Listen) Stop() {
	listen.listenCancel()
}

func GetList(listener Listener) [][]string {
	return listener.GetList()
}

func (listen *Listen) Add(uid int64, t int32, listener Listener, duration time.Duration) (context.Context, chan []info.Infoer, error) {
	ct, cl := context.WithCancel(listen.listenCtx)
	if err := listener.Add(uid, t, listen.api, ct, cl); err != nil {
		return nil, nil, err
	}
	ch := make(chan []info.Infoer, 1)

	go listen.listen(ct, uid, listener, time.NewTicker(duration).C, ch)

	return ct, ch, nil
}

func New(api api.API, entry *logrus.Entry) (*Listen, context.Context) {
	if entry == nil {
		entry = logrus.NewEntry(log.NewLog())
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Listen{
		listenCtx:    ctx,
		listenCancel: cancel,
		api:          api,
		log:          entry,
	}, ctx
}
