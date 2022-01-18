package api_test

import (
	"testing"

	"github.com/tsubasa597/BILIBILI-HELPER/api/user"
)

func TestGetUserInfo(t *testing.T) {
	for _, up := range _ups {
		if _, err := user.GetUserInfo(up); err != nil {
			t.Error(err)
		}

	}
}

func TestUserCheck(t *testing.T) {
	if !_do {
		t.SkipNow()
	}

	if err := user.UserCheck(_api); err != nil {
		t.Error(err)
	}

}

func TestGetUserName(t *testing.T) {
	for _, up := range _ups {
		if _, err := user.GetUserName(up); err != nil {
			t.Error(err)
		}
	}
}
