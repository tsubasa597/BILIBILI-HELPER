package dynamic

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/api/comment"
	"github.com/tsubasa597/BILIBILI-HELPER/api/proto"
	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/requests"
)

// Dynamic 监听的动态信息
type Info struct {
	Name    string
	Time    int64
	UID     int64
	Content string
	Card    string
	RID     int64
	Offect  int64
	Type    comment.Type
}

var (
	_dynamicPool *sync.Pool = &sync.Pool{
		New: func() interface{} {
			return &Info{}
		},
	}
)

// Get 获取目的 uid 的一页动态
func Get(hostUID, nextOffect int64) ([]Info, error) {
	resp := &proto.DynamicSvrSpaceHistoryResponse{}
	err := requests.Gets(fmt.Sprintf("%s?visitor_uid=0&host_uid=%d&offset_dynamic_id=%d&platform=web",
		api.DynamicSrvSpaceHistory, hostUID, nextOffect), resp)
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

	dynamics := make([]Info, 0, len(resp.Data.Cards))
	for _, card := range resp.Data.Cards {
		dynamic, err := getOriginCard(card)
		if err != nil {
			continue
		}

		dynamics = append(dynamics, dynamic)
	}

	return dynamics, nil
}

// GetAll 获取指定时间为止的所有动态
func GetAll(hostUID, t int64) (dynamics []Info, errs []error) {
	var offect int64
	for {
		infos, err := Get(hostUID, offect)
		if err != nil {
			errs = append(errs, ecode.APIErr{
				E:   ecode.ErrGetInfo,
				Msg: err.Error(),
			})
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

// getOriginCard 解析 Card 的动态内容
func getOriginCard(c *proto.Card) (Info, error) {
	dynamic := *_dynamicPool.Get().(*Info)
	defer _dynamicPool.Put(&dynamic)

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
		return Info{}, ecode.APIErr{
			E:   ecode.ErrLoad,
			Msg: ecode.MsgUnknowDynamic,
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

		var dy Info
		dy, err = getOriginCard(&proto.Card{
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

		dynamic.Type = comment.Dynamic
		dynamic.Content = orig.Item.Content

	case proto.DynamicDescType_WithImage:
		dynamic.Type = comment.DynamicImage

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
		dynamic.Type = comment.Dynamic

		item := &proto.CardTextOnly{}
		if err = json.Unmarshal([]byte(c.Card), item); err != nil {
			return dynamic, ecode.APIErr{
				E:   ecode.ErrLoad,
				Msg: err.Error(),
			}
		}
		dynamic.Content = item.Item.Content

	case proto.DynamicDescType_WithVideo:
		dynamic.Type = comment.Viedo
		dynamic.RID = c.Desc.Rid

	case proto.DynamicDescType_WithPost:
		dynamic.Type = comment.Column
		dynamic.RID = c.Desc.Rid

	case proto.DynamicDescType_WithMusic:
		dynamic.Type = comment.Audio

	case proto.DynamicDescType_WithMiss:
		dynamic.Type = comment.Dynamic

	default:
		return dynamic, ecode.APIErr{
			E: ecode.ErrLoad,
		}
	}

	return dynamic, nil
}
