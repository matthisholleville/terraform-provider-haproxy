package haproxy

import (
	"encoding/base64"
	"regexp"
	"strings"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func replaceSlashInString(value string) string {
	return strings.Replace(value, "/", "%2F", -1)
}

func ExtractStringWithRegex(value string, regex string) string {
	re := regexp.MustCompile(regex)
	res := re.FindAllStringSubmatch(value, 1)
	for i := range res {
		return res[i][1]
	}
	return value
}
