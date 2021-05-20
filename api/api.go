package api

import (
	fmt "fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/requests"
)

type API struct {
	conf     cookie
	Requests *requests.Requests
	Entry    *logrus.Entry
}

func New(c cookie, enrty *logrus.Entry) API {
	r := requests.New()
	r.SetHeader(http.Header{
		"Connection":   []string{"keep-alive"},
		"User-Agent":   []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.70"},
		"Cookie":       []string{c.getVerify()},
		"Content-Type": []string{"application/x-www-form-urlencoded"},
	})

	if enrty == nil {
		enrty = logrus.NewEntry(newLog())
	}

	return API{
		Requests: r,
		conf:     c,
		Entry:    enrty,
	}
}

func (api API) GetUserInfo(uid int64) (*XSpaceAccInfoResponse, error) {
	resp := &XSpaceAccInfoResponse{}
	err := api.Requests.Gets(fmt.Sprintf("%s?mid=%d", spaceAccInfo, uid), resp)

	return resp, err
}

// UserCheck 用户登录验证
func (api API) UserCheck() (*BaseResponse, error) {
	resp := &BaseResponse{}
	err := api.Requests.Gets(userLogin, resp)

	return resp, err
}

// WatchVideo 视频模拟观看，观看时间在 [0, 90) 之间
func (api API) WatchVideo(bvid string) (*BaseResponse, error) {
	data := url.Values{
		"bvid":        []string{bvid},
		"played_time": []string{strconv.Itoa(rand.Intn(90))},
	}

	resp := &BaseResponse{}
	err := api.Requests.Posts(videoHeartbeat, data, resp)

	return resp, err
}

// ShareVideo 分享视频
func (api API) ShareVideo(bvid string) (*BaseResponse, error) {
	data := url.Values{
		"bvid": []string{bvid},
		"csrf": []string{api.conf.BiliJct},
	}

	resp := &BaseResponse{}
	err := api.Requests.Posts(avShare, data, resp)

	return resp, err
}

// Sliver2CoinsStatus 获取银瓜子和硬币的数量
func (api API) Sliver2CoinsStatus() (*Sliver2CoinsStatusResponse, error) {
	resp := &Sliver2CoinsStatusResponse{}
	err := api.Requests.Gets(sliver2CoinsStatus, resp)

	return resp, err
}

// Sliver2Coins 将银瓜子兑换为硬币
func (api API) Sliver2Coins() (*BaseResponse, error) {
	resp := &BaseResponse{}
	err := api.Requests.Gets(sliver2Coins, resp)

	return resp, err
}

// LiveCheckin 直播签到
func (api API) LiveCheckin() (*BaseResponse, error) {
	resp := &BaseResponse{}
	err := api.Requests.Gets(liveCheckin, resp)

	return resp, err
}

// GetDynamicSrvSpaceHistory 获取目的 uid 的所有动态
func (api API) GetDynamicSrvSpaceHistory(hostUID int64) (*DynamicSvrSpaceHistoryResponse, error) {
	resp := &DynamicSvrSpaceHistoryResponse{}
	err := api.Requests.Gets(fmt.Sprintf("%s?host_uid=%d", dynamicSrvSpaceHistory, hostUID), resp)

	return resp, err
}

func (api API) LiverStatus(uid int64) (*GetRoomInfoOldResponse, error) {
	resp := &GetRoomInfoOldResponse{}
	err := api.Requests.Gets(fmt.Sprintf("%s?mid=%d", getRoomInfoOld, uid), resp)

	return resp, err
}

func (api API) GetUserName(uid int64) (string, error) {
	info, err := api.GetUserInfo(uid)
	if err != nil {
		return "", err
	}

	return info.Data.Name, nil
}
