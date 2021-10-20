package info

// Dynamic 监听的动态信息
type Dynamic struct {
	Info
	UID         int64
	Content     string
	Card        string
	RID         int64
	CommentType uint8
}

var _ Interface = (*Dynamic)(nil)

// GetInstance 将 Infoer 接口转为 interface{}
func (d *Dynamic) GetInstance() interface{} {
	return d
}

// GetType 判断监听类型，转换 interface{}
func (Dynamic) GetType() Type {
	return TDynamic
}
