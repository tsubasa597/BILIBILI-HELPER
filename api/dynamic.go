package api

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// GetDynamicMessage 获取目标 uid 的第一条记录
func (l *Listen) GetDynamicMessage(hostUID int64) Info {
	dynamicSvrSpaceHistoryResponse, err := l.api.GetDynamicSrvSpaceHistory(hostUID)
	if err != nil {
		l.api.entry.Debugln(err)
		return Info{
			Err: err,
		}
	}

	if dynamicSvrSpaceHistoryResponse.Code != 0 {
		l.api.entry.Debugln(dynamicSvrSpaceHistoryResponse.Message)
		return Info{
			Err: fmt.Errorf(errGetDynamic),
		}
	}

	if dynamicSvrSpaceHistoryResponse.Data.HasMore != 1 {
		l.api.entry.Debugln(errNoDynamic)
		return Info{
			Err: fmt.Errorf(errNoDynamic),
		}
	}

	return getOriginCard(dynamicSvrSpaceHistoryResponse.Data.Cards[0])
}

// GetOriginCard 获取 Card 的源动态
func getOriginCard(c *Card) (info Info) {
	info.T = c.Desc.Timestamp
	info.Name = c.Desc.UserProfile.Info.Uname

	switch c.Desc.Type {
	case 0:
		info.Err = fmt.Errorf(errUnknowDynamic)
		return
	case 1:
		dynamic := &CardWithOrig{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info = getOriginCard(&Card{
			Desc: &Card_Desc{
				Type:        dynamic.Item.OrigType,
				Timestamp:   c.Desc.Timestamp,
				UserProfile: c.Desc.UserProfile,
			},
			Card: dynamic.Origin,
		})

		info.Content = dynamic.Item.Content
		return
	case 2:
		dynamic := &CardWithImage{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic

		return
	case 4:
		dynamic := &CardTextOnly{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		return
	case 8:
		dynamic := &CardWithVideo{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		return
	case 64:
		dynamic := &CardWithPost{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		return
	case 256:
		dynamic := &CardWithMusic{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		return
	case 512:
		dynamic := &CardWithAnime{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		return
	case 2048:
		dynamic := &CardWithSketch{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		return
	case 4200:
		dynamic := &CardWithLive{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		return
	case 4308:
		dynamic := &CardWithLiveV2{}
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

func (l *Listen) DynamicListen(uid int64, ticker *time.Ticker) (context.Context, <-chan Info, error) {
	return l.listen(ticker.C, uid, l.GetDynamicMessage)
}
