// Package api bilibili 部分接口封装
package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/requests"
)

// API 发起请求所需的数据
type API struct {
	cookie *cookie
	Req    *requests.Requests
}

// New 初始化
func New(path string, enrty *logrus.Entry) (*API, error) {
	c, err := newCookie(path)
	if err != nil {
		return nil, err
	}

	return &API{
		cookie: c,
		Req: &requests.Requests{
			Client: &http.Client{},
			Headers: map[string]string{
				"Connection":   "keep-alive",
				"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.70",
				"Content-Type": "application/x-www-form-urlencoded",
			},
			Cookies: map[string]string{
				"DedeUserID": c.UserID,
				"SESSDATA":   c.SessData,
				"bili_jct":   c.BiliJct,
			},
		},
	}, nil
}
