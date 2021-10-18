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
	RID   int64
	Type  uint8
	Time  int32
	Pn    int
	ps    int
	state state.State
	log   *logrus.Entry
}

var _ Tasker = (*Comment)(nil)

// Run 开始运行
func (c *Comment) Run(ch chan<- interface{}) {
	if c.state != state.Runing {
		return
	}

	resp, err := api.GetComments(c.Type, 0, c.RID, c.ps, c.Pn)
	if err != nil {
		if err.Error() == ecode.ErrNoComment {
			c.state = state.Stop
			return
		}

		c.log.Error(err)
		return
	}

	infos := make([]*info.Comment, 0, len(resp.Data.Replies))
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
			DynamicID: c.RID,
		})
	}

	c.Pn++
	ch <- infos
}

// Next 下次运行时间
func (c Comment) Next(t time.Time) time.Time {
	return t.Add(time.Second * 2)
}

// NewComment 初始化
func NewComment(rid int64, typ uint8, ps int, log *logrus.Entry) *Comment {
	if log == nil {
		log = logrus.NewEntry(logrus.New())
	}

	return &Comment{
		RID:   rid,
		Type:  typ,
		Pn:    1,
		state: state.Runing,
		ps:    ps,
		log:   log,
	}
}
