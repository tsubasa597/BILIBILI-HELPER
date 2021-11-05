package api

import (
	fmt "fmt"
	"sync"

	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"github.com/tsubasa597/requests"
)

var (
	commentPool *sync.Pool = &sync.Pool{
		New: func() interface{} {
			return &info.Comment{}
		},
	}
)

// GetComments 根据评论区类型、排序类型(正序、逆序)、评论区 id、页码 获取评论
func GetComments(commentType info.Type, sort info.Sort, rid int64, ps, pn int) (*Comments_Data, error) {
	resp := &Comments{}
	if err := requests.Gets(fmt.Sprintf("%s?type=%d&oid=%d&sort=%d&ps=%d&pn=%d",
		reply, commentType, rid, sort, ps, pn), resp); err != nil {
		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	if len(resp.Data.Replies) == 0 {
		return nil, ecode.APIErr{
			E: ecode.ErrNoComment,
		}
	}

	if resp.Code != ecode.Sucess {
		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: resp.Message,
		}
	}

	return resp.Data, nil
}

// GetAllComments 获取指定时间为止的所有评论
func GetAllComments(commentType info.Type, rid, t int64) (comments []info.Comment) {
	for pn := 1; ; pn++ {
		data, err := GetComments(commentType, info.SortDesc, rid, info.MaxPs, pn)
		if err != nil {
			return
		}

		for _, reply := range data.Replies {
			if reply.Ctime <= t {
				return
			}

			comment := commentPool.Get().(*info.Comment)
			comment.Name = reply.Member.Uname
			comment.UserID = data.Upper.Mid
			comment.UID = reply.Mid
			comment.Rpid = reply.Rpid
			comment.LikeNum = uint32(reply.Like)
			comment.Content = reply.Content.Message
			comment.RID = rid

			comments = append(comments, *comment)
			commentPool.Put(comment)
		}
	}
}
