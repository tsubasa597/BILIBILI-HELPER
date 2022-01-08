package cron

import (
	"fmt"
	"sync"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api/comment"
	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Comment 评论区参数
type Comment struct {
	RID      int64
	Type     info.DynamicType
	Time     int64
	pn       int
	timeCell time.Duration
	log      *zap.Logger
	rwMutex  *sync.RWMutex
}

var (
	_ Tasker = (*Comment)(nil)
)

// NewComment 初始化
// 时间间隔 timeCell 的单位为 **秒**
func NewComment(rid, t int64, timeCell time.Duration, typ info.DynamicType, log *zap.Logger) *Comment {
	if log == nil {
		log = zap.NewExample()
	}

	return &Comment{
		RID:      rid,
		Type:     typ,
		Time:     t,
		pn:       1, // 开始爬取页数的初始值
		timeCell: timeCell * time.Second,
		log:      log,
		rwMutex:  &sync.RWMutex{},
	}
}

// Run 开始运行
func (c *Comment) Run(ch chan<- interface{}, wg *sync.WaitGroup) {
	// 防止请求过快
	defer func() {
		time.Sleep(time.Second * 1)
		wg.Done()
	}()

	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

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
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	if t.AddDate(0, 0, -7).Unix() > c.Time {
		return t.Add(info.TwoDay)
	}

	return t.Add(c.timeCell)
}

// Info 返回任务的信息
func (c Comment) Info() info.Task {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	return info.Task{
		ID:       c.RID,
		Type:     c.Type,
		TimeCell: c.timeCell,
		Time:     c.Time,
	}
}
