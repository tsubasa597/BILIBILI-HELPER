package api_test

import (
	"testing"

	"github.com/tsubasa597/BILIBILI-HELPER/api/comment"
)

type commentInfo struct {
	_type comment.Type
	sort  comment.Sort
	rid   int64
	ps    int
	pn    int
}

var (
	_comments = []commentInfo{
		{
			_type: comment.Dynamic,
			sort:  comment.Desc,
			rid:   593260684386323648,
			ps:    10,
			pn:    1,
		},
		{
			_type: comment.Dynamic,
			sort:  comment.Desc,
			rid:   593265499053075372,
			ps:    10,
			pn:    1,
		},
		{
			_type: comment.Dynamic,
			sort:  comment.Desc,
			rid:   593621130926097907,
			ps:    10,
			pn:    1,
		},
	}
)

func TestGetComment(t *testing.T) {
	for _, comm := range _comments {
		_, err := comment.Get(comm._type, comm.sort, comm.rid, comm.ps, comm.pn)
		if err != nil {
			t.Error(err)
		}
	}
}
