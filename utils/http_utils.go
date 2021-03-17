package utils

import (
	"bili/config"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	conf config.Config = *config.Init()
	ht   *HTTP         = &HTTP{
		request: &http.Request{
			Header: http.Header{
				"Connection":   []string{"keep-alive"},
				"User-Agent":   []string{conf.UserAgent},
				"Cookie":       []string{conf.Cookie.GetVerify()},
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			Method: "",
		},
	}
)

func (r *Response) HttpRequest(method string, url string, postBody []byte) (*Response, error) {
	ht.request.Method = method
	if postBody != nil {
		ht.request.Body = ioutil.NopCloser(bytes.NewReader(postBody))
	}
	response, err := ht.client.Do(ht.request)
	if err != nil {
		return r, err
	}
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return r, err
	}
	defer response.Body.Close()
	err = json.Unmarshal(res, &r)
	return r, err
}

func Get(url string) (resp *Response, err error) {
	return resp.HttpRequest("GET", url, nil)

}

func Post(url string, postBody []byte) (resp *Response, err error) {
	return resp.HttpRequest("POST", url, postBody)
}
