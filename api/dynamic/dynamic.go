package dynamic

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api/proto"
	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"github.com/tsubasa597/requests"
)

var (
	dynamicPool *sync.Pool = &sync.Pool{
		New: func() interface{} {
			return &info.Dynamic{}
		},
	}
)

// GetDynamics 获取目的 uid 的一页动态
func GetDynamics(hostUID, nextOffect int64) ([]info.Dynamic, error) {
	resp := &proto.DynamicSvrSpaceHistoryResponse{}
	err := requests.Gets(fmt.Sprintf("%s?visitor_uid=0&host_uid=%d&offset_dynamic_id=%d&platform=web",
		info.DynamicSrvSpaceHistory, hostUID, nextOffect), resp)
	if err != nil {
		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	if resp.Code != ecode.Sucess {
		return nil, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: resp.Message,
		}
	}

	if len(resp.Data.Cards) == 0 {
		return nil, ecode.APIErr{
			E: ecode.ErrNoDynamic,
		}
	}

	dynamics := make([]info.Dynamic, 0, len(resp.Data.Cards))
	for _, card := range resp.Data.Cards {
		dynamic, err := GetOriginCard(card)
		if err != nil {
			continue
		}

		dynamics = append(dynamics, dynamic)
	}

	return dynamics, nil
}

// GetAllDynamics 获取指定时间为止的所有动态
func GetAllDynamics(hostUID, t int64) (dynamics []info.Dynamic) {
	var offect int64
	for {
		infos, err := GetDynamics(hostUID, offect)
		if err != nil {
			return
		}

		for _, inf := range infos {
			if inf.Time <= t {
				return
			}

			offect = inf.Offect
			dynamics = append(dynamics, inf)
		}

		time.Sleep(time.Second)
	}
}

// GetOriginCard 获取 Card 的源动态
func GetOriginCard(c *proto.Card) (info.Dynamic, error) {
	dynamic := *dynamicPool.Get().(*info.Dynamic)
	defer dynamicPool.Put(&dynamic)

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
	case proto.DynamicDescType_Unknown:
		return info.Dynamic{}, ecode.APIErr{
			E: ecode.ErrUnknowDynamic,
		}
	case proto.DynamicDescType_WithOrigin:
		orig := &proto.CardWithOrig{}
		err = json.Unmarshal([]byte(c.Card), orig)
		if err != nil {
			return dynamic, ecode.APIErr{
				E:   ecode.ErrLoad,
				Msg: err.Error(),
			}
		}

		var dy info.Dynamic
		dy, err = GetOriginCard(&proto.Card{
			Desc: &proto.Card_Desc{
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

	case proto.DynamicDescType_WithImage:
		dynamic.Type = info.CommentDynamicImage

		item := &proto.CardWithImage{}
		if err = json.Unmarshal([]byte(c.Card), item); err != nil {
			return dynamic, ecode.APIErr{
				E:   ecode.ErrLoad,
				Msg: err.Error(),
			}
		}
		dynamic.RID = item.Item.Id
		dynamic.Content = item.Item.Description

	case proto.DynamicDescType_TextOnly:
		dynamic.Type = info.CommentDynamic

		item := &proto.CardTextOnly{}
		if err = json.Unmarshal([]byte(c.Card), item); err != nil {
			return dynamic, ecode.APIErr{
				E:   ecode.ErrLoad,
				Msg: err.Error(),
			}
		}
		dynamic.Content = item.Item.Content

	case proto.DynamicDescType_WithVideo:
		dynamic.Type = info.CommentViedo
		dynamic.RID = c.Desc.Rid

	case proto.DynamicDescType_WithPost:
		dynamic.Type = info.CommentColumn
		dynamic.RID = c.Desc.Rid

	case proto.DynamicDescType_WithMusic:
		dynamic.Type = info.CommentAudio

	case proto.DynamicDescType_WithMiss:
		dynamic.Type = info.CommentDynamic

	default:
		return dynamic, ecode.APIErr{
			E: ecode.ErrLoad,
		}
	}

	return dynamic, nil
}
