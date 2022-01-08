package api_test

import (
	"testing"

	"github.com/tsubasa597/BILIBILI-HELPER/api/dynamic"
)

var (
	_ups = []int64{
		351609538,
		672328094,
		672353429,
		672346917,
		672342685,
	}
)

func TestGetDynamics(t *testing.T) {
	for _, up := range _ups {
		dynamics, err := dynamic.GetDynamics(up, 0)
		if err != nil {
			t.Error(err)
			continue
		}

		t.Log(dynamics[0].Content)
	}
}
