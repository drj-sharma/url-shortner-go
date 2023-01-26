package utils

import (
	net_url "net/url"
)

func IsValidUrl(url string) {
	_, err := net_url.ParseRequestURI(url)
	if err != nil {
		panic(err)
	}
}
