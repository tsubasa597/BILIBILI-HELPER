package api_test

import (
	"testing"

	"github.com/tsubasa597/BILIBILI-HELPER/api/user"
)

func TestGetUserInfo(t *testing.T) {
	for _, up := range _ups {
		t.Log(user.GetUserInfo(up))
	}
}

func TestUserCheck(t *testing.T) {
	if !_do {
		t.SkipNow()
	}

	t.Log(user.UserCheck(_api))
}

func TestGetUserName(t *testing.T) {
	for _, up := range _ups {
		t.Log(user.GetUserName(up))
	}
}
