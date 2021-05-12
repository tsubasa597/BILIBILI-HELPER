package listen

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
)

// GetDynamicMessage 获取目标 uid 的第一条记录
func (l *Listen) GetDynamicMessage(hostUID int64) api.Info {
	dynamicSvrSpaceHistoryResponse, err := l.api.GetDynamicSrvSpaceHistory(hostUID)
	if err != nil {
		l.api.Entry.Debugln(err)
		return api.Info{
			Err: err,
		}
	}

	if dynamicSvrSpaceHistoryResponse.Code != 0 {
		l.api.Entry.Debugln(dynamicSvrSpaceHistoryResponse.Message)
		return api.Info{
			Err: fmt.Errorf(errGetDynamic),
		}
	}

	if dynamicSvrSpaceHistoryResponse.Data.HasMore != 1 {
		l.api.Entry.Debugln(errNoDynamic)
		return api.Info{
			Err: fmt.Errorf(errNoDynamic),
		}
	}

	return getOriginCard(dynamicSvrSpaceHistoryResponse.Data.Cards[0])
}

// GetOriginCard 获取 Card 的源动态
func getOriginCard(c *api.Card) (info api.Info) {
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

func (l *Listen) DynamicListen(uid int64) (context.Context, <-chan api.Info, error) {
	return l.AddListen(uid, l.GetDynamicMessage)
}
