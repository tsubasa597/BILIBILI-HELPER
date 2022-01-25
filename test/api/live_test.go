package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsubasa597/BILIBILI-HELPER/api/live"
)

func TestLive(t *testing.T) {
	assert := assert.New(t)

	for up, name := range _ups {
		if info, err := live.LiverStatus(up); err != nil {
			t.Error(err)
		} else {
			assert.Equal(name, info.Name)
		}
	}
}
