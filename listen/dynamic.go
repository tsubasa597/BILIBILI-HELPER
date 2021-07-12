package listen

import (
	"context"
	"fmt"
	"sync"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
)

type Dynamic struct {
	UPs sync.Map
}

// GetDynamicMessage 获取目标 uid 的指定记录
func getDynamicMessage(hostUID, t int64, a api.API) (dynamics []info.Dynamic) {
	var index int
	dynamicSvrSpaceHistoryResponse, err := a.GetDynamicSrvSpaceHistory(hostUID)
	if err != nil {
		return append(dynamics, info.Dynamic{
			Info: info.Info{
				Err: err,
			},
		})
	}

	if dynamicSvrSpaceHistoryResponse.Code != 0 {
		return append(dynamics, info.Dynamic{
			Info: info.Info{
				Err: fmt.Errorf(ErrGetDynamic),
			},
		})
	}

	if dynamicSvrSpaceHistoryResponse.Data.HasMore != 1 {
		return append(dynamics, info.Dynamic{
			Info: info.Info{
				Err: fmt.Errorf(ErrNoDynamic),
			},
		})
	}

	if t == int64(NewListen) {
		return append(dynamics, api.GetOriginCard(dynamicSvrSpaceHistoryResponse.Data.Cards[index]))
	}

	for int64(dynamicSvrSpaceHistoryResponse.Data.Cards[index].Desc.Timestamp) > t {
		dynamics = append(dynamics, api.GetOriginCard(dynamicSvrSpaceHistoryResponse.Data.Cards[index]))
		index++
	}

	return dynamics
}

var _ Listener = (*Dynamic)(nil)

func (dynamic *Dynamic) Listen(uid int64, a api.API) (infos []info.Infoer) {
	var dynamics []info.Dynamic

	if value, ok := dynamic.UPs.Load(uid); ok {
		dynamics = getDynamicMessage(uid, value.(*UpRoutine).Time, a)

		if len(dynamics) > 0 {
			value.(*UpRoutine).Time = int64(dynamics[0].T)
		}
	} else {
		dynamics = getDynamicMessage(uid, int64(NewListen), a)
	}

	for _, v := range dynamics {
		infos = append(infos, v)
	}

	return
}

func (dynamic *Dynamic) StopListenUP(uid int64) error {
	if _, ok := dynamic.UPs.Load(uid); ok {
		dynamic.UPs.Delete(uid)
		return nil
	} else {
		return fmt.Errorf(ErrNotListen)
	}
}

func (dynamic *Dynamic) GetList() (ups [][]string) {
	dynamic.UPs.Range(func(key, value interface{}) bool {
		ups = append(ups, []string{value.(*UpRoutine).Name, fmt.Sprint(value.(*UpRoutine).Time)})
		return true
	})

	return ups
}

func (dynamic *Dynamic) Add(uid, t int64, api api.API, ctx context.Context, cancel context.CancelFunc) error {
	if _, ok := dynamic.UPs.Load(uid); ok {
		return fmt.Errorf(ErrRepeatListen)
	}

	var name string

	if s, err := api.GetUserName(uid); err == nil {
		name = s
	}

	dynamic.UPs.Store(uid, &UpRoutine{
		Ctx:    ctx,
		Cancel: cancel,
		Name:   name,
		Time:   t,
	})
	return nil
}

func NewDynamic() *Dynamic {
	return &Dynamic{
		UPs: sync.Map{},
	}
}
