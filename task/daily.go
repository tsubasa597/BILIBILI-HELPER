package task

import (
	"fmt"
	"sync"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/api/daily"
	"github.com/tsubasa597/BILIBILI-HELPER/api/user"
	"github.com/tsubasa597/BILIBILI-HELPER/ecode"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
)

// Daily 日常任务
type Daily struct {
	VideoAvID string
	api       api.API
	timeCell  time.Duration
	rwMutex   *sync.RWMutex
}

var (
	_ Tasker = (*Daily)(nil)
	// 银瓜子最低兑换要求
	_exchangeRate int64 = 700
)

// NewDaily 初始化
// 默认时间间隔为 24 小时
func NewDaily(api api.API, av string) Daily {
	return Daily{
		api:       api,
		VideoAvID: av,
		timeCell:  time.Hour * 24, /** 24 小时 */
		rwMutex:   &sync.RWMutex{},
	}
}

// Run 运行日常任务
func (daily Daily) Run(ch chan<- interface{}, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	daily.rwMutex.Lock()
	defer daily.rwMutex.Unlock()

	if daily.VideoAvID == "" {
		daily.getRandomAV()
	}

	var res string
	if err, ok := daily.userCheck(); ok {
		res = fmt.Sprintf(
			`WatchVideo: %s
			ShareVideo: %s 
			Sliver2Coins: %s
			LiveCheckin: %s`,
			daily.watchVideo(), daily.shareVideo(), daily.sliver2Coins(), daily.liveCheckin())
	} else {
		res = fmt.Sprintf("UserCheck: %s", err)
	}

	ch <- res
}

// Next 下次运行时间
func (daily Daily) Next(t time.Time) time.Time {
	daily.rwMutex.Lock()
	defer daily.rwMutex.Unlock()

	return t.Add(daily.timeCell)
}

// Info 返回任务的信息
func (daily Daily) Info() info.Task {
	daily.rwMutex.RLock()
	defer daily.rwMutex.RUnlock()

	return info.Task{
		TimeCell: daily.timeCell,
	}
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

	if status.Data.Silver < _exchangeRate {
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
