package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func SHA256(data []byte) string {
	hash := sha256.New()
	hash.Write(data)

	result := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return result
}

func HmacSHA256(secret, data []byte) string {
	hash := hmac.New(sha256.New, secret)
	hash.Write(data)

	result := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return result
}
