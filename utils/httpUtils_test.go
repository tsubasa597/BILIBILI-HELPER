package utils_test

import (
	"bili/utils"
	"testing"
)

func TestGet(t *testing.T) {
	t.Log(utils.Get("https://api.live.bilibili.com/xlive/web-ucenter/v1/sign/DoSign"))
}

func TestPost(t *testing.T) {
	utils.Post("https://api.bilibili.com/x/click-interface/web/heartbeat", []byte(""))
}
