package listen

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"github.com/tsubasa597/BILIBILI-HELPER/log"
	"github.com/tsubasa597/requests"
)

type Listener interface {
	Listen(int64, api.API) []info.Infoer
	StopListenUP(int64) error
	GetList() [][]string
	Add(int64, int64, api.API, context.Context, context.CancelFunc) error
}

type State int64

const (
	NewListen State = iota + 1
)

var (
	listenCtx, listenCancel               = context.WithCancel(context.Background())
	duration                time.Duration = time.Minute * 5
	defaultAPI              api.API       = api.API{
		Entry: logrus.NewEntry(log.NewLog()),
		Req: &requests.Requests{
			Client: &http.Client{},
		},
	}
)

type UpRoutine struct {
	Name   string
	Cancel context.CancelFunc
	Ctx    context.Context
	Time   int64
}

func listen(ctx context.Context, uid int64, listener Listener, tick <-chan time.Time, ch chan<- []info.Infoer) {
	defaultAPI.Entry.Debugf("Start : %T %d", listener, uid)
	for {
		select {
		case <-ctx.Done():
			defaultAPI.Entry.Debugf("Stop : %T %d", listener, uid)
			return
		case <-tick:
			ch <- listener.Listen(uid, defaultAPI)
		}
	}
}

func StopUP(uid int64, listener Listener) error {
	listener.StopListenUP(uid)

	return nil
}

// Stop 释放资源
func Stop() {
	listenCancel()
}

func GetList(listener Listener) [][]string {
	return listener.GetList()
}

func Add(uid, t int64, listener Listener) (context.Context, chan []info.Infoer, error) {
	ct, cl := context.WithCancel(listenCtx)
	if err := listener.Add(uid, t, defaultAPI, ct, cl); err != nil {
		return nil, nil, err
	}
	ch := make(chan []info.Infoer, 1)

	go listen(ct, uid, listener, time.NewTicker(duration).C, ch)

	return ct, ch, nil
}

func SetDuration(d time.Duration) {
	duration = d
}

func SetAPI(api api.API) {
	defaultAPI = api
}
