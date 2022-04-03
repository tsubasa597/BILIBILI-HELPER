package api_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tsubasa597/BILIBILI-HELPER/api/live"
)

func TestLive(t *testing.T) {
	assert := assert.New(t)

	for up, name := range _ups {
		if info, err := live.Status(up); err != nil {
			t.Error(err)
		} else {
			assert.Equal(name, info.Name)
		}
	}
}

func TestLiveRoom(t *testing.T) {
	t.SkipNow()

	info, err := live.Status(351609538)
	if err != nil {
		t.Error(info)
	}

	room, err := live.NewRoom(_api, info)
	if err != nil {
		t.Error(err)
	}

	resp, err := room.Enter()
	if err != nil {
		t.Error(err)
	}

	<-time.After(
		time.Until(
			time.Unix(resp.Timestamp, 0).Add(time.Second * time.Duration(resp.HeartbeatInterval)),
		),
	)

	if err := room.In(resp); err != nil {
		t.Error(err)
	}
}
