package listen

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/e"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"github.com/tsubasa597/BILIBILI-HELPER/state"
)

// Dynamic 所有 uid 的状态
type Dynamic struct {
	ups state.ListenStater
}

var _ Listener = (*Dynamic)(nil)

const (
	//NewListen 新监听动态的时间，防止返回大量数据
	NewListen = iota + 1
)

// ListenInfo 返回从现在到指定时间为止所有 uid 的动态。若没有监听指定 uid，则返回空
func (dynamic *Dynamic) ListenInfo(uid int64) ([]info.Interface, error) {
	return dynamic.ups.Do(uid, getDynamicMessage)
}

// StopListenUP 停止监听 uid。若没有监听，则返回错误
func (dynamic *Dynamic) StopListenUP(uid int64) error {
	return dynamic.ups.StopOne(uid)
}

// GetList 获取监听的信息
func (dynamic *Dynamic) GetList() []state.Info {
	return dynamic.ups.GetAll()
}

// Add 添加 uid 进行监听
func (dynamic *Dynamic) Add(ctx context.Context, cancel context.CancelFunc, uid int64, t int32) error {
	var name string

	if s, err := api.GetUserName(uid); err == nil {
		name = s
	}

	return dynamic.ups.Put(uid, state.NewUpRoutine(ctx, cancel, t, name))
}

// GetState 获取状态
func (dynamic *Dynamic) GetState() state.State {
	return dynamic.ups.GetState()
}

// NewDynamic 初始化
func NewDynamic(ctx context.Context, log *logrus.Entry) *Dynamic {
	return &Dynamic{
		ups: state.NewDeListenState(ctx, log),
	}
}

// GetDynamicMessage 获取目标 uid 的指定记录
func getDynamicMessage(uid int64, l state.Listener, log *logrus.Entry) (dynamics []info.Interface) {
	var offect int64
	for {
		dynamicSvrSpaceHistoryResponse, err := api.GetDynamicSrvSpaceHistory(uid, offect)
		if err != nil {
			log.Errorln(err)
			return
		}

		if dynamicSvrSpaceHistoryResponse.Code != 0 {
			log.Errorln(e.ErrGetDynamic)
			return
		}

		if dynamicSvrSpaceHistoryResponse.Data.HasMore != 1 {
			log.Errorln(e.ErrNoDynamic)
			return
		}

		if l.GetInfo().Time == NewListen {
			info, err := api.GetOriginCard(dynamicSvrSpaceHistoryResponse.Data.Cards[0])
			if err != nil {
				log.Errorln(err)
				return
			}

			return append(dynamics, &info)
		}

		for _, card := range dynamicSvrSpaceHistoryResponse.Data.Cards {
			if card.Desc.Timestamp > l.GetInfo().Time {
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
