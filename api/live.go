package api

import fmt "fmt"

// LiveCheckin 直播签到
func (api API) LiveCheckin() (*BaseResponse, error) {
	resp := &BaseResponse{}
	err := api.Req.Gets(liveCheckin, resp)

	return resp, err
}

func (api API) LiverStatus(uid int64) (*GetRoomInfoOldResponse, error) {
	resp := &GetRoomInfoOldResponse{}
	err := api.Req.Gets(fmt.Sprintf("%s?mid=%d", getRoomInfoOld, uid), resp)

	return resp, err
}
