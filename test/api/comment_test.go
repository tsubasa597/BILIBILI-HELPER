package api_test

import (
	"testing"

	"github.com/tsubasa597/BILIBILI-HELPER/api/comment"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
)

type commentInfo struct {
	_type info.DynamicType
	sort  info.Sort
	rid   int64
	ps    int
	pn    int
}

var (
	comments = []commentInfo{
		{
			_type: info.CommentDynamic,
			sort:  info.SortDesc,
			rid:   593260684386323648,
			ps:    10,
			pn:    1,
		},
		{
			_type: info.CommentDynamic,
			sort:  info.SortDesc,
			rid:   593265499053075372,
			ps:    10,
			pn:    1,
		},
		{
			_type: info.CommentDynamic,
			sort:  info.SortDesc,
			rid:   593621130926097907,
			ps:    10,
			pn:    1,
		},
	}
)

func TestGetComment(t *testing.T) {
	for _, comm := range comments {
		data, err := comment.GetComments(comm._type, comm.sort, comm.rid, comm.ps, comm.pn)
		if err != nil {
			t.Error(err)
		}

		t.Log(data[0].Content)
	}
}
