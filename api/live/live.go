package live

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/api/proto"
	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/requests"
)

// SecretKey 中加密算法对应值
const (
	_md5 = iota
	_sha1
	_sha256
	_sha224
	_sha512
	_sha384
)

// Info 监听的直播信息
type Info struct {
	Name         string
	Title        string
	Status       proto.LiveStatus
	RoomID       int64
	RoomURL      string
	AreaId       int32
	ParentAreaId int32
	RoomId       int32
}

type Room struct {
	// parentAreaId 父分区号
	parentAreaId int32

	// areaId 分区号
	areaId int32

	// id 房间号
	id int32

	api api.API

	// uuid uuid4
	uuid string

	// seq 序列号
	seq int32

	// param 请求公共字段
	param url.Values
}

// LiverStatus 直播间状态
func Status(uid int64) (Info, error) {
	resp := &proto.XSpaceAccInfoResponse{}
	if err := requests.Gets(fmt.Sprintf("%s?mid=%d", api.SpaceAccInfo, uid), resp); err != nil {
		return Info{}, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	if resp.Code != ecode.Sucess {
		return Info{}, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: resp.Message,
		}
	}

	roomInfo := &proto.RoomInfo{}
	if err := requests.Gets(
		fmt.Sprintf("%s?room_id=%d", api.LiveInfo, resp.Data.LiveRoom.Roomid),
		roomInfo,
	); err != nil {
		return Info{}, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	if resp.Code != ecode.Sucess {
		return Info{}, ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: resp.Message,
		}
	}

	info := Info{
		Name:         resp.Data.Name,
		Title:        resp.Data.LiveRoom.Title,
		RoomURL:      resp.Data.LiveRoom.Url,
		RoomID:       resp.Data.LiveRoom.Roomid,
		Status:       roomInfo.Data.LiveStatus,
		AreaId:       roomInfo.Data.AreaId,
		ParentAreaId: roomInfo.Data.ParentAreaId,
		RoomId:       roomInfo.Data.RoomId,
	}

	return info, nil
}

// NewRoom 初始化直播间信息
func NewRoom(ap api.API, info Info) (*Room, error) {
	u := strings.ReplaceAll(uuid.NewString(), "-", "")
	return &Room{
		api:          ap,
		areaId:       info.AreaId,
		parentAreaId: info.ParentAreaId,
		id:           info.RoomId,
		uuid:         u,
		param: url.Values{
			"device": []string{
				fmt.Sprintf(
					"[\"%s\", \"%s\"]",
					ap.GetBuvid(), u),
			},
			"ua":         []string{"Mozilla/5.0 (X11; Linux x86_64; rv:98.0) Gecko/20100101 Firefox/98.0"},
			"csrf_token": []string{ap.GetJwt()},
			"csrf":       []string{ap.GetJwt()},
			"visit_id":   []string{},
		},
	}, nil
}

// Enter 进入直播间
func (room *Room) Enter() (*proto.EnterLive_Data, error) {
	param := room.param
	param.Set("id", room.buildID())
	param.Set("is_patch", "0")
	param.Set("heart_beat", "")
	param.Set("ts", fmt.Sprint(time.Now().UnixMilli()))

	resp := &proto.EnterLive{}
	if err := requests.Posts(api.EnterRoom, param, resp); err != nil {
		return nil, err
	}

	room.seq++
	return resp.Data, nil
}

// In 在直播间中发送心跳
func (room *Room) In(enter *proto.EnterLive_Data) error {
	param := room.param
	param.Set("id", room.buildID())
	param.Set("ts", fmt.Sprint(time.Now().UnixMilli()))
	param.Set("time", fmt.Sprint(enter.HeartbeatInterval))
	param.Set("ets", fmt.Sprint(enter.Timestamp))
	param.Set("benchmark", enter.SecretKey)

	s, err := room.clacSign(param, enter.SecretRule, []byte(enter.SecretKey))
	if err != nil {
		return ecode.APIErr{
			E:   ecode.ErrLoad,
			Msg: err.Error(),
		}
	}
	param.Set("s", s)

	resp := &proto.BaseResponse{}
	if err := requests.Posts(api.InRoom, param, resp); err != nil {
		return ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: err.Error(),
		}
	}

	if resp.Code != ecode.Sucess {
		return ecode.APIErr{
			E:   ecode.ErrGetInfo,
			Msg: resp.Message,
		}
	}

	room.seq++

	return nil
}

// clacSign 构建心跳请求所需字段
func (room Room) clacSign(param url.Values, rules []int32, key []byte) (string, error) {
	// 构建加密字符串
	var builder strings.Builder
	builder.WriteString("{")
	builder.WriteString("\"platform\":\"web\",")
	builder.WriteString(fmt.Sprintf("\"parent_id\":%d,", room.parentAreaId))
	builder.WriteString(fmt.Sprintf("\"area_id\":%d,", room.areaId))
	builder.WriteString(fmt.Sprintf("\"seq_id\":%d,", room.seq))
	builder.WriteString(fmt.Sprintf("\"room_id\":%d,", room.id))
	builder.WriteString(fmt.Sprintf("\"buvid\":\"%s\",", room.api.GetBuvid()))
	builder.WriteString(fmt.Sprintf("\"uuid\":\"%s\",", room.uuid))
	builder.WriteString(fmt.Sprintf("\"ets\":%s,", param.Get("ets")))
	builder.WriteString(fmt.Sprintf("\"time\":%s,", param.Get("time")))
	builder.WriteString(fmt.Sprintf("\"ts\":%s", param.Get("ts")))
	builder.WriteString("}")
	b := builder.String()

	var h hash.Hash
	for _, rule := range rules {
		switch rule {
		case _md5:
			h = hmac.New(md5.New, key)
		case _sha1:
			h = hmac.New(sha1.New, key)
		case _sha256:
			h = hmac.New(sha256.New, key)
		case _sha224:
			h = hmac.New(sha256.New224, key)
		case _sha512:
			h = hmac.New(sha512.New, key)
		case _sha384:
			h = hmac.New(sha512.New384, key)
		}

		if _, err := h.Write([]byte(b)); err != nil {
			return "", err
		}
		b = hex.EncodeToString(h.Sum([]byte{}))

		h.Reset()
	}

	return b, nil
}

// buildID 构建 id 字段
func (room Room) buildID() string {
	return fmt.Sprintf("[%d,%d,%d,%d]", room.parentAreaId, room.areaId, room.seq, room.id)
}
