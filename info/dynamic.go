package info

var _ Infoer = (*Dynamic)(nil)

// Dynamic 监听的动态信息
type Dynamic struct {
	Name        string
	Time        int32
	Content     string
	Card        string
	RID         int64
	CommentType uint8
	DynamicID   int64
}

// Type 判断监听类型，转换 interface{}
func (Dynamic) Type() Type {
	return TDynamic
}

// GetInstance 将 Infoer 接口转为 interface{}
func (dynamic *Dynamic) GetInstance() interface{} {
	return dynamic
}
