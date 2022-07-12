package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsubasa597/BILIBILI-HELPER/api/user"
)

func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)

	for up, name := range _ups {
		if info, err := user.GetInfo(up); err != nil {
			t.Error(err)
		} else {
			assert.Equal(name, info.Data.Name)
		}

	}
}

func TestUserCheck(t *testing.T) {
	if !_do {
		t.SkipNow()
	}

	if err := user.Check(_api); err != nil {
		t.Error(err)
	}

}
