package api

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Listen struct {
	Duration time.Duration
	ctx      context.Context
	cancel   context.CancelFunc
	ups      map[int64]*upRoutine
	api      API
	mutex    sync.RWMutex
}

type upRoutine struct {
	name   string
	cancel context.CancelFunc
	ctx    context.Context
}

func (listen *Listen) listen(tick <-chan time.Time, uid int64, f func(int64) Info, ch chan<- Info) {
	listen.api.entry.Debugf("Start : %T %d", f, uid)
	for {
		select {
		case <-listen.ctx.Done():
			listen.api.entry.Debugf("Stop : %T %d", f, uid)
			return
		case <-tick:
			in := f(uid)

			listen.mutex.Lock()
			if v, ok := listen.ups[uid]; ok && in.Name != "" {
				v.name = in.Name
			}
			listen.mutex.Unlock()

			ch <- in
		}
	}
}

func (listen *Listen) StopListenUP(uid int64) error {
	listen.mutex.Lock()
	defer listen.mutex.Unlock()

	if v, ok := listen.ups[uid]; !ok {
		return fmt.Errorf(errNotListen)
	} else {
		v.cancel()
	}

	return nil
}

// Stop 释放资源
func (l *Listen) Stop() {
	l.api.entry.Infoln("All Goroutine Quit")
	l.cancel()
}

func NewListen(api API) *Listen {
	c, cl := context.WithCancel(context.Background())
	return &Listen{
		Duration: time.Minute * 5,
		ctx:      c,
		cancel:   cl,
		ups:      make(map[int64]*upRoutine),
		api:      api,
	}
}

func (l *Listen) GetListenList() (ups []string) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	for _, v := range l.ups {
		ups = append(ups, v.name)
	}

	return
}

func (listen *Listen) AddListen(uid int64, f func(int64) Info) (context.Context, chan Info, error) {
	listen.mutex.Lock()
	defer listen.mutex.Unlock()

	if v, ok := listen.ups[uid]; ok {
		listen.api.entry.Debugln(uid, errRepeatListen)
		return v.ctx, nil, fmt.Errorf(errRepeatListen)
	}

	ct, cl := context.WithCancel(listen.ctx)
	listen.ups[uid] = &upRoutine{
		cancel: cl,
		ctx:    ct,
	}
	ch := make(chan Info, 1)

	go listen.listen(time.NewTicker(listen.Duration).C, uid, f, ch)
	return ct, ch, nil
}
