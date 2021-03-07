package utils

import (
	"bili/config"
	"bytes"
	"io/ioutil"
	"net/http"
	urlpkg "net/url"
)

// HTTP 请求的结构体
type HTTP struct {
	client  *http.Client
	request *http.Request
}

var (
	// Http 结构体初始化
	Http HTTP = HTTP{}
)

func init() {
	Http.client = &http.Client{}
	Http.request = &http.Request{Header: make(http.Header), Method: "GET"}
	Http.request.Header.Add("Connection", "keep-alive")
	Http.request.Header.Add("User-Agent", config.Conf.UserAgent)
	Http.request.Header.Add("Cookie", config.Conf.Cookie.GetVerify())
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
