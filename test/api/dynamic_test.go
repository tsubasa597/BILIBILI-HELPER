package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsubasa597/BILIBILI-HELPER/api/dynamic"
)

func TestGetDynamics(t *testing.T) {
	assert := assert.New(t)

	for up, name := range _ups {
		dynamics, err := dynamic.Get(up, 0)
		if err != nil {
			t.Error(err)
			continue
		}

		assert.Equal(name, dynamics[0].Name)
	}
}
