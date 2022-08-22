package strparser

import (
	urlpkg "net/url"
	"path"
	"strings"
)

func ParseUrl(u string) string {
	u = strings.TrimSpace(u)
	if u == "" {
		return ""
	}

	url, err := urlpkg.Parse(u)
	if err != nil {
		return ""
	}
	return path.Base(url.Path)
}
