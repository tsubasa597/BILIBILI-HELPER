package dynamic

import (
	"context"
	"fmt"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
)

type DynamicListen struct {
	C      chan []interface{}
	ctx    context.Context
	cancel context.CancelFunc
	ups    map[int64]context.CancelFunc
}

func (d *DynamicListen) ListenDynamic(uid int64) {
	if _, ok := d.ups[uid]; ok {
		return
	}

	ct, cl := context.WithCancel(d.ctx)
	d.ups[uid] = cl
	go listen(uid, ct, d.C)
}

func (d *DynamicListen) StopListenUP(uid int64) error {
	if _, ok := d.ups[uid]; !ok {
		return fmt.Errorf("该用户未监听")
	}

	d.ups[uid]()
	return nil
}

// Stop 释放资源
func (d DynamicListen) Stop() {
	d.cancel()
}

func listen(uid int64, ctx context.Context, ch chan<- []interface{}) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.NewTicker(time.Minute).C:
			t, c, e := api.GetDynamicMessage(uid)
			ch <- []interface{}{uid, t, c, e}
		}
	}
}

func New() *DynamicListen {
	c, cl := context.WithCancel(context.Background())
	return &DynamicListen{
		C:      make(chan []interface{}, 1),
		ctx:    c,
		cancel: cl,
		ups:    make(map[int64]context.CancelFunc),
	}
}
