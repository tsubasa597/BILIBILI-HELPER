package info

type Infoer interface {
	Type() Type
	GetInstance() interface{}
}

type Type int16

const (
	TLive Type = iota + 1
	TDynamic
)
