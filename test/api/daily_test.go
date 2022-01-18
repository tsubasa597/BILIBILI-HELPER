package api_test

import (
	"testing"

	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/api/daily"
)

var (
	_bvs = []string{
		"BV1PN411X7QW",
		"BV1M64y1a7zh",
		"BV1ER4y1E7qn",
		"BV1f54y1j7X8",
		"BV1ER4y1E7qn",
	}
)

var (
	_api api.API
	_do  bool

	_ups = []int64{
		351609538,
		672328094,
		672353429,
		672346917,
		672342685,
	}
)

func TestMain(t *testing.T) {
	ap, err := api.New("./cookie.yaml")
	if err != nil {
		t.Error(err)
		t.Log("跳过测试")
		return
	}

	_api = ap
	_do = true
}

func TestWatchVideo(t *testing.T) {
	if !_do {
		t.SkipNow()
	}

	for _, bv := range _bvs {
		if err := daily.WatchVideo(_api, bv); err != nil {
			t.Error(err)
		}
	}
}

func TestShareVideo(t *testing.T) {
	if !_do {
		t.SkipNow()
	}

	if err := daily.ShareVideo(_api, _bvs[0]); err != nil {
		t.Error(err)
	}

}

func TestSliver2CoinsStatus(t *testing.T) {
	if !_do {
		t.SkipNow()
	}

	if _, err := daily.Sliver2CoinsStatus(_api); err != nil {
		t.Error(err)
	}
}

func TestSliver2Coins(t *testing.T) {
	if !_do {
		t.SkipNow()
	}

	if err := daily.Sliver2Coins(_api); err != nil {
		t.Error(err)
	}
}

func TestGetRandomAV(t *testing.T) {
	if !_do {
		t.SkipNow()
	}

	if _, err := daily.GetRandomAV(_api); err != nil {
		t.Error(err)
	}
}

func TestLiveCheckin(t *testing.T) {
	if !_do {
		t.SkipNow()
	}

	if err := daily.LiveCheckin(_api); err != nil {
		t.Error(err)
	}
}
