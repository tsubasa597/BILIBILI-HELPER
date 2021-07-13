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

// GetDynamicSrvSpaceHistory 获取目的 uid 的所有动态
func GetDynamicSrvSpaceHistory(hostUID int64) (*DynamicSvrSpaceHistoryResponse, error) {
	resp := &DynamicSvrSpaceHistoryResponse{}
	err := requests.Gets(fmt.Sprintf("%s?host_uid=%d", dynamicSrvSpaceHistory, hostUID), resp)

	return resp, err
}

func GetComments(commentType uint8, oid int64) (*Comments, error) {
	resp := &Comments{}
	err := requests.Gets(fmt.Sprintf("%s?type=%d&oid=%d", reply, commentType, oid), resp)
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
	case DynamicDescType_Unknown:
		info.Err = fmt.Errorf(ErrUnknowDynamic)
		return
	case DynamicDescType_WithOrigin:
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
				Rid:         c.Desc.Origin.Rid,
			},
			Card: dynamic.Origin,
		})
		info.Content = dynamic.Item.Content
		info.RID = c.Desc.Origin.Rid
		return
	case DynamicDescType_WithImage:
		dynamic := &CardWithImage{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		info.CommentType = CommentDynamicImage
		info.RID = c.Desc.Rid
		return
	case DynamicDescType_TextOnly:
		dynamic := &CardTextOnly{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		info.CommentType = CommentDynamic
		info.RID = c.Desc.Rid
		return
	case DynamicDescType_WithVideo:
		dynamic := &CardWithVideo{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		info.CommentType = CommentViedo
		info.RID = c.Desc.Rid
		return
	case DynamicDescType_WithPost:
		dynamic := &CardWithPost{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		info.RID = c.Desc.Rid
		return
	case DynamicDescType_WithMusic:
		dynamic := &CardWithMusic{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		info.CommentType = CommentAudio
		info.RID = c.Desc.Rid
		return
	case DynamicDescType_WithAnime:
		dynamic := &CardWithAnime{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		info.RID = c.Desc.Rid
		return
	case DynamicDescType_WithSketch:
		dynamic := &CardWithSketch{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		info.RID = c.Desc.Rid
		return
	case DynamicDescType_WithLive:
		dynamic := &CardWithLive{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		info.RID = c.Desc.Rid
		return
	case DynamicDescType_WithLiveV2:
		dynamic := &CardWithLiveV2{}
		err := json.Unmarshal([]byte(c.Card), dynamic)
		if err != nil {
			info.Err = err
			return
		}

		info.Card = dynamic
		info.RID = c.Desc.Rid
		return
	}
	info.Err = fmt.Errorf(ErrLoad)
	return
}
