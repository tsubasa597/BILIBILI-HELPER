package live

import (
	"fmt"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/api/proto"
	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/requests"
)

// Info 监听的直播信息
type Info struct {
	Name    string
	Title   string
	Status  bool
	RoomURL string
}

// LiverStatus 直播状态
func LiverStatus(uid int64) (Info, error) {
	resp := &proto.XSpaceAccInfoResponse{}
	if err := requests.Gets(fmt.Sprintf("%s?mid=%d", api.SpaceAccInfo, uid), resp); err != nil {
		return Info{}, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	if resp.Code != ecode.Sucess {
		return Info{}, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: resp.Message,
		}
	}

	info := Info{
		Name:    resp.Data.Name,
		Title:   resp.Data.LiveRoom.Title,
		RoomURL: resp.Data.LiveRoom.Url,
	}

	if resp.Data.LiveRoom.LiveStatus == proto.LiveStatus_Living {
		info.Status = true
	}

	return info, nil
}
