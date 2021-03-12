package utils

import (
	"bili/config"
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

func (r *Response) HttpRequest(method string, url string, postBody []byte) (rep *Response, e error) {
	h := &HTTP{}
	var err error
	h, err = initHttp(h, method, url)
	if postBody != nil {
		h.request.Body = ioutil.NopCloser(bytes.NewReader(postBody))
	}
	response, err := h.client.Do(h.request)
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

func initHttp(h *HTTP, method string, url string) (ht *HTTP, e error) {
	h.client = &http.Client{}
	var err error
	h.request = &http.Request{Header: make(http.Header), Method: method}
	h.request.Header.Add("Connection", "keep-alive")
	h.request.Header.Add("User-Agent", config.Conf.UserAgent)
	h.request.Header.Add("Cookie", config.Conf.Cookie.GetVerify())
	h.request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	h.request.URL, err = urlpkg.Parse(url)
	if err != nil {
		return h, err
	}
	return h, err
}

func Get(url string) (resp *Response, err error) {
	return resp.HttpRequest("GET", url, nil)

}

func Post(url string, postBody []byte) (resp *Response, err error) {
	return resp.HttpRequest("POST", url, postBody)
}
