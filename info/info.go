package info

type Infoer interface{}

var _ Infoer = (*Info)(nil)

type Info struct {
	T    int32
	Err  error
	Name string
}

func (info Info) GetName() string {
	return info.Name
}
