package info

// Comment 爬取的评论信息
type Comment struct {
	Info
	UserID  int64
	RID     int64
	UID     int64
	Rpid    int64
	Like    uint32
	Content string
}

var _ Interface = (*Comment)(nil)

// GetInstance 将 Infoer 接口转为 interface{}
func (c *Comment) GetInstance() interface{} {
	return c
}

// Sort 排序
type Sort uint8

const (
	// SortDesc 评论区按时间倒序排序
	SortDesc Sort = iota
	// SortAsc 评论区按时间正序排序
	SortAsc

	// MaxPs 一页评论的最大数量
	MaxPs int = 49
	// MinPs 一页评论的最小数量
	MinPs int = 20
)
