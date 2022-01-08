package comment

import (
	fmt "fmt"
	"sync"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/api/proto"
	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/requests"
)

// Info 爬取的评论信息
type Info struct {
	Name       string
	Time       int64
	DynamicUID int64
	RID        int64
	UID        int64
	Rpid       int64
	LikeNum    uint32
	Content    string
}

// Sort 排序
type Sort uint8

const (
	// Desc 评论区按时间倒序排序
	Desc Sort = iota
	// Asc 评论区按时间正序排序
	Asc

	// MaxPs 一页评论的最大数量
	MaxPs = 49
	// MinPs 一页评论的最小数量
	MinPs = 20
)

// Type 动态类型
type Type uint8

const (
	// Viedo 视频
	Viedo Type = iota + 1
	// Topic 话题
	Topic
	_
	// Activity 活动
	Activity
	_
	_
	// Notice 公告
	Notice
	// LiveActivity 直播活动
	LiveActivity
	// ActivityViedo 活动稿件
	ActivityViedo
	// LiveNotice 直播公告
	LiveNotice
	// DynamicImage 相簿（图片动态）
	DynamicImage
	// Column 专栏
	Column
	_
	// Audio 音频
	Audio
	_
	_
	// Dynamic 动态（纯文字动态&分享）
	Dynamic
)

var (
	_commentPool *sync.Pool = &sync.Pool{
		New: func() interface{} {
			return &Info{}
		},
	}
)

// Get 根据评论区类型、排序类型(正序、逆序)、评论区 id、页码 获取评论
func Get(commentType Type, sort Sort, rid int64, ps, pn int) ([]Info, error) {
	resp := &proto.Comments{}
	err := requests.Gets(fmt.Sprintf("%s?type=%d&oid=%d&sort=%d&ps=%d&pn=%d",
		api.Reply, commentType, rid, sort, ps, pn), resp)
	if err != nil {
		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	if resp.Code != ecode.Sucess {
		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: resp.Message,
		}
	}

	if len(resp.Data.Replies) == 0 {
		return nil, ecode.APIErr{
			E: ecode.ErrNoComment,
		}
	}

	comments := make([]Info, 0, len(resp.Data.Replies))
	for _, reply := range resp.Data.Replies {
		comment := _commentPool.Get().(*Info)
		comment.Name = reply.Member.Uname
		comment.Time = reply.Ctime
		comment.DynamicUID = resp.Data.Upper.Mid
		comment.UID = reply.Mid
		comment.Rpid = reply.Rpid
		comment.LikeNum = uint32(reply.Like)
		comment.Content = reply.Content.Message
		comment.RID = rid

		comments = append(comments, *comment)
		_commentPool.Put(comment)
	}

	return comments, nil
}

// GetAll 获取指定时间为止的所有评论
func GetAll(commentType Type, rid, t int64) (comments []Info) {
	for pn := 1; ; pn++ {
		infos, err := Get(commentType, Desc, rid, MaxPs, pn)
		if err != nil {
			return
		}

		for _, inf := range infos {
			if inf.Time <= t {
				return
			}

			comments = append(comments, inf)
		}

		time.Sleep(time.Second)
	}
}
