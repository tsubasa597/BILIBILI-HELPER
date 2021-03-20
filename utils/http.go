package utils

import (
	"bili/conf"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	urlpkg "net/url"
)

// HTTP 请求的结构体
type HTTP struct {
	client  *http.Client
	request *http.Request
}

// Response 返回值解析
type Response struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	TTL     int                    `json:"ttl"`
	Data    map[string]interface{} `json:"data"`
}

var (
	config conf.Config = *conf.Init()
	ht     *HTTP       = &HTTP{
		request: &http.Request{
			Header: http.Header{
				"Connection":   []string{"keep-alive"},
				"User-Agent":   []string{config.UserAgent},
				"Cookie":       []string{config.Cookie.GetVerify()},
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			Method: "",
		},
		client: &http.Client{},
	}
)

func httpRequest(method string, url string, postBody []byte) (r *Response, err error) {
	ht.request.Method = method
	ht.request.URL, err = urlpkg.Parse(url)
	if postBody != nil {
		ht.request.Body = ioutil.NopCloser(bytes.NewReader(postBody))
	}
	response, err := ht.client.Do(ht.request)
	if err != nil {
		return
	}
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	defer response.Body.Close()
	err = json.Unmarshal(res, &r)
	return
}

func Get(url string) (resp *Response, err error) {
	return httpRequest("GET", url, nil)

}

func Post(url string, postBody []byte) (resp *Response, err error) {
	return httpRequest("POST", url, postBody)
}
