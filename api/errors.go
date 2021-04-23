package api

type RequestErr struct{}
type DynamicUnknownErr struct{}
type LoadErr struct{}

var _ error = (*DynamicUnknownErr)(nil)
var _ error = (*RequestErr)(nil)
var _ error = (*LoadErr)(nil)

func (r RequestErr) Error() string {
	return "请求错误"
}

func (d DynamicUnknownErr) Error() string {
	return "未知动态"
}

func (l LoadErr) Error() string {
	return "解析错误"
}
