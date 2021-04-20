package task

import (
	"errors"
	"sync"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/global"
)

type Bili struct {
	api.API

	info []string
	errs []error
	wg   *sync.WaitGroup
}

func New(c global.Cookie) *Bili {
	return &Bili{
		API: *api.New(c),

		info: make([]string, 0),
		errs: make([]error, 0),
		wg:   &sync.WaitGroup{},
	}
}

func Run(b *Bili) ([]string, []error) {
	if ok := b.userCheck(); ok {
		b.wg.Add(1)
		go b.watchVideo("BV1NT4y137Jc")
		time.Sleep(time.Second)

		b.wg.Add(1)
		go b.shareVideo("BV1NT4y137Jc")
		time.Sleep(time.Second)

		b.wg.Add(1)
		go b.sliver2Coins()
		time.Sleep(time.Second)

		b.wg.Add(1)
		go b.liveCheckin()
		time.Sleep(time.Second)
	}

	b.wg.Wait()

	return b.info, b.errs
}

func (b *Bili) userCheck() bool {
	resp, err := b.UserCheck()
	if err != nil {
		b.errs = append(b.errs, err)
		return false
	}

	if resp.Code == 0 {
		return true
	}

	b.errs = append(b.errs, errors.New(resp.Message))
	return false
}

func (b *Bili) watchVideo(bvid string) {
	defer b.wg.Done()

	resp, err := b.WatchVideo(bvid)
	if err != nil {
		b.errs = append(b.errs, err)
		return
	}

	if resp.Code == 0 {
		b.info = append(b.info, "播放成功")
		return
	}

	b.errs = append(b.errs, errors.New(resp.Message))
}

func (b *Bili) sliver2Coins() {
	defer b.wg.Done()

	const exchangeRate int64 = 700
	status, err := b.Sliver2CoinsStatus()
	if err != nil {
		b.errs = append(b.errs, err)
		return
	}

	if status.Data.Silver < exchangeRate {
		b.errs = append(b.errs, errors.New("当前银瓜子余额不足700,不进行兑换"))
		return
	}

	resp, err := b.Sliver2Coins()

	if resp.Code == 0 {
		b.info = append(b.info, "兑换成功")
		return
	}

	if resp.Code == 403 {
		b.errs = append(b.errs, errors.New(resp.Message))
		return
	}

	if err != nil {
		b.errs = append(b.errs, err)
		return
	}

	b.errs = append(b.errs, errors.New(resp.Message))
}

func (b *Bili) shareVideo(bvid string) {
	defer b.wg.Done()

	resp, err := b.ShareVideo(bvid)
	if err != nil {
		b.errs = append(b.errs, err)
		return
	}

	if resp.Code == 0 {
		b.info = append(b.info, "分享成功")
		return
	}

	b.errs = append(b.errs, errors.New(resp.Message))
}

func (b *Bili) liveCheckin() {
	defer b.wg.Done()

	resp, err := b.LiveCheckin()
	if err != nil {
		b.errs = append(b.errs, err)
		return
	}

	if resp.Code == 0 {
		b.info = append(b.info, resp.Message)
		return
	}

	b.errs = append(b.errs, errors.New("重复签到"))
}
