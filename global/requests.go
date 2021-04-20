package global

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	URL "net/url"
)

type Requests struct {
	cli *http.Client
	req *http.Request
}

func New() *Requests {
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

	if u, err := URL.Parse(url); err != nil {
		return nil, err
	} else {
		r.req.URL = u
	}

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

	if u, err := URL.Parse(url); err != nil {
		return nil, err
	} else {
		r.req.URL = u
	}

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
