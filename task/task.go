package task

// Tasker 任务
type Tasker interface {
	// LiveCheckin 直播签到
	LiveCheckin()
	// UserCheck 用户检查
	UserCheck()
	// Sliver2CoinsStatus 银瓜子换硬币状态
	Sliver2CoinsStatus()
	// Sliver2Coins 银瓜子换硬币
	Sliver2Coins()
	// VideoWatch 观看视频
	VideoWatch(string)
	// GetJsonResponse 返回 JSONResponse
	GetJSONResponse() *JSONResponse
}

// JSONResponse 返回 json 的结构
type JSONResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	TTL     int                    `json:"ttl"`
	Data    map[string]interface{} `json:"data"`
}

// GetJSONResponse 返回 JSONResponse
func (rs *JSONResponse) GetJSONResponse() *JSONResponse {
	return rs
}
