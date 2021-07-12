package api

import (
	"encoding/json"
	fmt "fmt"

	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"github.com/tsubasa597/requests"
)

const (
	ErrGetDynamic    = "请求发生错误"
	ErrNoDynamic     = "该用户没有动态"
	ErrUnknowDynamic = "未知动态"
	ErrNotListen     = "该用户未监听"
	ErrRepeatListen  = "重复监听"
	ErrLoad          = "解析错误"
)

const (
	CommentViedo = iota + 1
	CommentTopic
	_
	CommentActivity
	_
	_
	CommentNotice
	CommentLiveActivity
	CommentActivityViedo
	CommentLiveNotice
	CommentDynamicImage
	CommentColumn
	_
	CommentAudio
	_
	_
	CommentDynamic
)

const (
	DynamicUnknow = iota
	DynamicOrig
	DynamicImage
	DynamicText    = (iota + 1)
	DynamicVideo   = iota << 1
	DynamicPost    = (iota - 1) << 4
	DynamicMusic   = (iota + 2) << 5
	DynamicAnime   = (iota + 1) << 6
	DynamichSketch = iota << 8
	DynamicLive    = 4200
	DynamicLive2   = 4308
)

// GetDynamicSrvSpaceHistory 获取目的 uid 的所有动态
func (api API) GetDynamicSrvSpaceHistory(hostUID int64) (*DynamicSvrSpaceHistoryResponse, error) {
	resp := &DynamicSvrSpaceHistoryResponse{}
	err := api.Req.Gets(fmt.Sprintf("%s?host_uid=%d", dynamicSrvSpaceHistory, hostUID), resp)

	return resp, err
}

func GetComments(type_, oid int) (*Comments, error) {
	resp := &Comments{}
	err := requests.Gets(fmt.Sprintf("%s?type=%d&oid=%d", reply, type_, oid), resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// GetOriginCard 获取 Card 的源动态
func GetOriginCard(c *Card) (info info.Dynamic) {
	info.T = c.Desc.Timestamp
	info.Name = c.Desc.UserProfile.Info.Uname

	switch c.Desc.Type {
	case DynamicUnknow:
		info.Err = fmt.Errorf(ErrUnknowDynamic)
		return
	case DynamicOrig:
		dynamic := &CardWithOrig{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info = GetOriginCard(&Card{
			Desc: &Card_Desc{
				Type:        dynamic.Item.OrigType,
				Timestamp:   c.Desc.Timestamp,
				UserProfile: c.Desc.UserProfile,
			},
			Card: dynamic.Origin,
		})
		info.Content = dynamic.Item.Content
		return
	case DynamicImage:
		dynamic := &CardWithImage{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		info.CommentType = CommentDynamicImage
		return
	case DynamicText:
		dynamic := &CardTextOnly{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		info.CommentType = CommentDynamic
		return
	case DynamicVideo:
		dynamic := &CardWithVideo{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		info.CommentType = CommentViedo
		return
	case DynamicPost:
		dynamic := &CardWithPost{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		return
	case DynamicMusic:
		dynamic := &CardWithMusic{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		info.CommentType = CommentAudio
		return
	case DynamicAnime:
		dynamic := &CardWithAnime{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		return
	case DynamichSketch:
		dynamic := &CardWithSketch{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		return
	case DynamicLive:
		dynamic := &CardWithLive{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		return
	case DynamicLive2:
		dynamic := &CardWithLiveV2{}
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
