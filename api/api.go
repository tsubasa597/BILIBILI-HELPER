package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/BILIBILI-HELPER/global"
)

const (
	BaseHost     = "https://api.bilibili.com"
	BaseLiveHost = "https://api.live.bilibili.com"
	BaseVCHost   = "https://api.vc.bilibili.com"
	VideoView    = "https://www.bilibili.com/video"
	DynamicView  = "https://t.bilibili.com"

	RoomInit               = BaseLiveHost + "/room/v1/Room/room_init"
	SpaceAccInfo           = BaseHost + "/x/space/acc/info"
	DynamicSrvSpaceHistory = BaseVCHost + "/dynamic_svr/v1/dynamic_svr/space_history"
	GetRoomInfoOld         = BaseLiveHost + "/room/v1/Room/getRoomInfoOld"
	DynamicSrvDynamicNew   = BaseVCHost + "/dynamic_svr/v1/dynamic_svr/dynamic_new"
	RelationModify         = BaseHost + "/x/relation/modify"
	RelationFeedList       = BaseLiveHost + "/relation/v1/feed/feed_list"
	GetAttentionList       = BaseVCHost + "/feed/v1/feed/get_attention_list"
	UserLogin              = BaseHost + "/x/web-interface/nav"
	VideoHeartbeat         = BaseHost + "/x/click-interface/web/heartbeat"
	AvShare                = BaseHost + "/x/web-interface/share/add"
	Sliver2CoinsStatus     = BaseLiveHost + "/pay/v1/Exchange/getStatus"
	Sliver2Coins           = BaseLiveHost + "/pay/v1/Exchange/silver2coin"
	LiveCheckin            = BaseLiveHost + "/xlive/web-ucenter/v1/sign/DoSign"
	LiverStatus            = BaseHost + "/x/space/acc/info"

	errGetDynamic    = "请求发生错误"
	errNoDynamic     = "该用户没有动态"
	errUnknowDynamic = "未知动态"
	errNotListen     = "该用户未监听"
	errRepeatListen  = "重复监听"
)

type API struct {
	conf global.Cookie
	r    *global.Requests
	log  *logrus.Logger
}

type Info struct {
	T       int32
	Err     error
	Content string
	Card    interface{}
	Name    string
	Live
}

type Live struct {
	LiveStatus  bool
	LiveRoomURL string
	LiveTitle   string
}

func New(c global.Cookie, log *logrus.Logger) *API {
	r := global.NewRequests()
	r.SetHeader(http.Header{
		"Connection":   []string{"keep-alive"},
		"User-Agent":   []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.70"},
		"Cookie":       []string{c.GetVerify()},
		"Content-Type": []string{"application/x-www-form-urlencoded"},
	})
	return &API{
		r:    r,
		conf: c,
		log:  log,
	}
}

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
		l.api.log.Debugln(l.ups[uid], fmt.Errorf(errRepeatListen))
		return l.ups[uid].ctx, nil, fmt.Errorf(errRepeatListen)
	}

	ct, cl := context.WithCancel(l.ctx)
	l.ups[uid] = upRoutine{
		cancel: cl,
		ctx:    ct,
	}
	ch := make(chan Info, 1)

	go func() {
		l.api.log.Debugln(fmt.Sprintf("Start : %T %d", f, uid))
		for {
			select {
			case <-l.ctx.Done():
				l.api.log.Debugln(fmt.Sprintf("Stop : %T %d", f, uid))
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
	d.api.log.Infoln("All Goroutine Done")
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
