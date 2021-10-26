// Package info 监听返回的消息内容
package info

// Interface 监听信息的接口
type Interface interface {
	GetInstance() interface{}
	GetTime() int64
	GetName() string
}

// Info 监听返回值的基础结构体
type Info struct {
	Name string
	Time int64
}

var _ Interface = (*Info)(nil)

// GetInstance 将 Infoer 接口转为 interface{}
func (i *Info) GetInstance() interface{} {
	return i
}

// GetTime 返回时间
func (i Info) GetTime() int64 {
	return i.Time
}

// GetName 返回名称
func (i Info) GetName() string {
	return i.Name
}
