package client

import (
	"io"
	"leetcode-tool/config"
	"net/http"
)

var DefaultClient *Client

func init() {
	endpoint := "https://leetcode.com"
	// endpoint := "http://localhost:9900"
	cookies := config.ReadCookies()
	jar, err := NewCookie(endpoint, cookies)
	if err != nil {
		panic(err)
	}
	DefaultClient = NewClient(Config{
		Url:     endpoint,
		Token:   cookies.Get("csrftoken"),
		Cookies: jar,
	})
}

func Get(path string) (*http.Response, error) {
	req, err := DefaultClient.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	return DefaultClient.Do(req)
}

func Post(path string, body io.Reader) (*http.Response, error) {
	req, err := DefaultClient.NewRequest(http.MethodPost, path, body)
	if err != nil {
		return nil, err
	}
	return DefaultClient.Do(req)
}
