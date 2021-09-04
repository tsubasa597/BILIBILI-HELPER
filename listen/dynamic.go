package listen

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
)

// Dynamic 所有 uid 的状态
type Dynamic struct {
	ups Infoer
	log *logrus.Entry
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
func (dynamic *Dynamic) Listen(uid int64) (infos []info.Infoer) {
	val, err := dynamic.ups.Get(uid)
	if err != nil {
		return getDynamicMessage(uid, NewListen, dynamic.log)
	}

	if val.GetState() != StateRuning {
		dynamic.log.Infof("uid: %d, name:%s 正在暂停", uid, val.Name)
		return nil
	}

	infos = getDynamicMessage(uid, val.Time, dynamic.log)

	if len(infos) > 0 {
		val.Time = infos[0].(*info.Dynamic).Time
	}

	return infos
}

// StopListenUP 停止监听 uid。若没有监听，则返回错误
func (dynamic *Dynamic) StopListenUP(uid int64) error {
	return dynamic.ups.StopOne(uid)
}

// GetList 获取监听的信息
func (dynamic *Dynamic) GetList() []*UpRoutine {
	return dynamic.ups.GetAll()
}

// Add 添加 uid 进行监听
func (dynamic *Dynamic) Add(ctx context.Context, cancel context.CancelFunc, uid int64, t int32) error {
	if _, err := dynamic.ups.Get(uid); err == nil {
		return fmt.Errorf(errRepeatListen)
	}

	var name string

	if s, err := api.GetUserName(uid); err == nil {
		name = s
	}

	dynamic.ups.Put(uid, NewUpRoutine(ctx, cancel, StateRuning, t, name))
	return nil
}

// NewDynamic 初始化
func NewDynamic(log *logrus.Entry) *Dynamic {
	return &Dynamic{
		ups: NewDefaultInfos(),
		log: log,
	}
}
