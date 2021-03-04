package login_test

import (
	"bili/login"
	"testing"
)

func TestVerify(t *testing.T) {
	info := login.UserInfo{"1", "2", "3"}
	t.Log(info.GetVerify())
}
