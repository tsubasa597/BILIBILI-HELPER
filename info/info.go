package info

type Infoer interface {
	Type() Type
	GetData() []interface{}
	GetT() int32
	GetErr() error
	GetName() string
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

func (info Info) GetT() int32 {
	return info.T
}

func (info Info) GetErr() error {
	return info.Err
}

func (info Info) GetName() string {
	return info.Name
}
