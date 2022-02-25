package haproxy

import (
	"encoding/base64"
	"net/url"
	"regexp"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func encodeUrl(s string) string {
	return url.QueryEscape(s)
}

func ExtractStringWithRegex(value string, regex string) string {
	re := regexp.MustCompile(regex)
	res := re.FindAllStringSubmatch(value, 1)
	for i := range res {
		return res[i][1]
	}
	return value
}
