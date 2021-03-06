package bilitest

import (
	"bili/utils"
	"testing"
)

var (
	HTTP utils.HTTP = utils.HTTP{}
)

func TestGet(t *testing.T) {
	// t.Log(HTTP.Get("https://api.live.bilibili.com/xlive/web-ucenter/v1/sign/DoSign"))
}

func TestPost(t *testing.T) {
	// HTTP.Post("https://api.bilibili.com/x/click-interface/web/heartbeat", []byte(""))
}
