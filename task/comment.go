package task

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"github.com/tsubasa597/BILIBILI-HELPER/state"
)

// Comment 评论区参数
type Comment struct {
	RID      int64
	Type     info.Type
	Pn       int
	ps       int
	timeCell time.Duration
	state    state.State
	log      *logrus.Entry
}

var (
	_           Tasker     = (*Comment)(nil)
	commentPool *sync.Pool = &sync.Pool{
		New: func() interface{} {
			return &info.Comment{}
		},
	}
)

// Run 开始运行
func (c *Comment) Run(ch chan<- interface{}) {
	if c.state == state.Stop {
		return
	}

	data, err := api.GetComments(c.Type, 0, c.RID, c.ps, c.Pn)
	if err != nil {
		if err.Error() == ecode.ErrNoComment {
			c.state = state.Stop
			return
		}

		c.state = state.Pause
		c.log.Error(err)
		return
	}

	infos := make([]info.Comment, 0, len(data.Replies))
	for _, inf := range data.Replies {
		comm := commentPool.Get().(*info.Comment)
		comm.Name = inf.Member.Uname
		comm.Time = inf.Ctime
		comm.DynamicUID = data.Upper.Mid
		comm.UID = inf.Mid
		comm.Rpid = inf.Rpid
		comm.LikeNum = uint32(inf.Like)
		comm.Content = inf.Content.Message
		comm.RID = c.RID

		infos = append(infos, *comm)
		commentPool.Put(comm)
	}

	c.Pn++
	if c.state == state.Pause {
		c.state = state.Runing
	}

	ch <- infos
}

// State 获取运行状态
func (c Comment) State() state.State {
	return c.state
}

// Next 下次运行时间
func (c Comment) Next(t time.Time) time.Time {
	if c.state == state.Pause {
		return t.Add(PauseDuration)
	}

	return t.Add(time.Second * c.timeCell)
}

// NewComment 初始化
func NewComment(rid int64, timeCell time.Duration, typ info.Type, ps int, log *logrus.Entry) *Comment {
	if log == nil {
		log = logrus.NewEntry(logrus.New())
	}

	return &Comment{
		RID:      rid,
		Type:     typ,
		Pn:       1,
		state:    state.Runing,
		ps:       ps,
		timeCell: timeCell,
		log:      log,
	}
}
