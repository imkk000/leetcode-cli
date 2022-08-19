package client

import (
	"fmt"
	"io"
	"leetcode-tool/config"
	"net/http"
	"os"
)

var DefaultClient *Client

func init() {
	endpoint := "https://leetcode.com"
	configDir, err := config.GetConfigDir()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cookies := config.ReadCookies(configDir)
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

func Head(path string) (*http.Response, error) {
	req, err := DefaultClient.NewRequest(http.MethodHead, path, nil)
	if err != nil {
		return nil, err
	}
	return DefaultClient.Do(req)
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
