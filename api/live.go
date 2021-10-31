package api

import (
	"fmt"

	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
)

// LiveCheckin 直播签到
func (api API) LiveCheckin() (*BaseResponse, error) {
	resp := &BaseResponse{}
	if err := api.Req.Gets(liveCheckin, resp); err != nil {
		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	if resp.Code != ecode.Sucess {
		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: resp.Message,
		}
	}

	return resp, nil
}

// LiverStatus 直播状态
func (api API) LiverStatus(uid int64) (*GetRoomInfoOldResponse, error) {
	resp := &GetRoomInfoOldResponse{}
	if err := api.Req.Gets(fmt.Sprintf("%s?mid=%d", getRoomInfoOld, uid), resp); err != nil {
		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	if resp.Code != ecode.Sucess {
		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: resp.Message,
		}
	}

	return resp, nil
}
