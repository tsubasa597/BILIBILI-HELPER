package info

type Infoer interface {
	Type() Type
}

type Type int16

const (
	TLive Type = iota + 1
	TDynamic
)

type Info struct {
	T    int32
	Err  error
	Name string
}
