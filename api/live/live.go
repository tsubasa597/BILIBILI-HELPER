package live

import (
	"fmt"

	"github.com/tsubasa597/BILIBILI-HELPER/api/proto"
	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"github.com/tsubasa597/requests"
)

// LiverStatus 直播状态
func LiverStatus(uid int64) (*proto.GetRoomInfoOldResponse, error) {
	resp := &proto.GetRoomInfoOldResponse{}
	if err := requests.Gets(fmt.Sprintf("%s?mid=%d", info.GetRoomInfoOld, uid), resp); err != nil {
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
