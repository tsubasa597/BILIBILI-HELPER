package info

var _ Infoer = (*Dynamic)(nil)

type Dynamic struct {
	Info
	Content     string
	Card        string
	RID         int64
	CommentType uint8
}

func (Dynamic) Type() Type {
	return TDynamic
}
