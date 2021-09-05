// Package info 监听返回的消息内容
package info

// Interface 监听信息的接口
type Interface interface {
	GetType() Type
	GetInstance() interface{}
	GetTime() int32
	GetName() string
}

type Info struct {
	Name string
	Time int32
}

var _ Interface = (*Info)(nil)

func (Info) GetType() Type {
	return TBase
}
func (i *Info) GetInstance() interface{} {
	return i
}

func (i Info) GetTime() int32 {
	return i.Time
}
func (i Info) GetName() string {
	return i.Name
}

// Type Info 的类型
type Type uint8

const (
	TBase Type = iota
	// TLive 直播
	TLive
	// TDynamic 动态
	TDynamic
	// TComment 评论
	TComment

	TDynamicState
)
