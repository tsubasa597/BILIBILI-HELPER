package api

type RequestErr struct{}

type DynamicUnknown struct{}

var _ error = (*DynamicUnknown)(nil)
var _ error = (*RequestErr)(nil)

func (r RequestErr) Error() string {
	return "请求错误"
}

func (d DynamicUnknown) Error() string {
	return "未知动态"
}
