package listen

import (
	"context"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
)

type Dynamic struct {
	UPs sync.Map
}

// GetDynamicMessage 获取目标 uid 的指定记录
func getDynamicMessage(hostUID int64, t int32, log *logrus.Entry) (dynamics []info.Infoer) {
	var offect int64
	for {

		dynamicSvrSpaceHistoryResponse, err := api.GetDynamicSrvSpaceHistory(hostUID, offect)
		if err != nil {
			log.Errorln(err)
			return
		}

		if dynamicSvrSpaceHistoryResponse.Code != 0 {
			log.Errorln(ErrGetDynamic)
			return
		}

		if dynamicSvrSpaceHistoryResponse.Data.HasMore != 1 {
			log.Errorln(ErrNoDynamic)
			return
		}

		if t == NewListen {
			info, err := api.GetOriginCard(dynamicSvrSpaceHistoryResponse.Data.Cards[0])
			if err != nil {
				log.Errorln(err)
				return
			}

			return append(dynamics, &info)
		}

		for _, card := range dynamicSvrSpaceHistoryResponse.Data.Cards {
			if card.Desc.Timestamp > t {
				info, err := api.GetOriginCard(card)
				if err != nil {
					log.Errorln(err)
					continue
				}
				offect = info.DynamicID
				dynamics = append(dynamics, &info)

			} else {
				return
			}
		}
	}
}

var _ Listener = (*Dynamic)(nil)

func (dynamic *Dynamic) Listen(uid int64, _ api.API, log *logrus.Entry) (infos []info.Infoer) {
	if value, ok := dynamic.UPs.Load(uid); ok {
		infos = getDynamicMessage(uid, value.(*UpRoutine).Time, log)

		if len(infos) > 0 {
			value.(*UpRoutine).Time = infos[0].(*info.Dynamic).Time
		}
	} else {
		return getDynamicMessage(uid, NewListen, log)
	}

	return
}

func (dynamic *Dynamic) StopListenUP(uid int64) error {
	if val, ok := dynamic.UPs.Load(uid); ok {
		dynamic.UPs.Delete(uid)
		val.(*UpRoutine).Cancel()
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

func (dynamic *Dynamic) Add(uid int64, t int32, _ api.API, ctx context.Context, cancel context.CancelFunc) error {
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
