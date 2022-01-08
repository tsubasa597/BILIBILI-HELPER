package cron

import (
	"fmt"
	"sync"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/api/daily"
	"github.com/tsubasa597/BILIBILI-HELPER/api/user"
	"go.uber.org/zap"
)

// Daily 日常任务
type Daily struct {
	VideoAvID string
	api       api.API
	timeCell  time.Duration
	log       *zap.Logger
	mutex     *sync.Mutex
}

var (
	_ Tasker = (*Daily)(nil)
)

// NewDaily 初始化
// 默认时间间隔为 24 小时
func NewDaily(api api.API, av string, log *zap.Logger) Daily {
	return Daily{
		api:       api,
		VideoAvID: av,
		timeCell:  time.Hour * 24, /** 24 小时 */
		log:       log,
		mutex:     &sync.Mutex{},
	}
}

// Run 运行日常任务
func (d *Daily) Run(ch chan<- interface{}, wg *sync.WaitGroup) {
	var (
		err error
	)

	defer func() {
		wg.Done()
	}()

	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.VideoAvID == "" {
		d.VideoAvID, err = daily.GetRandomAV(d.api)
		if err != nil {
			d.log.Error(err.Error())
		}
	}

	if err := user.UserCheck(d.api); err == nil {
		d.log.Info(fmt.Sprintf(
			`WatchVideo: %s
			ShareVideo: %s 
			Sliver2Coins: %s
			LiveCheckin: %s`,
			daily.WatchVideo(d.api, d.VideoAvID).Error(),
			daily.ShareVideo(d.api, d.VideoAvID).Error(),
			daily.Sliver2Coins(d.api).Error(),
			daily.LiveCheckin(d.api).Error(),
		))
	} else {
		d.log.Error(err.Error())
	}
}

// Next 下次运行时间
func (daily Daily) Next(t time.Time) time.Time {
	return t.Add(daily.timeCell)
}

// Info 返回任务的信息
func (daily Daily) Info() Info {
	return Info{
		TimeCell: daily.timeCell,
	}
}
