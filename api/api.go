package api

import "bili/utils"

// JSON 返回值解析
type JSON utils.Response

// API 部分 api 简易封装
type API interface {
	// GiveGift 赠送礼物
	GiveGift()
}
