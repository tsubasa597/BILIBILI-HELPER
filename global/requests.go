package global

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	URL "net/url"
	"sync"
)

type Requests struct {
	cli *http.Client
	req *http.Request
}

func NewRequests() *Requests {
	return &Requests{
		cli: &http.Client{},
		req: &http.Request{},
	}
}

func (r *Requests) SetHeader(h http.Header) {
	r.req.Header = h
}

func (r Requests) Post(url string, params URL.Values) ([]byte, error) {
	r.req.Method = "POST"

	u, err := URL.Parse(url)
	if err != nil {
		return nil, err
	}
	r.req.URL = u

	r.req.PostForm = params

	rep, err := r.cli.PostForm(url, params)
	if err != nil {
		return nil, err
	}
	defer rep.Body.Close()

	data, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r Requests) Get(url string) ([]byte, error) {
	r.req.Method = "GET"

	u, err := URL.Parse(url)
	if err != nil {
		return nil, err
	}
	r.req.URL = u

	rep, err := r.cli.Do(r.req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Json(resp []byte, v interface{}) error {
	return json.Unmarshal(resp, v)
}

func Get(url string) ([]byte, error) {
	req := &http.Request{}
	cli := http.Client{}
	req.Method = "GET"

	req.Header = pool.Get().(http.Header)

	u, err := URL.Parse(url)
	if err != nil {
		return nil, err
	}
	req.URL = u

	rep, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Post(url string, params URL.Values) ([]byte, error) {
	req := &http.Request{}
	cli := http.Client{}
	req.Method = "POST"

	req.Header = pool.Get().(http.Header)

	u, err := URL.Parse(url)
	if err != nil {
		return nil, err
	}
	req.URL = u

	req.PostForm = params

	rep, err := cli.PostForm(url, params)
	if err != nil {
		return nil, err
	}
	defer rep.Body.Close()

	data, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// 增加 http.Header 对象重用率
var pool *sync.Pool = &sync.Pool{
	New: func() interface{} {
		return http.Header{
			"Connection":   []string{"keep-alive"},
			"User-Agent":   []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.70"},
			"Content-Type": []string{"application/x-www-form-urlencoded"},
		}
	},
}
