package api

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"github.com/tsubasa597/requests"
)

const (
	// CommentViedo 视频
	CommentViedo = iota + 1
	// CommentTopic 话题
	CommentTopic
	_
	// CommentActivity 活动
	CommentActivity
	_
	_
	// CommentNotice 公告
	CommentNotice
	// CommentLiveActivity 直播活动
	CommentLiveActivity
	// CommentActivityViedo 活动稿件
	CommentActivityViedo
	// CommentLiveNotice 直播公告
	CommentLiveNotice
	// CommentDynamicImage 相簿（图片动态）
	CommentDynamicImage
	// CommentColumn 专栏
	CommentColumn
	_
	// CommentAudio 音频
	CommentAudio
	_
	_
	// CommentDynamic 动态（纯文字动态&分享）
	CommentDynamic
)

// GetDynamicSrvSpaceHistory 获取目的 uid 的所有动态
func GetDynamicSrvSpaceHistory(hostUID, nextOffect int64) (*DynamicSvrSpaceHistoryResponse, error) {
	resp := &DynamicSvrSpaceHistoryResponse{}
	err := requests.Gets(fmt.Sprintf("%s?visitor_uid=0&host_uid=%d&offset_dynamic_id=%d&platform=web",
		dynamicSrvSpaceHistory, hostUID, nextOffect), resp)

	if len(resp.Data.Cards) == 0 {
		return nil, fmt.Errorf(ecode.ErrNoDynamic)
	}

	if resp.Code != ecode.Sucess {
		return nil, fmt.Errorf(resp.Message)
	}

	return resp, err
}

// GetDynamic 获取指定时间前的动态
func GetDynamic(hostUID, t int64) (dynamics []info.Dynamic, err error) {
	var nextOffect int64
DynamicCards:
	for {
		history, err := GetDynamicSrvSpaceHistory(hostUID, nextOffect)
		if err != nil {
			return nil, err
		}

		for _, card := range history.Data.Cards {
			if int64(card.Desc.Timestamp) >= t {
				break DynamicCards
			}

			if dynamic, err := GetOriginCard(card); err == nil {
				dynamics = append(dynamics, dynamic)
			}
		}
	}

	return
}

// GetComments 获取评论
func GetComments(commentType, sort uint8, oid int64, ps, pn int) (*Comments, error) {
	resp := &Comments{}
	err := requests.Gets(fmt.Sprintf("%s?type=%d&oid=%d&sort=%d&ps=%d&pn=%d",
		reply, commentType, oid, sort, ps, pn), resp)
	if err != nil {
		return nil, err
	}

	if len(resp.Data.Replies) == 0 {
		return nil, fmt.Errorf(ecode.ErrNoComment)
	}

	if resp.Code != ecode.Sucess {
		return nil, fmt.Errorf(resp.Message)
	}

	return resp, nil
}

// GetOriginCard 获取 Card 的源动态
func GetOriginCard(c *Card) (dynamic info.Dynamic, err error) {
	dynamic.UID = c.Desc.Uid
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
		return info.Dynamic{}, fmt.Errorf(ecode.ErrUnknowDynamic)
	case DynamicDescType_WithOrigin:
		orig := &CardWithOrig{}
		err = json.Unmarshal([]byte(c.Card), orig)
		if err != nil {
			return
		}

		var dy info.Dynamic
		dy, err = GetOriginCard(&Card{
			Desc: &Card_Desc{
				Uid:          c.Desc.Uid,
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
		dynamic.RID = item.Item.Id
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

	case DynamicDescType_WithMiss:
		dynamic.CommentType = CommentDynamic

	default:
		err = fmt.Errorf(ecode.ErrLoad)
	}
	return
}
