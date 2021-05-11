package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	URL "net/url"
	"sync"
)

type requests struct {
	cli  *http.Client
	req  *http.Request
	pool *sync.Pool
}

func newRequests() requests {
	return requests{
		cli: &http.Client{},
		req: &http.Request{},
		pool: &sync.Pool{
			New: func() interface{} {
				return http.Header{
					"Connection":   []string{"keep-alive"},
					"User-Agent":   []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.70"},
					"Content-Type": []string{"application/x-www-form-urlencoded"},
				}
			},
		},
	}
}

func (r *requests) setHeader(h http.Header) {
	r.req.Header = h
}

func (r requests) Post(url string, params URL.Values) ([]byte, error) {
	if r.req.Header == nil {
		r.req.Header = r.pool.Get().(http.Header)
	}

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

func (r requests) Get(url string) ([]byte, error) {
	if r.req.Header == nil {
		r.req.Header = r.pool.Get().(http.Header)
	}
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

func (r requests) gets(url string, v interface{}) error {
	rep, err := r.Get(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rep, v)
	return err
}

func (r requests) posts(url string, params URL.Values, v interface{}) error {
	rep, err := r.Post(url, params)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rep, v)
	return err
}
