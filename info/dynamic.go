package info

var _ Infoer = (*Dynamic)(nil)

type Dynamic struct {
	Info
	Content string
	Card    interface{}
}
