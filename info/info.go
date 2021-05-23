package info

type Infoer interface {
	GetErr() error
	GetT() int32
	GetName() string
	GetData() []interface{}
}

type Info struct {
	T    int32
	Err  error
	Name string
}

func (info Info) GetName() string {
	return info.Name
}

func (info Info) GetErr() error {
	return info.Err
}

func (info Info) GetT() int32 {
	return info.T
}
