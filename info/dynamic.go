package info

var _ Infoer = (*Dynamic)(nil)

type Dynamic struct {
	Info
	Content string
	Card    interface{}
}

func (dynamic Dynamic) GetData() []interface{} {
	return []interface{}{dynamic.Card, dynamic.Content}
}
