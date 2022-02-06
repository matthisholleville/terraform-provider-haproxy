package haproxy

import (
	"encoding/base64"
	"strings"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func replaceSlashInString(value string) string {
	return strings.Replace(value, "/", "%2F", -1)
}
