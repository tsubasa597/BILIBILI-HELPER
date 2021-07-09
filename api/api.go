package api

import (
	"encoding/json"
	fmt "fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/BILIBILI-HELPER/log"
	"github.com/tsubasa597/requests"
)

type API struct {
	cookie cookie
	Entry  *logrus.Entry
	Req    requests.Requests
}

func New(path string, enrty *logrus.Entry) API {
	c := newCookie(path)

	if enrty == nil {
		enrty = logrus.NewEntry(log.NewLog())
	}

	return API{
		cookie: c,
		Entry:  enrty,
		Req: requests.Requests{
			Client: &http.Client{},
			Headers: map[string]string{
				"Connection":   "keep-alive",
				"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.70",
				"Content-Type": "application/x-www-form-urlencoded",
			},
			Cookies: map[string]string{
				"DedeUserID": c.UserID,
				"SESSDATA":   c.SessData,
				"bili_jct":   c.BiliJct,
			},
		},
	}
}

func (api API) GetUserInfo(uid int64) (*XSpaceAccInfoResponse, error) {
	resp := &XSpaceAccInfoResponse{}
	err := api.Req.Gets(fmt.Sprintf("%s?mid=%d", spaceAccInfo, uid), resp)

	return resp, err
}

// UserCheck 用户登录验证
func (api API) UserCheck() (*BaseResponse, error) {
	resp := &BaseResponse{}
	err := api.Req.Gets(userLogin, resp)

	return resp, err
}

// WatchVideo 视频模拟观看，观看时间在 [0, 90) 之间
func (api API) WatchVideo(bvid string) (*BaseResponse, error) {
	data := url.Values{
		"bvid":        []string{bvid},
		"played_time": []string{strconv.Itoa(rand.Intn(90))},
	}

	resp := &BaseResponse{}
	err := api.Req.Posts(videoHeartbeat, data, resp)

	return resp, err
}

// ShareVideo 分享视频
func (api API) ShareVideo(bvid string) (*BaseResponse, error) {
	data := url.Values{
		"bvid": []string{bvid},
		"csrf": []string{api.cookie.BiliJct},
	}

	resp := &BaseResponse{}
	err := api.Req.Posts(avShare, data, resp)

	return resp, err
}

// Sliver2CoinsStatus 获取银瓜子和硬币的数量
func (api API) Sliver2CoinsStatus() (*Sliver2CoinsStatusResponse, error) {
	resp := &Sliver2CoinsStatusResponse{}
	err := api.Req.Gets(sliver2CoinsStatus, resp)

	return resp, err
}

// Sliver2Coins 将银瓜子兑换为硬币
func (api API) Sliver2Coins() (*BaseResponse, error) {
	resp := &BaseResponse{}
	err := api.Req.Gets(sliver2Coins, resp)

	return resp, err
}

// LiveCheckin 直播签到
func (api API) LiveCheckin() (*BaseResponse, error) {
	resp := &BaseResponse{}
	err := api.Req.Gets(liveCheckin, resp)

	return resp, err
}

// GetDynamicSrvSpaceHistory 获取目的 uid 的所有动态
func (api API) GetDynamicSrvSpaceHistory(hostUID int64) (*DynamicSvrSpaceHistoryResponse, error) {
	resp := &DynamicSvrSpaceHistoryResponse{}
	err := api.Req.Gets(fmt.Sprintf("%s?host_uid=%d", dynamicSrvSpaceHistory, hostUID), resp)

	return resp, err
}

func (api API) LiverStatus(uid int64) (*GetRoomInfoOldResponse, error) {
	resp := &GetRoomInfoOldResponse{}
	err := api.Req.Gets(fmt.Sprintf("%s?mid=%d", getRoomInfoOld, uid), resp)

	return resp, err
}

func (api API) GetUserName(uid int64) (string, error) {
	info, err := api.GetUserInfo(uid)
	if err != nil {
		return "", err
	}

	return info.Data.Name, nil
}

func (api API) GetRandomAV() (string, error) {
	resp := &RandomAvResponse{}
	err := api.Req.Gets(randomAV, resp)
	if err != nil {
		return "", err
	}
	parms := strings.Split(resp.Data.Url, "/")
	if strings.HasPrefix(parms[len(parms)-1], "BV") {
		return parms[len(parms)-1], nil
	}
	return "", nil
}

func GetComments() (*Comments, error) {
	resp := &Comments{}
	rep, err := requests.Get(reply)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(rep, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
