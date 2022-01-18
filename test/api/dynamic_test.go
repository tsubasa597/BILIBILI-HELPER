package api_test

import (
	"testing"

	"github.com/tsubasa597/BILIBILI-HELPER/api/dynamic"
)

func TestGetDynamics(t *testing.T) {
	for _, up := range _ups {
		_, err := dynamic.Get(up, 0)
		if err != nil {
			t.Error(err)
			continue
		}
	}
}
