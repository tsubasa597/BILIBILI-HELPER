package task

import (
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/api/daily"
	"github.com/tsubasa597/BILIBILI-HELPER/api/user"
	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/BILIBILI-HELPER/state"
)

var (
	_ Tasker = (*Daily)(nil)
	// 银瓜子最低兑换要求
	exchangeRate int64 = 700
)

// Daily 日常任务
type Daily struct {
	VideoAvID string
	api       api.API
	state     state.State
}

// NewDaily 初始化
func NewDaily(api api.API, av string) Daily {
	return Daily{
		api:       api,
		VideoAvID: av,
		state:     state.Runing,
	}
}

// State 获取运行状态
func (d Daily) State() state.State {
	return d.state
}

// Run 运行日常任务
func (daily Daily) Run(ch chan<- interface{}) {
	if daily.state != state.Runing {
		return
	}

	if daily.VideoAvID == "" {
		daily.getRandomAV()
	}

	var res string
	if err, ok := daily.userCheck(); ok {
		res += "WatchVideo: " + daily.watchVideo() + "\n"
		res += "ShareVideo: " + daily.shareVideo() + "\n"
		res += "Sliver2Coins: " + daily.sliver2Coins() + "\n"
		res += "LiveCheckin: " + daily.liveCheckin()
	} else {
		res += "UserCheck: " + err
		daily.state = state.Stop
	}

	ch <- res
}

// Next 下次运行时间
func (daily Daily) Next(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}

func (d Daily) userCheck() (string, bool) {
	if _, err := user.UserCheck(d.api); err != nil {
		return err.Error(), false
	}

	return ecode.SucessLogin, true
}

func (d Daily) watchVideo() string {
	if d.VideoAvID == "" {
		return ecode.ErrNoBvID
	}

	if _, err := daily.WatchVideo(d.api, d.VideoAvID); err != nil {
		return err.Error()
	}

	return ecode.SucessPlay

}

func (d Daily) sliver2Coins() string {
	status, err := daily.Sliver2CoinsStatus(d.api)
	if err != nil {
		return err.Error()
	}

	if status.Data.Silver < exchangeRate {
		return ecode.ErrExchange
	}

	resp, err := daily.Sliver2Coins(d.api)
	if err != nil {
		return err.Error()
	}

	return resp.Message

}

func (d Daily) shareVideo() string {
	if d.VideoAvID == "" {
		return ecode.ErrNoBvID
	}

	if _, err := daily.ShareVideo(d.api, d.VideoAvID); err != nil {
		return err.Error()
	}

	return ecode.SucessShare
}

func (d Daily) liveCheckin() string {
	if _, err := daily.LiveCheckin(d.api); err != nil {
		return err.Error()
	}

	return ecode.SucessSignIn

}

func (d *Daily) getRandomAV() {
	s, err := daily.GetRandomAV(d.api)
	if err != nil || s == "" {
		d.VideoAvID = ""
	}

	d.VideoAvID = s
}
