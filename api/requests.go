package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	URL "net/url"
	"strings"
)

type requests struct {
	header http.Header
	cli    *http.Client
	req    *http.Request
}

func newRequests() requests {
	return requests{
		cli:    &http.Client{},
		req:    &http.Request{},
		header: make(http.Header),
	}
}

func (r *requests) setHeader(h http.Header) {
	r.header = h
}

func (r requests) post(url string, params URL.Values) ([]byte, error) {
	var err error

	if len(r.header) == 0 {
		r.header = http.Header{
			"Connection":   []string{"keep-alive"},
			"User-Agent":   []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.70"},
			"Content-Type": []string{"application/x-www-form-urlencoded"},
		}
	}

	r.req, err = http.NewRequest("POST", url, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}

	r.req.Header = r.header
	rep, err := r.cli.Do(r.req)
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

func (r requests) get(url string) ([]byte, error) {
	var err error

	if len(r.header) == 0 {
		r.header = http.Header{
			"Connection":   []string{"keep-alive"},
			"User-Agent":   []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.70"},
			"Content-Type": []string{"application/x-www-form-urlencoded"},
		}
	}

	r.req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	r.req.Header = r.header
	rep, err := r.cli.Do(r.req)
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

func (r requests) gets(url string, v interface{}) error {
	rep, err := r.get(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rep, v)
	return err
}

func (r requests) posts(url string, params URL.Values, v interface{}) error {
	rep, err := r.post(url, params)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rep, v)
	return err
}
