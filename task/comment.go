package task

import (
	"sync"
	"sync/atomic"
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
		pn:       1, // 开始爬取页数的初始值
		state:    state.Stop,
		timeCell: timeCell, // 爬取时间间隔
		log:      log,
	}
}

// Run 开始运行
func (c *Comment) Run(ch chan<- interface{}, wg *sync.WaitGroup) {
	// 防止请求过快
	defer func() {
		time.Sleep(time.Second * 1)
		wg.Done()
	}()

	if atomic.LoadInt32((*int32)(&c.state)) != int32(state.Stop) {
		return
	}

	atomic.SwapInt32((*int32)(&c.state), int32(state.Runing))
	defer func() {
		atomic.SwapInt32((*int32)(&c.state), int32(state.Stop))
	}()

	infos, err := comment.GetComments(c.Type, info.SortDesc, c.RID, info.MaxPs, c.pn)
	if err != nil {
		// 爬取完成,在间隔之间之后继续更新
		if err.Error() == ecode.ErrNoComment {
			c.pn = 1
			return
		}

		c.log.Error(err)
		return
	}

	c.pn++

	ch <- infos
}

// Next 下次运行时间
func (c Comment) Next(t time.Time) time.Time {
	if t.AddDate(0, 0, -7).Unix() > c.Time {
		return t.Add(info.TwoDay)
	}

	return t.Add(time.Second * c.timeCell)
}
