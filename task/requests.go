package task

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	urlpkg "net/url"
)

// HTTP 请求的结构体
type Requests struct {
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

func (r Requests) httpRequest(method string, url string, postBody []byte) (rep *Response, err error) {
	r.request.Method = method
	r.request.URL, err = urlpkg.Parse(url)
	if postBody != nil {
		r.request.Body = ioutil.NopCloser(bytes.NewReader(postBody))
	}
	response, err := r.client.Do(r.request)
	if err != nil {
		return
	}
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	defer response.Body.Close()
	err = json.Unmarshal(res, &rep)
	return
}

func (r Requests) Get(url string) (resp *Response, err error) {
	return r.httpRequest("GET", url, nil)

}

func (r Requests) Post(url string, postBody []byte) (resp *Response, err error) {
	return r.httpRequest("POST", url, postBody)
}

func newRequests(conf config) (r Requests) {
	return Requests{
		request: &http.Request{
			Header: http.Header{
				"Connection":   []string{"keep-alive"},
				"User-Agent":   []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.70"},
				"Cookie":       []string{conf.Cookie.GetVerify()},
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			Method: "",
		},
		client: &http.Client{},
	}
}
