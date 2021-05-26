package info

var _ Infoer = (*Dynamic)(nil)

type Dynamic struct {
	Info
	Content string
	Card    interface{}
}

func (Dynamic) Type() Type {
	return TDynamic
}
