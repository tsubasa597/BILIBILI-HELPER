package dynamic

import (
	"context"
	"fmt"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
)

var (
	errNotLinsten = fmt.Errorf("该用户未监听")
)

type upRoutine struct {
	cancel context.CancelFunc
	ctx    context.Context
}

type DynamicListen struct {
	T      *time.Ticker
	ctx    context.Context
	cancel context.CancelFunc
	ups    map[int64]upRoutine
	api    api.API
}

func (d *DynamicListen) Listen(uid int64) (context.Context, chan api.Info, error) {
	if _, ok := d.ups[uid]; ok {
		return d.ups[uid].ctx, nil, fmt.Errorf("重复监听")
	}

	ct, cl := context.WithCancel(d.ctx)
	d.ups[uid] = upRoutine{
		cancel: cl,
		ctx:    ct,
	}

	c := make(chan api.Info, 1)
	go listen(d.api, *d.T, uid, ct, c)

	return ct, c, nil
}

func (d *DynamicListen) StopListenUP(uid int64) error {
	if _, ok := d.ups[uid]; !ok {
		return errNotLinsten
	}

	d.ups[uid].cancel()
	return nil
}

// Stop 释放资源
func (d DynamicListen) Stop() {
	d.cancel()
}

func listen(api api.API, ticker time.Ticker, uid int64, ctx context.Context, ch chan<- api.Info) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ch <- api.GetDynamicMessage(uid)
		}
	}
}

func New(api api.API) *DynamicListen {
	c, cl := context.WithCancel(context.Background())
	return &DynamicListen{
		T:      time.NewTicker(time.Minute),
		ctx:    c,
		cancel: cl,
		ups:    make(map[int64]upRoutine),
		api:    api,
	}
}
