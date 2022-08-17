package client

import (
	"errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	urlpkg "net/url"
	"time"
)

type Config struct {
	Url     string
	Token   string
	Cookies *cookiejar.Jar
}

type Client struct {
	url   *urlpkg.URL
	token string
	c     *http.Client
}

func NewClient(cfg Config) *Client {
	url, err := urlpkg.Parse(cfg.Url)
	if err != nil {
		url = &urlpkg.URL{}
	}

	return &Client{
		url:   url,
		token: cfg.Token,
		c: &http.Client{
			Jar:     cfg.Cookies,
			Timeout: 3 * time.Minute,
		},
	}
}

func NewCookie(url string, v urlpkg.Values) (*cookiejar.Jar, error) {
	u, err := urlpkg.Parse(url)
	if err != nil {
		return nil, err
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	cookies := make([]*http.Cookie, 0)
	for k := range v {
		c := &http.Cookie{
			Name:    k,
			Value:   v.Get(k),
			Path:    "/",
			Expires: time.Now().AddDate(10, 0, 0),
		}
		if err := c.Valid(); err != nil {
			continue
		}
		cookies = append(cookies, c)
	}
	jar.SetCookies(u, cookies)
	return jar, nil
}

func (c Client) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	url, err := c.url.Parse(path)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:103.0) Gecko/20100101 Firefox/103.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("TE", "trailers")
	req.Header.Set("DNT", "1")
	req.Header.Set("Referer", "https://leetcode.com")
	req.Header.Set("x-csrftoken", c.token)
	req.Header.Set("authorization", "")
	req.Header.Set("Origin", "https://leetcode.com")
	req.Header.Set("content-type", "application/json")

	return req, nil
}

func (c Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	return resp, nil
}

func (c Client) Get(path string) (*http.Response, error) {
	req, err := c.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c Client) Post(path string, body io.Reader) (*http.Response, error) {
	req, err := c.NewRequest(http.MethodPost, path, body)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}
