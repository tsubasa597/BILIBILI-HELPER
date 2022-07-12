package cron

import (
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
func NewDaily(api api.API, av string, log *zap.Logger) *Daily {
	if log == nil {
		log = zap.NewExample()
	}

	return &Daily{
		api:       api,
		VideoAvID: av,
		timeCell:  time.Hour * 24, /** 24 小时 */
		log:       log,
		mutex:     &sync.Mutex{},
	}
}

// Run 运行日常任务
func (d *Daily) Run() interface{} {
	var (
		err error
	)

	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.VideoAvID == "" {
		d.VideoAvID, err = daily.GetRandomAV(d.api)
		if err != nil {
			d.log.Error(err.Error())
		}
	}

	err = user.Check(d.api)
	if err == nil {
		return daily.Info{
			WatchVideo:   daily.WatchVideo(d.api, d.VideoAvID).Error(),
			ShareVideo:   daily.ShareVideo(d.api, d.VideoAvID).Error(),
			Sliver2Coins: daily.Sliver2Coins(d.api).Error(),
			LiveCheckin:  daily.LiveCheckin(d.api).Error(),
		}
	}

	d.log.Error(err.Error())
	return nil
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
