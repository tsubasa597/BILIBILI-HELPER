package info

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

// Sort 排序
type Sort uint8

// Comment 爬取的评论信息
type Comment struct {
	Info
	DynamicUID int64
	RID        int64
	UID        int64
	Rpid       int64
	LikeNum    uint32
	Content    string
}
