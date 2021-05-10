package api

import (
	"context"
	"fmt"
	"time"
)

type Listen struct {
	Tick   *time.Ticker
	ctx    context.Context
	cancel context.CancelFunc
	ups    map[int64]upRoutine
	api    API
}

type upRoutine struct {
	cancel context.CancelFunc
	ctx    context.Context
}

func (l *Listen) listen(uid int64, f func(int64) Info) (context.Context, chan Info, error) {
	if _, ok := l.ups[uid]; ok {
		l.api.entry.Debugln(uid, errRepeatListen)
		return l.ups[uid].ctx, nil, fmt.Errorf(errRepeatListen)
	}

	ct, cl := context.WithCancel(l.ctx)
	l.ups[uid] = upRoutine{
		cancel: cl,
		ctx:    ct,
	}
	ch := make(chan Info, 1)

	go func() {
		l.api.entry.Debugf("Start : %T %d", f, uid)
		for {
			select {
			case <-l.ctx.Done():
				l.api.entry.Debugf("Stop : %T %d", f, uid)
				return
			case <-l.Tick.C:
				ch <- f(uid)
			}
		}
	}()

	return ct, ch, nil
}

func (d *Listen) StopListenUP(uid int64) error {
	if _, ok := d.ups[uid]; !ok {
		return fmt.Errorf(errNotListen)
	}

	d.ups[uid].cancel()
	return nil
}

// Stop 释放资源
func (d *Listen) Stop() {
	d.api.entry.Infoln("All Goroutine Quit")
	d.cancel()
}

func NewListen(api API) *Listen {
	c, cl := context.WithCancel(context.Background())
	return &Listen{
		Tick:   time.NewTicker(time.Minute * 5),
		ctx:    c,
		cancel: cl,
		ups:    make(map[int64]upRoutine),
		api:    api,
	}
}
