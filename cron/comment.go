package cron

import (
	"fmt"
	"sync"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api/comment"
	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Comment 评论区参数
type Comment struct {
	RID      int64
	Type     comment.Type
	time     int64
	pn       int
	timeCell time.Duration
	log      *zap.Logger
	mutex    *sync.Mutex
}

const (
	TwoDay time.Duration = time.Hour * 24 * 2
)

const (
	_errCKey = "评论区参数"
)

var (
	_ Tasker = (*Comment)(nil)
)

// NewComment 初始化
// 时间间隔 timeCell 的单位为 **秒**
func NewComment(rid, t int64, timeCell time.Duration, typ comment.Type, log *zap.Logger) *Comment {
	if log == nil {
		log = zap.NewExample()
	}

	return &Comment{
		RID:      rid,
		Type:     typ,
		time:     t,
		pn:       1, // 开始爬取页数的初始值
		timeCell: timeCell * time.Second,
		log:      log,
		mutex:    &sync.Mutex{},
	}
}

// Run 开始运行
func (c *Comment) Run(ch chan<- interface{}, wg *sync.WaitGroup) {
	// 防止请求过快
	defer func() {
		time.Sleep(time.Second * 1)
		wg.Done()
	}()

	c.mutex.Lock()

	infos, err := comment.Get(c.Type, comment.Desc, c.RID, comment.MaxPs, c.pn)
	if err != nil {
		// 爬取完成,在间隔之间之后继续更新
		if err.Error() == ecode.ErrNoComment {
			c.pn = 1
			return
		}

		c.log.Error(err.Error(), zapcore.Field{
			Key:    _errCKey,
			Type:   zapcore.StringType,
			String: fmt.Sprintf("RID: %d, Type: %d", c.RID, c.Type),
		})
		return
	}

	c.pn++

	c.mutex.Unlock()

	ch <- infos
}

// Next 下次运行时间
func (c Comment) Next(t time.Time) time.Time {
	if t.AddDate(0, 0, -7).Unix() > c.time {
		return t.Add(TwoDay)
	}

	return t.Add(c.timeCell)
}

// Info 返回任务的信息
func (c Comment) Info() Info {
	return Info{
		ID:       c.RID,
		Type:     c.Type,
		TimeCell: c.timeCell,
		Time:     c.time,
	}
}
