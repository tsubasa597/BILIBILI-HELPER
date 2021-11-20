package task

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/BILIBILI-HELPER/api/comment"
	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"github.com/tsubasa597/BILIBILI-HELPER/state"
)

// Comment 评论区参数
type Comment struct {
	RID      int64
	Type     info.Type
	Time     int64
	pn       int
	timeCell time.Duration
	state    state.State
	log      *logrus.Logger
}

var (
	_ Tasker = (*Comment)(nil)
)

// NewComment 初始化
func NewComment(rid, t int64, timeCell time.Duration, typ info.Type, log *logrus.Logger) *Comment {
	if log == nil {
		log = logrus.New()
	}

	return &Comment{
		RID:      rid,
		Type:     typ,
		Time:     t,
		pn:       1, /** 开始爬取页数的初始值 */
		state:    state.Runing,
		timeCell: timeCell, /** 爬取时间间隔 */
		log:      log,
	}
}

// Run 开始运行
func (c *Comment) Run(ch chan<- interface{}) {
	// 防止请求过快
	defer time.Sleep(time.Second * 1)

	if c.state == state.Stop {
		return
	}

	infos, err := comment.GetComments(c.Type, info.SortDesc, c.RID, info.MaxPs, c.pn)
	if err != nil {
		c.state = state.Pause

		// 爬取完成,在间隔之间之后继续更新
		if err.Error() == ecode.ErrNoComment {
			c.pn = 1
			return
		}

		c.log.Error(err)
		return
	}

	c.pn++

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
		if time.Now().AddDate(0, 0, -7).Unix() < c.Time {
			return t.Add(info.TwoDay)
		}

		return t.Add(info.Pause)
	}

	return t.Add(time.Second * c.timeCell)
}
