package listen

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"github.com/tsubasa597/BILIBILI-HELPER/state"
	"golang.org/x/sync/semaphore"
)

// Comment 评论
type Comment struct {
	comms state.ListenStater
	limit *semaphore.Weighted
}

var _ Listener = (*Comment)(nil)

// ListenInfo 获取指定评论区的评论
func (c *Comment) ListenInfo(rid int64) (infos []info.Interface, err error) {
	var limit bool

	infos, err = c.comms.Do(rid, func(rid int64, l state.Listener, log *logrus.Entry) (infos []info.Interface) {
		comm := l.(*CommLinsten)
		c.limit.Acquire(comm.Ctx, 1)
		defer c.limit.Release(1)

		resp, err := api.GetComments(comm.Type, 0, comm.RID, comm.ps, comm.Pn)
		if err != nil {
			return
		}

		for _, inf := range resp.Data.Replies {
			infos = append(infos, &info.Comment{
				Info: info.Info{
					Name: inf.Member.Uname,
					Time: inf.Ctime,
				},
				UserID:    resp.Data.Upper.Mid,
				UID:       inf.Mid,
				Rpid:      inf.Rpid,
				Like:      uint32(inf.Like),
				Content:   inf.Content.Message,
				DynamicID: rid,
			})
		}

		comm.Pn++
		return
	})

	if limit {
		c.comms.Pause(30)
	}

	return
}

// GetState 获取状态
func (c *Comment) GetState() state.State {
	return c.comms.GetState()
}

// StopListenUP 停止获取评论区评论
func (c Comment) StopListenUP(uid int64) error {
	return c.comms.StopOne(uid)
}

// GetList 获取状态
func (c Comment) GetList() []state.Info {
	return c.comms.GetAll()
}

// Add 添加监听 评论区一页数量为 49
func (c Comment) Add(ctx context.Context, cancel context.CancelFunc, rid, typ int64) error {
	return c.comms.Put(rid, NewCommLinsten(ctx, cancel, rid, uint8(typ), 49))
}

// NewComment 初始化
func NewComment(ctx context.Context, weight int64, log *logrus.Entry) *Comment {
	return &Comment{
		limit: semaphore.NewWeighted(weight),
		comms: state.NewDeListenState(ctx, log),
	}
}

// CommLinsten 监听信息  !!非线程安全!!
type CommLinsten struct {
	Ctx    context.Context
	RID    int64
	Type   uint8
	Time   int64
	Pn     int
	ps     int
	state  state.State
	cancel context.CancelFunc
}

var _ state.Listener = (*CommLinsten)(nil)

// NewCommLinsten 初始化
func NewCommLinsten(ctx context.Context, cancel context.CancelFunc, rid int64, typ uint8, ps int) *CommLinsten {
	return &CommLinsten{
		Ctx:    ctx,
		RID:    rid,
		Type:   typ,
		Pn:     1,
		state:  state.Runing,
		ps:     ps,
		cancel: cancel,
	}
}

// GetInfo 获取监听信息
func (c CommLinsten) GetInfo() state.Info {
	return state.Info{
		Time: c.Time,
		Ctx:  c.Ctx,
	}
}

// GetState 获取当前状态
func (c CommLinsten) GetState() state.State {
	return c.state
}

// Update 方法体是空
func (c CommLinsten) Update(num interface{}) {
	return
}

// Pause 从运行状态暂停
func (c CommLinsten) Pause() bool {
	switch c.state {
	case state.Pause:
		return true
	case state.Runing:
		c.state = state.Pause
		return true
	default:
		return false
	}
}

// Start 从暂停状态开始
func (c CommLinsten) Start() bool {
	switch c.state {
	case state.Pause:
		c.state = state.Runing
		return true
	case state.Runing:
		return true
	default:
		return false
	}
}

// Stop 停止
func (c CommLinsten) Stop() {
	c.cancel()
}
