package api

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Listen struct {
	ctx    context.Context
	cancel context.CancelFunc
	ups    map[int64]*upRoutine
	api    API
	mutex  sync.RWMutex
}

type upRoutine struct {
	name   string
	cancel context.CancelFunc
	ctx    context.Context
}

func (l *Listen) listen(tick <-chan time.Time, uid int64, f func(int64) Info) (context.Context, chan Info, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if v, ok := l.ups[uid]; ok {
		l.api.entry.Debugln(uid, errRepeatListen)
		return v.ctx, nil, fmt.Errorf(errRepeatListen)
	}

	ct, cl := context.WithCancel(l.ctx)
	l.ups[uid] = &upRoutine{
		cancel: cl,
		ctx:    ct,
	}
	ch := make(chan Info, 1)

	go func(tick <-chan time.Time, uid int64, f func(int64) Info) {
		l.api.entry.Debugf("Start : %T %d", f, uid)
		for {
			select {
			case <-l.ctx.Done():
				l.api.entry.Debugf("Stop : %T %d", f, uid)
				return
			case <-tick:
				in := f(uid)
				if v, ok := l.ups[uid]; ok {
					v.name = in.Name
				}
				ch <- in
			}
		}
	}(tick, uid, f)

	return ct, ch, nil
}

func (l *Listen) StopListenUP(uid int64) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if v, ok := l.ups[uid]; !ok {
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
		ctx:    c,
		cancel: cl,
		ups:    make(map[int64]*upRoutine),
		api:    api,
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
