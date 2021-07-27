package api

import (
	"encoding/json"
	fmt "fmt"
	"strconv"

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
func GetDynamicSrvSpaceHistory(hostUID, nextOffect int64) (*DynamicSvrSpaceHistoryResponse, error) {
	resp := &DynamicSvrSpaceHistoryResponse{}
	err := requests.Gets(fmt.Sprintf("%s?host_uid=%d&offset_dynamic_id=%d",
		dynamicSrvSpaceHistory, hostUID, nextOffect), resp)

	return resp, err
}

func GetComments(commentType uint8, oid int64, ps, pn int) (*Comments, error) {
	resp := &Comments{}
	err := requests.Gets(fmt.Sprintf("%s?type=%d&oid=%d&ps=%d&pn=%d",
		reply, commentType, oid, ps, pn), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetOriginCard 获取 Card 的源动态
func GetOriginCard(c *Card) (dynamic info.Dynamic, err error) {
	dynamic.Time = c.Desc.Timestamp
	dynamic.Name = c.Desc.UserProfile.Info.Uname
	dynamic.Card = c.Card

	i, err := strconv.Atoi(c.Desc.DynamicIdStr)
	if err != nil {
		return
	}
	dynamic.RID = int64(i)
	dynamic.DynamicID = int64(i)

	switch c.Desc.Type {
	case DynamicDescType_Unknown:
		return info.Dynamic{}, fmt.Errorf(ErrUnknowDynamic)
	case DynamicDescType_WithOrigin:
		orig := &CardWithOrig{}
		err = json.Unmarshal([]byte(c.Card), orig)
		if err != nil {
			return
		}

		var dy info.Dynamic
		dy, err = GetOriginCard(&Card{
			Desc: &Card_Desc{
				Type:         orig.Item.OrigType,
				Timestamp:    c.Desc.Timestamp,
				UserProfile:  c.Desc.UserProfile,
				DynamicIdStr: c.Desc.DynamicIdStr,
			},
			Card: orig.Origin,
		})
		if err != nil {
			return dy, err
		}

		dynamic.CommentType = CommentDynamic
		dynamic.Content = orig.Item.Content

	case DynamicDescType_WithImage:
		dynamic.CommentType = CommentDynamicImage

		item := &CardWithImage{}
		if err = json.Unmarshal([]byte(c.Card), item); err != nil {
			return
		}
		dynamic.Content = item.Item.Description

	case DynamicDescType_TextOnly:
		dynamic.CommentType = CommentDynamic

		item := &CardTextOnly{}
		if err = json.Unmarshal([]byte(c.Card), item); err != nil {
			return
		}
		dynamic.Content = item.Item.Content

	case DynamicDescType_WithVideo:
		dynamic.CommentType = CommentViedo
		dynamic.RID = c.Desc.Rid

	case DynamicDescType_WithPost:
		dynamic.CommentType = CommentColumn
		dynamic.RID = c.Desc.Rid

	case DynamicDescType_WithMusic:
		dynamic.CommentType = CommentAudio

	default:
		err = fmt.Errorf(ErrLoad)
	}
	return
}
