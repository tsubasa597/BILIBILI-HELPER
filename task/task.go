package task

import "bili/utils"

// Tasker 任务
type Tasker interface {
	// LiveCheckin 直播签到
	liveCheckin()
	// UserCheck 用户检查
	userCheck()
	// Sliver2CoinsStatus 银瓜子换硬币状态
	sliver2CoinsStatus()
	// Sliver2Coins 银瓜子换硬币
	sliver2Coins()
	// VideoWatch 观看视频
	videoWatch(string)
	// GetJsonResponse 返回 JSONResponse
	getJSONResponse() utils.JSON
	videoShare(string)
}

// Response 返回 json 的结构
type Response struct {
	http utils.HTTP
	json utils.JSON
}

// GetJSONResponse 返回 JSONResponse
func (rs *Response) getJSONResponse() utils.JSON {
	return rs.json
}
