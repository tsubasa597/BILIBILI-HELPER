package utils

import (
	"bili/login"
	"bytes"
	"io/ioutil"
	"net/http"
	urlpkg "net/url"
)

var (
	userAgent string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 " +
		"(KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.70"
	verify   login.UserInfo = login.UserInfo{UserID: "35656398", SessData: "c0763d4c,1630379068,8c9fa*31", BiliJct: "e5ca623ee5d94759cf6d7a7b62cbf6c9"}
	client   http.Client    = http.Client{}
	request  *http.Request  = &http.Request{Header: make(http.Header), Method: "GET"}
	response *http.Response = &http.Response{}
)

func init() {
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", userAgent)
	request.Header.Add("Cookie", verify.GetVerify())
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
}

// Get 请求封装
func Get(url string) (resp []byte, err error) {
	request.URL, err = urlpkg.Parse(url)
	request.Method = "GET"
	if err != nil {
		return []byte(""), err
	}
	response, err := client.Do(request)
	if err != nil {
		return []byte(""), err
	}
	res, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return res, nil
}

// Post 请求封装
func Post(url string, postBody []byte) (resp []byte, err error) {
	request.URL, err = urlpkg.Parse(url)
	request.Method = "POST"
	request.Body = ioutil.NopCloser(bytes.NewReader(postBody))
	if err != nil {
		return []byte(""), err
	}
	response, err = client.Do(request)
	if err != nil {
		return []byte(""), err
	}
	res, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return res, nil
}
