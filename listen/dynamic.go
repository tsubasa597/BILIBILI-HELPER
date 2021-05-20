package listen

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
)

type Dynamic struct {
	ups map[int64]*UpRoutine
}

// GetDynamicMessage 获取目标 uid 的第一条记录
func getDynamicMessage(hostUID int64, a api.API) (dynamic info.Dynamic) {
	dynamicSvrSpaceHistoryResponse, err := a.GetDynamicSrvSpaceHistory(hostUID)
	if err != nil {
		dynamic.Err = err
		return
	}

	if dynamicSvrSpaceHistoryResponse.Code != 0 {
		dynamic.Err = fmt.Errorf(errGetDynamic)
		return
	}

	if dynamicSvrSpaceHistoryResponse.Data.HasMore != 1 {
		dynamic.Err = fmt.Errorf(errNoDynamic)
		return
	}

	return getOriginCard(dynamicSvrSpaceHistoryResponse.Data.Cards[0])
}

// GetOriginCard 获取 Card 的源动态
func getOriginCard(c *api.Card) (info info.Dynamic) {
	info.T = c.Desc.Timestamp
	info.Name = c.Desc.UserProfile.Info.Uname

	switch c.Desc.Type {
	case 0:
		info.Err = fmt.Errorf(errUnknowDynamic)
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

	info.Err = fmt.Errorf(errLoad)
	return
}

var _ Listener = (*Dynamic)(nil)

func (dynamic Dynamic) Listen(uid int64, a api.API) info.Infoer {
	return getDynamicMessage(uid, a)
}

func (dynamic *Dynamic) StopListenUP(uid int64) error {
	if _, ok := dynamic.ups[uid]; ok {
		delete(dynamic.ups, uid)
		return nil
	} else {
		return fmt.Errorf("错误")
	}
}

func (dynamic Dynamic) GetList() string {
	var ups string
	for _, v := range dynamic.ups {
		ups += fmt.Sprintf("%s\n", v.Name)
	}
	return ups
}

func (dynamic *Dynamic) Add(uid int64, name string, ctx context.Context, cancel context.CancelFunc) error {
	if _, ok := dynamic.ups[uid]; ok {
		return fmt.Errorf("错误")
	}
	dynamic.ups[uid] = &UpRoutine{
		Ctx:    ctx,
		Cancel: cancel,
		Name:   name,
	}
	return nil
}

func NewDynamic() *Dynamic {
	return &Dynamic{
		ups: make(map[int64]*UpRoutine),
	}
}
