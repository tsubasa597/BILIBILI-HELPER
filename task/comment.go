package task

import (
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
	Time     int32
	Pn       int
	ps       int
	timeCell time.Duration
	state    state.State
	log      *logrus.Entry
}

var _ Tasker = (*Comment)(nil)

// Run 开始运行
func (c *Comment) Run(ch chan<- interface{}) {
	if c.state != state.Runing {
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

	infos := make([]*info.Comment, 0, len(data.Replies))
	for _, inf := range data.Replies {
		infos = append(infos, &info.Comment{
			Info: info.Info{
				Name: inf.Member.Uname,
				Time: inf.Ctime,
			},
			UserID:  data.Upper.Mid,
			UID:     inf.Mid,
			Rpid:    inf.Rpid,
			Like:    uint32(inf.Like),
			Content: inf.Content.Message,
			RID:     c.RID,
		})
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
