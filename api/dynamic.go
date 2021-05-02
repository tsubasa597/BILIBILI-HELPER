package api

import (
	"encoding/json"
	"fmt"

	"github.com/tsubasa597/BILIBILI-HELPER/global"
)

// GetDynamicMessage 获取目标 uid 的第一条记录
//
// return : 时间戳、Card 实例、错误信息
func GetDynamicMessage(hostUID int64) (int32, interface{}, error) {
	dynamicSvrSpaceHistoryResponse, err := GetDynamicSrvSpaceHistory(hostUID)
	if err != nil {
		return -1, nil, err
	}

	if dynamicSvrSpaceHistoryResponse.Code != 0 {
		return -1, nil, fmt.Errorf("请求发生错误: " + dynamicSvrSpaceHistoryResponse.Message)
	}

	c, e := GetOriginCard(dynamicSvrSpaceHistoryResponse.Data.Cards[0])
	if err != nil {
		return -1, nil, e
	}
	return c[0].(int32), c[1], nil
}

// GetOriginCard 获取 Card 的源动态
//
// return []interface{} : 时间戳、Card 实例、错误信息
func GetOriginCard(c *Card) ([]interface{}, error) {
	switch c.Desc.Type {
	case 0:
		return []interface{}{-1, nil}, fmt.Errorf("未知动态")
	case 1:
		dynamic := &CardWithOrig{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			return []interface{}{-1, nil}, err
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
			return []interface{}{-1, nil}, err
		}

		return []interface{}{c.Desc.Timestamp, dynamic}, nil
	case 4:
		dynamic := &CardTextOnly{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			return []interface{}{-1, nil}, err
		}

		return []interface{}{c.Desc.Timestamp, dynamic}, nil
	case 8:
		dynamic := &CardWithVideo{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			return []interface{}{-1, nil}, err
		}

		return []interface{}{c.Desc.Timestamp, dynamic}, nil
	case 64:
		dynamic := &CardWithPost{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			return []interface{}{-1, nil}, err
		}

		return []interface{}{c.Desc.Timestamp, dynamic}, nil
	case 256:
		dynamic := &CardWithMusic{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			return []interface{}{-1, nil}, err
		}

		return []interface{}{c.Desc.Timestamp, dynamic}, nil
	case 512:
		dynamic := &CardWithAnime{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			return []interface{}{-1, nil}, err
		}

		return []interface{}{c.Desc.Timestamp, dynamic}, nil
	case 2048:
		dynamic := &CardWithSketch{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			return []interface{}{-1, nil}, err
		}

		return []interface{}{c.Desc.Timestamp, dynamic}, nil
	case 4200:
		dynamic := &CardWithLive{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			return []interface{}{-1, nil}, err
		}

		return []interface{}{c.Desc.Timestamp, dynamic}, nil
	case 4308:
		dynamic := &CardWithLiveV2{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			return []interface{}{-1, nil}, err
		}

		return []interface{}{c.Desc.Timestamp, dynamic}, nil
	}

	return []interface{}{-1, nil}, fmt.Errorf("解析错误")
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
