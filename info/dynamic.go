package info

var _ Infoer = (*Dynamic)(nil)

type Dynamic struct {
	Name        string
	Time        int32
	Content     string
	Card        string
	RID         int64
	CommentType uint8
}

func (Dynamic) Type() Type {
	return TDynamic
}

func (dynamic *Dynamic) GetInstance() interface{} {
	return dynamic
}
