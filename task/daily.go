package task

import (
	"strings"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
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
	api       *api.API
	state     state.State
}

// NewDaily 初始化
func NewDaily(api *api.API, av string) Daily {
	return Daily{
		api:       api,
		VideoAvID: av,
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

	res := strings.Builder{}
	if err, ok := daily.userCheck(); ok {
		res.WriteString("WatchVideo: " + daily.watchVideo() + "\n")
		res.WriteString("ShareVideo: " + daily.shareVideo() + "\n")
		res.WriteString("Sliver2Coins: " + daily.sliver2Coins() + "\n")
		res.WriteString("LiveCheckin: " + daily.liveCheckin())
	} else {
		res.WriteString("UserCheck: " + err)
		daily.state = state.Stop
	}

	ch <- res.String()
	res.Reset()
}

// Next 下次运行时间
func (daily Daily) Next(time time.Time) time.Time {
	return time.AddDate(0, 0, 1)
}

func (daily Daily) userCheck() (string, bool) {
	if _, err := daily.api.UserCheck(); err != nil {
		return err.Error(), false
	}

	return ecode.SucessLogin, true
}

func (daily Daily) watchVideo() string {
	if daily.VideoAvID == "" {
		return ecode.ErrNoBvID
	}

	if _, err := daily.api.WatchVideo(daily.VideoAvID); err != nil {
		return err.Error()
	}

	return ecode.SucessPlay

}

func (daily Daily) sliver2Coins() string {
	status, err := daily.api.Sliver2CoinsStatus()
	if err != nil {
		return err.Error()
	}

	if status.Data.Silver < exchangeRate {
		return ecode.ErrExchange
	}

	resp, err := daily.api.Sliver2Coins()
	if err != nil {
		return err.Error()
	}

	return resp.Message

}

func (daily Daily) shareVideo() string {
	if daily.VideoAvID == "" {
		return ecode.ErrNoBvID
	}

	if _, err := daily.api.ShareVideo(daily.VideoAvID); err != nil {
		return err.Error()
	}

	return ecode.SucessShare
}

func (daily Daily) liveCheckin() string {
	if _, err := daily.api.LiveCheckin(); err != nil {
		return err.Error()
	}

	return ecode.SucessSignIn

}

func (daily *Daily) getRandomAV() {
	s, err := daily.api.GetRandomAV()
	if err != nil || s == "" {
		daily.VideoAvID = ""
	}
	daily.VideoAvID = s
}
