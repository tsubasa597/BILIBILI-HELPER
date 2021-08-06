package task

import "github.com/tsubasa597/BILIBILI-HELPER/api"

var _ Tasker = (*Daily)(nil)

type Daily struct {
	VideoAvID string
	api       api.API
}

func (daily *Daily) Init(api api.API) {
	daily.api = api
}

func (daily Daily) Run() (res string) {
	if daily.VideoAvID == "" {
		daily.getRandomAV()
	}

	if err, ok := daily.userCheck(); ok {
		res += "WatchVideo: " + daily.watchVideo() + "\n"
		res += "ShareVideo: " + daily.shareVideo() + "\n"
		res += "Sliver2Coins: " + daily.sliver2Coins() + "\n"
		res += "LiveCheckin: " + daily.liveCheckin()
	} else {
		res += "UserCheck: " + err
	}
	return
}

func (daily Daily) userCheck() (string, bool) {
	resp, err := daily.api.UserCheck()
	if err != nil {
		return err.Error(), false
	}

	if resp.Code == 0 {
		return "登录成功", true
	}

	return resp.Message, false
}

func (daily Daily) watchVideo() string {
	if daily.VideoAvID == "" {
		return "Bvid 为空，跳过"
	}

	resp, err := daily.api.WatchVideo(daily.VideoAvID)
	if err != nil && resp.Code != 0 {
		return err.Error()
	}

	if resp.Code == 0 {
		return "播放成功"
	}

	return resp.Message
}

func (daily Daily) sliver2Coins() string {
	const exchangeRate int64 = 700
	status, err := daily.api.Sliver2CoinsStatus()
	if err != nil {
		return err.Error()
	}

	if status.Data.Silver < exchangeRate {
		return "当前银瓜子余额不足700,不进行兑换"
	}

	resp, err := daily.api.Sliver2Coins()

	if resp.Code == 0 {
		return "兑换成功"
	}

	if resp.Code == 403 {
		return resp.Message
	}

	if err != nil {
		return err.Error()
	}

	return resp.Message
}

func (daily Daily) shareVideo() string {
	if daily.VideoAvID == "" {
		return "Bvid 为空，跳过"
	}

	resp, err := daily.api.ShareVideo(daily.VideoAvID)
	if err != nil && resp.Code != 0 {
		return err.Error()
	}

	if resp.Code == 0 {
		return "分享成功"
	}

	return resp.Message
}

func (daily Daily) liveCheckin() string {
	resp, err := daily.api.LiveCheckin()
	if err != nil {
		return err.Error()
	}

	if resp.Code == 0 {
		return "签到成功"
	}

	return "重复签到"
}

func (daily *Daily) getRandomAV() {
	s, err := daily.api.GetRandomAV()
	if err != nil || s == "" {
		daily.VideoAvID = ""
	}
	daily.VideoAvID = s
}
