package api

import (
	"encoding/json"
	"fmt"

	"github.com/tsubasa597/BILIBILI-HELPER/global"
)

type Info struct {
	T    int32
	Err  error
	Card interface{}
}

// GetDynamicMessage 获取目标 uid 的第一条记录
func GetDynamicMessage(hostUID int64) Info {
	dynamicSvrSpaceHistoryResponse, err := GetDynamicSrvSpaceHistory(hostUID)
	if err != nil {
		return Info{
			T:   -1,
			Err: err,
		}
	}

	if dynamicSvrSpaceHistoryResponse.Code != 0 {
		return Info{
			T:   -1,
			Err: fmt.Errorf("请求发生错误: " + dynamicSvrSpaceHistoryResponse.Message),
		}
	}

	return GetOriginCard(dynamicSvrSpaceHistoryResponse.Data.Cards[0])
}

// GetOriginCard 获取 Card 的源动态
func GetOriginCard(c *Card) (info Info) {
	info = Info{
		T: c.Desc.Timestamp,
	}
	switch c.Desc.Type {
	case 0:
		info.Err = fmt.Errorf("未知动态")
		return
	case 1:
		dynamic := &CardWithOrig{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		return GetOriginCard(&Card{
			Desc: &Card_Desc{
				Type: dynamic.Item.OrigType,
			},
			Card: dynamic.Origin,
		})
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

	info.Err = fmt.Errorf("解析错误")
	return
}

// GetDynamicSrvSpaceHistory 获取目的 uid 的所有动态
func GetDynamicSrvSpaceHistory(hostUID int64) (*DynamicSvrSpaceHistoryResponse, error) {
	rep, err := global.Get(fmt.Sprintf("%s?host_uid=%d", DynamicSrvSpaceHistory, hostUID))
	if err != nil {
		return nil, err
	}

	resp := &DynamicSvrSpaceHistoryResponse{}

	err = json.Unmarshal(rep, &resp)

	return resp, err
}