package main

import "strings"

func httpsStrip(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	}
	if strings.HasPrefix(url, "https://") {
		// http stuff
		url = strings.TrimPrefix(url, "https://")
	}
	return "http://" + url
}
