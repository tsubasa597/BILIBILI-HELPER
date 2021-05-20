package listen

import (
	"context"
	"sync"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
)

type Listener interface {
	Listen(int64, api.API) info.Infoer
	StopListenUP(int64) error
	GetList() string
	Add(int64, string, context.Context, context.CancelFunc) error
}

var (
	listenCtx, listenCancel               = context.WithCancel(context.Background())
	duration                time.Duration = time.Minute * 5
	mutex                   sync.RWMutex
)

type UpRoutine struct {
	Name   string
	Cancel context.CancelFunc
	Ctx    context.Context
}

func listen(tick <-chan time.Time, uid int64, listener Listener, ch chan<- info.Infoer, api api.API, ctx context.Context) {
	api.Entry.Debugf("Start : %T %d", listener, uid)
	for {
		select {
		case <-ctx.Done():
			api.Entry.Debugf("Stop : %T %d", listener, uid)
			return
		case <-tick:
			in := listener.Listen(uid, api)
			ch <- in
		}
	}
}

func StopListenUP(uid int64, listener Listener) error {
	mutex.Lock()
	defer mutex.Unlock()

	listener.StopListenUP(uid)

	return nil
}

// Stop 释放资源
func Stop() {
	listenCancel()
}

func GetListenList(listener Listener) string {
	mutex.RLock()
	defer mutex.RUnlock()

	return listener.GetList()
}

func AddListen(uid int64, listener Listener, api api.API) (context.Context, chan info.Infoer, error) {
	mutex.Lock()
	defer mutex.Unlock()

	ct, cl := context.WithCancel(listenCtx)
	name, err := api.GetUserName(uid)
	if err != nil {
		api.Entry.Warningf("错误")
	}
	if err := listener.Add(uid, name, ct, cl); err != nil {
		return nil, nil, err
	}
	ch := make(chan info.Infoer, 1)

	go listen(time.NewTicker(duration).C, uid, listener, ch, api, ct)

	return ct, ch, nil
}

func SetDuration(d time.Duration) {
	duration = d
}
