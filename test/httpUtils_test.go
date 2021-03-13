package bilitest

import (
	"bili/utils"
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	res, _ := utils.Get("http://passport.bilibili.com/login?act=exit")
	fmt.Println(res)
}

func TestPost(t *testing.T) {
	// HTTP.Post("https://api.bilibili.com/x/click-interface/web/heartbeat", []byte(""))
}
