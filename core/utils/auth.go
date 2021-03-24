package utils

import (
	"net/http"
	"strings"
	"time"
)

var algorithm = "hmac-sha256"
var head_auth_headers = "date request-line digest"

func BuildHttpAuthHeader(apiKey, apiSecret, httpMethod, path string, data []byte) map[string]string {
	headers := make(map[string]string)

	date := time.Now().UTC().Format(http.TimeFormat)
	headers["Date"] = date

	if data == nil {
		data = []byte("{}")
	}

	digest := "SHA-256=" + SHA256(data)
	headers["Digest"] = digest

	textToSign := strings.Join(
		[]string{
			"date: " + date,
			httpMethod + " " + path + " HTTP/1.1",
			"digest: " + digest,
		},
		"\n",
	)
	signature := HmacSHA256([]byte(apiSecret), []byte(textToSign))

	head_auth := "hmac apikey=\"" + apiKey + "\", algorithm=\"" + algorithm + "\", headers=\"" + head_auth_headers + "\", signature=\"" + signature + "\""

	headers["Authorization"] = head_auth

	return headers
}
