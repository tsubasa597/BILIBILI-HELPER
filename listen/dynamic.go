package listen

import (
	"context"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
)

// Dynamic 所有 uid 的状态 sync.Map -> map[int64]UpRoutine
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
			log.Errorln(errGetDynamic)
			return
		}

		if dynamicSvrSpaceHistoryResponse.Data.HasMore != 1 {
			log.Errorln(errNoDynamic)
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

// Listen 返回从现在到指定时间为止所有 uid 的动态。若没有监听指定 uid，则返回第一条记录。
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

// StopListenUP 停止监听 uid。若没有监听，则返回错误
func (dynamic *Dynamic) StopListenUP(uid int64) error {
	if val, ok := dynamic.UPs.Load(uid); ok {
		dynamic.UPs.Delete(uid)
		val.(*UpRoutine).Cancel()
		return nil
	}

	return fmt.Errorf(errNotListen)
}

// GetList 获取正在监听的状态。returns []string{Name, Time}
func (dynamic *Dynamic) GetList() (ups [][2]string) {
	dynamic.UPs.Range(func(key, value interface{}) bool {
		ups = append(ups, [2]string{value.(*UpRoutine).Name, fmt.Sprint(value.(*UpRoutine).Time)})
		return true
	})

	return ups
}

// Add 添加 uid 进行监听
func (dynamic *Dynamic) Add(ctx context.Context, cancel context.CancelFunc, uid int64, t int32, _ api.API) error {
	if _, ok := dynamic.UPs.Load(uid); ok {
		return fmt.Errorf(errRepeatListen)
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

// NewDynamic 初始化
func NewDynamic() *Dynamic {
	return &Dynamic{
		UPs: sync.Map{},
	}
}
