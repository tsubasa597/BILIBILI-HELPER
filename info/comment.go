package info

// Comment 爬取的评论信息
type Comment struct {
	Info
	UID     int64
	Rpid    int64
	Like    uint32
	Content string
}

var _ Interface = (*Comment)(nil)

func (*Comment) GetTime() int32 {
	return 1
}

// GetInstance 将 Infoer 接口转为 interface{}
func (c *Comment) GetInstance() interface{} {
	return c
}

// GetType 判断监听类型，转换 interface{}
func (Comment) GetType() Type {
	return TComment
}
