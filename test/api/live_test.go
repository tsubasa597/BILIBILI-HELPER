package api_test

import (
	"testing"

	"github.com/tsubasa597/BILIBILI-HELPER/api/live"
)

func TestLive(t *testing.T) {
	for _, up := range _ups {
		if _, err := live.LiverStatus(up); err != nil {
			t.Error(err)
		}
	}
}
