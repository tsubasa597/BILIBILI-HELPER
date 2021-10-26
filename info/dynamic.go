package info

// Dynamic 监听的动态信息
type Dynamic struct {
	Info
	UID     int64
	Content string
	Card    string
	RID     int64
	Offect  int64
	Type    Type
}

var _ Interface = (*Dynamic)(nil)

// GetInstance 将 Infoer 接口转为 interface{}
func (d *Dynamic) GetInstance() interface{} {
	return d
}

// Type 动态类型
type Type uint8

const (
	// CommentViedo 视频
	CommentViedo Type = iota + 1
	// CommentTopic 话题
	CommentTopic
	_
	// CommentActivity 活动
	CommentActivity
	_
	_
	// CommentNotice 公告
	CommentNotice
	// CommentLiveActivity 直播活动
	CommentLiveActivity
	// CommentActivityViedo 活动稿件
	CommentActivityViedo
	// CommentLiveNotice 直播公告
	CommentLiveNotice
	// CommentDynamicImage 相簿（图片动态）
	CommentDynamicImage
	// CommentColumn 专栏
	CommentColumn
	_
	// CommentAudio 音频
	CommentAudio
	_
	_
	// CommentDynamic 动态（纯文字动态&分享）
	CommentDynamic
)
