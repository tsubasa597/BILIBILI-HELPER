package task

import (
	"fmt"
	"sync"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api/comment"
	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"github.com/tsubasa597/BILIBILI-HELPER/state"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Comment 评论区参数
type Comment struct {
	RID      int64
	Type     info.Type
	Time     int64
	pn       int
	timeCell time.Duration
	state    *atomic.Int32
	log      *zap.Logger
}

var (
	_ Tasker = (*Comment)(nil)
)

// NewComment 初始化
// 时间间隔 timeCell 的单位为 **秒**
func NewComment(rid, t int64, timeCell time.Duration, typ info.Type, log *zap.Logger) *Comment {
	if log == nil {
		log = zap.NewExample()
	}

	return &Comment{
		RID:      rid,
		Type:     typ,
		Time:     t,
		pn:       1, // 开始爬取页数的初始值
		state:    atomic.NewInt32(state.Stop),
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

	if c.state.Load() != state.Stop {
		return
	}

	c.state.Swap(state.Running)
	defer func() {
		c.state.Swap(state.Stop)
	}()

	infos, err := comment.GetComments(c.Type, info.SortDesc, c.RID, info.MaxPs, c.pn)
	if err != nil {
		// 爬取完成,在间隔之间之后继续更新
		if err.Error() == ecode.ErrNoComment {
			c.pn = 1
			return
		}

		c.log.Error(err.Error(), zapcore.Field{
			Key:    "评论区参数",
			Type:   zapcore.StringType,
			String: fmt.Sprintf("RID: %d, Type: %d", c.RID, c.Type),
		})
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
