package core

import (
	"fmt"
)

type Config struct {
	BaseUrl   string //openapi baseurl,eg: https://api2.hopex.com
	Apikey    string //api key
	ApiSecret string //api Secret
	UserAgent string //<product>/<product-version> <comment>
}

func NewConfig(baseUrl, apiKey, apiSecret, userAgent string) *Config {
	if len(userAgent) == 0 {
		panic(fmt.Errorf("userAgent must not empty"))
	}
	return &Config{
		BaseUrl:   baseUrl,
		Apikey:    apiKey,
		ApiSecret: apiSecret,
		UserAgent: userAgent,
	}
}
