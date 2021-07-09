package info

var _ Infoer = (*Dynamic)(nil)

type Dynamic struct {
	Info
	Content string
	Card    interface{}
	URL     string
}

func (Dynamic) Type() Type {
	return TDynamic
}

func (dynamic Dynamic) GetData() []interface{} {
	return []interface{}{dynamic.Content, dynamic.Card}
}
