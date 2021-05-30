package listen

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
)

type Dynamic struct {
	ups sync.Map
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
		return append(dynamics, getOriginCard(dynamicSvrSpaceHistoryResponse.Data.Cards[index]))
	}

	for int64(dynamicSvrSpaceHistoryResponse.Data.Cards[index].Desc.Timestamp) > t {
		dynamics = append(dynamics, getOriginCard(dynamicSvrSpaceHistoryResponse.Data.Cards[index]))
		index++
	}

	return dynamics
}

// GetOriginCard 获取 Card 的源动态
func getOriginCard(c *api.Card) (info info.Dynamic) {
	info.T = c.Desc.Timestamp
	info.Name = c.Desc.UserProfile.Info.Uname

	switch c.Desc.Type {
	case 0:
		info.Err = fmt.Errorf(ErrUnknowDynamic)
		return
	case 1:
		dynamic := &api.CardWithOrig{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info = getOriginCard(&api.Card{
			Desc: &api.Card_Desc{
				Type:        dynamic.Item.OrigType,
				Timestamp:   c.Desc.Timestamp,
				UserProfile: c.Desc.UserProfile,
			},
			Card: dynamic.Origin,
		})

		info.Content = dynamic.Item.Content
		return
	case 2:
		dynamic := &api.CardWithImage{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic

		return
	case 4:
		dynamic := &api.CardTextOnly{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		return
	case 8:
		dynamic := &api.CardWithVideo{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		return
	case 64:
		dynamic := &api.CardWithPost{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		return
	case 256:
		dynamic := &api.CardWithMusic{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		return
	case 512:
		dynamic := &api.CardWithAnime{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		return
	case 2048:
		dynamic := &api.CardWithSketch{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		return
	case 4200:
		dynamic := &api.CardWithLive{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		return
	case 4308:
		dynamic := &api.CardWithLiveV2{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		return
	}

	info.Err = fmt.Errorf(ErrLoad)
	return
}

var _ Listener = (*Dynamic)(nil)

func (dynamic *Dynamic) Listen(uid int64, a api.API) (infos []info.Infoer) {
	var dynamics []info.Dynamic

	if value, ok := dynamic.ups.Load(uid); ok {
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
	if _, ok := dynamic.ups.Load(uid); ok {
		dynamic.ups.Delete(uid)
		return nil
	} else {
		return fmt.Errorf(ErrNotListen)
	}
}

func (dynamic *Dynamic) GetList() (ups [][]string) {
	dynamic.ups.Range(func(key, value interface{}) bool {
		ups = append(ups, []string{value.(*UpRoutine).Name, fmt.Sprint(value.(*UpRoutine).Time)})
		return true
	})

	return ups
}

func (dynamic *Dynamic) Add(uid, t int64, api api.API, ctx context.Context, cancel context.CancelFunc) error {
	if _, ok := dynamic.ups.Load(uid); ok {
		return fmt.Errorf(ErrRepeatListen)
	}

	var name string

	if s, err := api.GetUserName(uid); err == nil {
		name = s
	}

	dynamic.ups.Store(uid, &UpRoutine{
		Ctx:    ctx,
		Cancel: cancel,
		Name:   name,
		Time:   t,
	})
	return nil
}

func NewDynamic() *Dynamic {
	return &Dynamic{
		ups: sync.Map{},
	}
}
