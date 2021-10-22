package api

import (
	fmt "fmt"

	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"github.com/tsubasa597/requests"
)

const (
	// SortDesc 评论区按时间倒序排序
	SortDesc uint8 = iota
	// SortAsc 评论区按时间正序排序
	SortAsc

	// MaxPs 一页评论的最大数量
	MaxPs int = 49
	// MinPs 一页评论的最小数量
	MinPs int = 20
)

// GetComments 根据评论区类型、排序类型(正序、逆序)、评论区 id、页码 获取评论
func GetComments(commentType, sort uint8, oid int64, ps, pn int) (*Comments_Data, error) {
	resp := &Comments{}
	if err := requests.Gets(fmt.Sprintf("%s?type=%d&oid=%d&sort=%d&ps=%d&pn=%d",
		reply, commentType, oid, sort, ps, pn), resp); err != nil {
		return nil, err
	}

	if len(resp.Data.Replies) == 0 {
		return nil, fmt.Errorf(ecode.ErrNoComment)
	}

	if resp.Code != ecode.Sucess {
		return nil, fmt.Errorf(resp.Message)
	}

	return resp.Data, nil
}

// GetAllComments 获取指定时间为止的所有评论
func GetAllComments(commentType uint8, oid, t int64) (comments []*info.Comment) {
	for pn := 1; ; pn++ {
		data, err := GetComments(commentType, SortDesc, oid, MaxPs, pn)
		if err != nil {
			return
		}

		for _, reply := range data.Replies {
			if reply.Ctime <= t {
				return
			}

			comments = append(comments, &info.Comment{
				Info: info.Info{
					Name: reply.Member.Uname,
					Time: reply.Ctime,
				},
				UserID:    data.Upper.Mid,
				UID:       reply.Mid,
				Rpid:      reply.Rpid,
				Like:      uint32(reply.Like),
				Content:   reply.Content.Message,
				DynamicID: oid,
			})
		}
	}
}
