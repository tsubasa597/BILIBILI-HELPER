// Package info 监听返回的消息内容
package info

// Infoer 监听信息的接口
type Infoer interface {
	Type() Type
	GetInstance() interface{}
}

// Type Info 的类型
type Type uint8

const (
	// TLive 直播
	TLive Type = iota + 1
	// TDynamic 动态
	TDynamic
)
