package cron_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/api/daily"
	"github.com/tsubasa597/BILIBILI-HELPER/api/dynamic"
	"github.com/tsubasa597/BILIBILI-HELPER/api/user"
	"github.com/tsubasa597/BILIBILI-HELPER/cron"
	"go.uber.org/zap"
)

var (
	_ti  int64       = time.Now().AddDate(0, 0, -1).Unix()
	_log *zap.Logger = zap.NewExample()

	_dynamic = map[int64]*cron.Dynamic{
		351609538: cron.NewDynamic(351609538, _ti, 5*time.Minute, _log),
		672328094: cron.NewDynamic(672328094, _ti, 5*time.Minute, _log),
		672353429: cron.NewDynamic(672353429, _ti, 5*time.Minute, _log),
		672346917: cron.NewDynamic(672346917, _ti, 5*time.Minute, _log),
	}
	// _comment = map[int64]
)

func TestDynamic(t *testing.T) {
	var (
		now    = time.Now()
		assert = assert.New(t)
	)

	c := cron.New[[]dynamic.Info]()
	c.Start()

	for id, task := range _dynamic {
		c.Add(id, task, now)
		time.Sleep(time.Second)
	}

	assert.Equal(4, len(c.Info()))

	c.StopByID(672328094)
	assert.Equal(3, len(c.Info()))

	c.Add(672342685, cron.NewDynamic(672342685, _ti, 5*time.Minute, _log), time.Now())
	assert.Equal(4, len(c.Info()))

	<-c.Ch

	c.Stop()
}

func TestDaily(t *testing.T) {
	t.SkipNow()

	ap, err := api.New("../cookie.yaml")
	if err != nil {
		t.Logf("%s, %s", "跳过测试", err)
		return
	}

	task := cron.NewDaily(ap, "", nil)
	c := cron.New[daily.Info]()

	c.Start()
	c.Add(1, task, time.Now())

	t.Log(<-c.Ch)
}

func TestUser(t *testing.T) {
	assert := assert.New(t)

	c := cron.New[user.Info]()
	c.Start()

	task := cron.NewUser(351609538, time.Minute, nil)
	c.Add(351609538, task, time.Now())

	info := <-c.Ch
	assert.Equal(int64(0), info.UID)

	c.Stop()
}
