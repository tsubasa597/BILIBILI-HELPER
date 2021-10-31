package api

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"github.com/tsubasa597/requests"
)

// GetDynamics 获取目的 uid 的一页动态
func GetDynamics(hostUID, nextOffect int64) ([]*Card, error) {
	resp := &DynamicSvrSpaceHistoryResponse{}
	if err := requests.Gets(fmt.Sprintf("%s?visitor_uid=0&host_uid=%d&offset_dynamic_id=%d&platform=web",
		dynamicSrvSpaceHistory, hostUID, nextOffect), resp); err != nil {

		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	if len(resp.Data.Cards) == 0 {
		return nil, ecode.APIErr{
			E: ecode.ErrNoDynamic,
		}
	}

	if resp.Code != ecode.Sucess {
		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: resp.Message,
		}
	}

	return resp.Data.Cards, nil
}

// GetAllDynamics 获取指定时间为止的所有动态
func GetAllDynamics(hostUID, t int64) (dynamics []*info.Dynamic) {
	var offect int64
	for {
		cards, err := GetDynamics(hostUID, offect)
		if err != nil {
			return
		}

		for _, card := range cards {
			if card.Desc.Timestamp <= t {
				return
			}

			dy, err := GetOriginCard(card)
			if err != nil {
				continue
			}

			offect = dy.Offect
			dynamics = append(dynamics, &dy)
		}
	}
}

// GetOriginCard 获取 Card 的源动态
func GetOriginCard(c *Card) (dynamic info.Dynamic, err error) {
	dynamic.UID = c.Desc.Uid
	dynamic.Time = c.Desc.Timestamp
	dynamic.Name = c.Desc.UserProfile.Info.Uname
	dynamic.Card = c.Card

	offect, err := strconv.Atoi(c.Desc.DynamicIdStr)
	if err != nil {
		return dynamic, ecode.APIErr{
			E:   ecode.ErrLoad,
			Msg: err.Error(),
		}
	}
	dynamic.RID = int64(offect)
	dynamic.Offect = int64(offect)

	switch c.Desc.Type {
	case DynamicDescType_Unknown:
		return info.Dynamic{}, ecode.APIErr{
			E: ecode.ErrUnknowDynamic,
		}
	case DynamicDescType_WithOrigin:
		orig := &CardWithOrig{}
		err = json.Unmarshal([]byte(c.Card), orig)
		if err != nil {
			return dynamic, ecode.APIErr{
				E:   ecode.ErrLoad,
				Msg: err.Error(),
			}
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

		dynamic.Type = info.CommentDynamic
		dynamic.Content = orig.Item.Content

	case DynamicDescType_WithImage:
		dynamic.Type = info.CommentDynamicImage

		item := &CardWithImage{}
		if err = json.Unmarshal([]byte(c.Card), item); err != nil {
			return dynamic, ecode.APIErr{
				E:   ecode.ErrLoad,
				Msg: err.Error(),
			}
		}
		dynamic.RID = item.Item.Id
		dynamic.Content = item.Item.Description

	case DynamicDescType_TextOnly:
		dynamic.Type = info.CommentDynamic

		item := &CardTextOnly{}
		if err = json.Unmarshal([]byte(c.Card), item); err != nil {
			return dynamic, ecode.APIErr{
				E:   ecode.ErrLoad,
				Msg: err.Error(),
			}
		}
		dynamic.Content = item.Item.Content

	case DynamicDescType_WithVideo:
		dynamic.Type = info.CommentViedo
		dynamic.RID = c.Desc.Rid

	case DynamicDescType_WithPost:
		dynamic.Type = info.CommentColumn
		dynamic.RID = c.Desc.Rid

	case DynamicDescType_WithMusic:
		dynamic.Type = info.CommentAudio

	case DynamicDescType_WithMiss:
		dynamic.Type = info.CommentDynamic

	default:
		return dynamic, ecode.APIErr{
			E: ecode.ErrLoad,
		}
	}

	return dynamic, nil
}
