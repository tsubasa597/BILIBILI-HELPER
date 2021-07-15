package info

type Infoer interface {
	Type() Type
	GetT() int32
	GetName() string
}

type Type int16

const (
	TLive Type = iota + 1
	TDynamic
)

type Info struct {
	T    int32
	Name string
}

func (info Info) GetT() int32 {
	return info.T
}

func (info Info) GetName() string {
	return info.Name
}
