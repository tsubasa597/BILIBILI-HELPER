package utils

import (
	"bili/login"
	"bytes"
	"io/ioutil"
	"net/http"
	urlpkg "net/url"
)

// HTTP 请求的结构体
type HTTP struct {
	Verify  *login.UserInfo
	client  *http.Client
	request *http.Request
}

var (
	Http *HTTP = &HTTP{}
)

func init() {
	var userAgent string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 " +
		"(KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.70"
	Http.client = &http.Client{}
	Http.request = &http.Request{Header: make(http.Header), Method: "GET"}
	Http.Verify = &login.UserInfo{UserID: "35656398", SessData: "c0763d4c,1630379068,8c9fa*31", BiliJct: "e5ca623ee5d94759cf6d7a7b62cbf6c9"}
	Http.request.Header.Add("Connection", "keep-alive")
	Http.request.Header.Add("User-Agent", userAgent)
	Http.request.Header.Add("Cookie", Http.Verify.GetVerify())
	Http.request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
}

// Get 请求封装
func Get(url string) (resp []byte, err error) {
	Http.request.URL, err = urlpkg.Parse(url)
	Http.request.Method = "GET"
	if err != nil {
		return []byte(""), err
	}
	response, err := Http.client.Do(Http.request)
	if err != nil {
		return []byte(""), err
	}
	res, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return res, nil
}

// Post 请求封装
func Post(url string, postBody []byte) (resp []byte, err error) {
	Http.request.URL, err = urlpkg.Parse(url)
	Http.request.Method = "POST"
	Http.request.Body = ioutil.NopCloser(bytes.NewReader(postBody))
	if err != nil {
		return []byte(""), err
	}
	response, err := Http.client.Do(Http.request)
	if err != nil {
		return []byte(""), err
	}
	res, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return res, nil
}
