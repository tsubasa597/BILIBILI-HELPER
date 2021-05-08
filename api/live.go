package api

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/global"
)

func (api API) LiverStatus(uid int64) (*XSpaceAccInfoResponse, error) {
	rep, err := global.Get(fmt.Sprintf("%s?mid=%d", LiverStatus, uid))
	if err != nil {
		return nil, err
	}

	resp := &XSpaceAccInfoResponse{}
	err = json.Unmarshal(rep, resp)
	return resp, err
}

func (api API) GetLiverStatus(uid int64) (info Info) {
	rep, err := api.LiverStatus(uid)
	if err != nil {
		info.Err = err
		return
	}

	info.Name = rep.Data.Name
	info.T = int32(time.Now().Unix())
	info.LiveRoomURL = rep.Data.LiveRoom.Url
	info.LiveTitle = rep.Data.LiveRoom.Title
	if rep.Data.LiveRoom.RoomStatus == 1 {
		info.LiveStatus = true
	}
	return
}

func (l Listen) LiveListen(uid int64) (context.Context, chan Info, error) {
	return l.listen(uid, l.api.GetLiverStatus)
}
