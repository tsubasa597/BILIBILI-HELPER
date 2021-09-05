package state

import "context"

// Info 获取的状态信息
type Info struct {
	Name string
	Time int32
	Ctx  context.Context
}
