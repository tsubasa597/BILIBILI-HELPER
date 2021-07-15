package task

import "github.com/tsubasa597/BILIBILI-HELPER/api"

var _ Tasker = (*Daily)(nil)

type Daily struct {
	VideoAvID string
}

func (daily Daily) Run(api api.API) (res string) {
	if daily.VideoAvID == "" {
		daily.VideoAvID = getRandomAV(api)
	}

	if err, ok := userCheck(api); ok {
		res += "WatchVideo: " + watchVideo(api, daily.VideoAvID) + "\n"
		res += "ShareVideo: " + shareVideo(api, daily.VideoAvID) + "\n"
		res += "Sliver2Coins: " + sliver2Coins(api) + "\n"
		res += "LiveCheckin: " + liveCheckin(api)
	} else {
		res += "UserCheck: " + err
	}
	return
}

func userCheck(api api.API) (string, bool) {
	resp, err := api.UserCheck()
	if err != nil {
		return err.Error(), false
	}

	if resp.Code == 0 {
		return "登录成功", true
	}

	return resp.Message, false
}

func watchVideo(api api.API, bvid string) string {
	if bvid == "" {
		return "Bvid 为空，跳过"
	}

	resp, err := api.WatchVideo(bvid)
	if err != nil && resp.Code != 0 {
		return err.Error()
	}

	if resp.Code == 0 {
		return "播放成功"
	}

	return resp.Message
}

func sliver2Coins(api api.API) string {
	const exchangeRate int64 = 700
	status, err := api.Sliver2CoinsStatus()
	if err != nil {
		return err.Error()
	}

	if status.Data.Silver < exchangeRate {
		return "当前银瓜子余额不足700,不进行兑换"
	}

	resp, err := api.Sliver2Coins()

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

func shareVideo(api api.API, bvid string) string {
	if bvid == "" {
		return "Bvid 为空，跳过"
	}

	resp, err := api.ShareVideo(bvid)
	if err != nil && resp.Code != 0 {
		return err.Error()
	}

	if resp.Code == 0 {
		return "分享成功"
	}

	return resp.Message
}

func liveCheckin(api api.API) string {
	resp, err := api.LiveCheckin()
	if err != nil {
		return err.Error()
	}

	if resp.Code == 0 {
		return "签到成功"
	}

	return "重复签到"
}

func getRandomAV(api api.API) string {
	s, err := api.GetRandomAV()
	if err != nil || s == "" {
		return ""
	}
	return s
}
