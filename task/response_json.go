package task

// Response 返回 json 的结构
type Response struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	TTL     int                    `json:"ttl"`
	Data    map[string]interface{} `json:"data"`
}
