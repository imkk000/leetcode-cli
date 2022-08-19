package config

import (
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

func ReadCookies(configDir string) url.Values {
	cookies := url.Values{}

	vp := viper.New()
	vp.AddConfigPath(configDir)
	vp.SetConfigType("json")
	vp.SetConfigName("cookies")
	if err := vp.ReadInConfig(); err != nil {
		return cookies
	}

	for k, v := range vp.GetStringMapString("Request Cookies") {
		if k == "new_problemlist_page" || k == "leetcode_session" {
			k = strings.ToUpper(k)
		}
		cookies.Set(k, v)
	}
	return cookies
}
