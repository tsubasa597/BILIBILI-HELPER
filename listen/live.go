package listen

import (
	"context"
	"fmt"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
)

type Live struct {
	ups map[int64]*UpRoutine
}

func getLiverStatus(uid int64, api api.API) (info info.Live) {
	rep, err := api.GetUserInfo(uid)
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

var _ Listener = (*Live)(nil)

func (live Live) Listen(uid int64, api api.API) info.Infoer {
	return getLiverStatus(uid, api)
}

func (live *Live) StopListenUP(uid int64) error {
	if _, ok := live.ups[uid]; ok {
		delete(live.ups, uid)
		return nil
	} else {
		return fmt.Errorf("错误")
	}
}

func (live Live) GetList() string {
	var ups string
	for _, v := range live.ups {
		ups += fmt.Sprintf("%s\n", v.Name)
	}
	return ups
}

func (live *Live) Add(uid int64, name string, ctx context.Context, cancel context.CancelFunc) error {
	if _, ok := live.ups[uid]; ok {
		return fmt.Errorf("错误")
	}
	live.ups[uid] = &UpRoutine{
		Ctx:    ctx,
		Cancel: cancel,
		Name:   name,
	}
	return nil
}

func NewLive() *Live {
	return &Live{
		ups: make(map[int64]*UpRoutine),
	}
}
