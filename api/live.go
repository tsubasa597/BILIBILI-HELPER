package api

import (
	"encoding/json"
	"fmt"

	"github.com/tsubasa597/BILIBILI-HELPER/global"
)

func (api API) GetLiverStatus(uid int64) (*GetRoomInfoOldResponse, error) {
	rep, err := api.r.Get(fmt.Sprintf("%s?mid=%d", LiverRoomID, uid))
	if err != nil {
		return nil, err
	}

	resp := &GetRoomInfoOldResponse{}
	err = json.Unmarshal(rep, resp)
	return resp, err
}

func GetLiverStatus(uid int64) (*XSpaceAccInfoResponse, error) {
	rep, err := global.Get(fmt.Sprintf("%s?mid=%d", LiverStatus, uid))
	if err != nil {
		return nil, err
	}

	resp := &XSpaceAccInfoResponse{}
	err = json.Unmarshal(rep, resp)
	return resp, err
}
