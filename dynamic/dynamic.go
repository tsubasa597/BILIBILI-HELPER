package dynamic

import (
	"context"
	"fmt"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
)

type DynamicListen struct {
	T             *time.Ticker
	ctx           context.Context
	cancel        context.CancelFunc
	upsCancelFunc map[int64]context.CancelFunc
	upsContext    map[int64]context.Context
}

func (d *DynamicListen) ListenDynamic(uid int64) (context.Context, chan api.Info, error) {
	if _, ok := d.upsCancelFunc[uid]; ok {
		return d.upsContext[uid], nil, fmt.Errorf("重复监听")
	}

	ct, cl := context.WithCancel(d.ctx)
	d.upsContext[uid] = ct
	d.upsCancelFunc[uid] = cl
	c := make(chan api.Info, 1)
	go listen(*d.T, uid, ct, c)

	return ct, c, nil
}

func (d *DynamicListen) StopListenUP(uid int64) error {
	if _, ok := d.upsCancelFunc[uid]; !ok {
		return fmt.Errorf("该用户未监听")
	}

	d.upsCancelFunc[uid]()
	return nil
}

// Stop 释放资源
func (d DynamicListen) Stop() {
	d.cancel()
}

func listen(ticker time.Ticker, uid int64, ctx context.Context, ch chan<- api.Info) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ch <- api.GetDynamicMessage(uid)
		}
	}
}

func New() *DynamicListen {
	c, cl := context.WithCancel(context.Background())
	return &DynamicListen{
		T:             time.NewTicker(time.Minute),
		ctx:           c,
		cancel:        cl,
		upsCancelFunc: make(map[int64]context.CancelFunc),
		upsContext:    make(map[int64]context.Context),
	}
}
