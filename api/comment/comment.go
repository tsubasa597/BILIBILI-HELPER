package comment

import (
	fmt "fmt"
	"sync"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api/proto"
	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"github.com/tsubasa597/requests"
)

var (
	_commentPool *sync.Pool = &sync.Pool{
		New: func() interface{} {
			return &info.Comment{}
		},
	}
)

// GetComments 根据评论区类型、排序类型(正序、逆序)、评论区 id、页码 获取评论
func GetComments(commentType info.DynamicType, sort info.Sort, rid int64, ps, pn int) ([]info.Comment, error) {
	resp := &proto.Comments{}
	err := requests.Gets(fmt.Sprintf("%s?type=%d&oid=%d&sort=%d&ps=%d&pn=%d",
		info.Reply, commentType, rid, sort, ps, pn), resp)
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

	comments := make([]info.Comment, 0, len(resp.Data.Replies))
	for _, reply := range resp.Data.Replies {
		comment := _commentPool.Get().(*info.Comment)
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

// GetAllComments 获取指定时间为止的所有评论
func GetAllComments(commentType info.DynamicType, rid, t int64) (comments []info.Comment) {
	for pn := 1; ; pn++ {
		infos, err := GetComments(commentType, info.SortDesc, rid, info.MaxPs, pn)
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
