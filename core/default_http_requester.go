package core

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	pathAlias "path"
	"strings"
	"time"

	"github.com/hopex-hk/go_sdk/core/logging"
	"github.com/hopex-hk/go_sdk/core/logging/zaplogger"
	"github.com/hopex-hk/go_sdk/core/utils"
)

type DefaultHttpRequester struct {
	baseUrl *url.URL
	config  *Config
	logger  logging.Logger
}

func (req *DefaultHttpRequester) Init(config *Config, logger logging.Logger) *DefaultHttpRequester {
	url, err := url.Parse(config.BaseUrl)
	if err != nil {
		panic(err)
	}

	req.baseUrl = url
	req.config = config
	req.logger = logger

	return req
}

func (req *DefaultHttpRequester) InitByZapLogger(config *Config) *DefaultHttpRequester {
	return req.Init(config, new(zaplogger.ZapLogger))
}

func (req *DefaultHttpRequester) newUrl(path string, queries map[string]string) *url.URL {
	url := &url.URL{
		Scheme: req.baseUrl.Scheme,
		Host:   req.baseUrl.Host,
		Path:   pathAlias.Join(req.baseUrl.Path, path),
	}
	urls := url.Query()

	if queries != nil {
		for key, value := range queries {
			urls.Add(key, value)
		}
	}
	urls.Add("culture", "en")

	url.RawQuery = urls.Encode()

	return url
}

func (req *DefaultHttpRequester) Get(path string, queries map[string]string, auth bool, response Response) error {
	start := time.Now()

	url := req.newUrl(path, queries)
	urlStr := url.String()

	if req.logger.Enable(logging.DEBUG) {
		req.logger.Debug("get %s", urlStr)
	}

	getReq, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return err
	}
	getReq.Header = map[string][]string{
		"User-Agent": {req.config.UserAgent},
	}

	if auth {
		authHeaders := utils.BuildHttpAuthHeader(req.config.Apikey, req.config.ApiSecret, "GET", url.Path, nil)
		for k, v := range authHeaders {
			getReq.Header.Add(k, v)
		}
	}

	if req.logger.Enable(logging.DEBUG) {
		req.logger.Debug("-------------begin http headers----------")

		for k, v := range getReq.Header {
			req.logger.Debug("%s=%s", k, strings.Join(v, ","))
		}

		req.logger.Debug("-------------end http headers-------------")
	}
	resp, err := http.DefaultClient.Do(getReq)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if req.logger.Enable(logging.DEBUG) {
		req.logger.Debug("raw res: %s", string(result))
	}
	req.logger.Info("get %s %d ms", urlStr, time.Since(start).Milliseconds())

	err = json.Unmarshal(result, response)
	if err != nil {
		return err
	}
	err = response.CheckRet()
	if err != nil {
		return err
	}

	return err
}

func (req *DefaultHttpRequester) Post(path string, body []byte, queries map[string]string, auth bool, response Response) error {
	start := time.Now()

	url := req.newUrl(path, queries)
	urlStr := url.String()

	if req.logger.Enable(logging.DEBUG) {
		req.logger.Debug("post %s, body: %s", urlStr, string(body))
	}

	getReq, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(body))
	getReq.Header = map[string][]string{}
	if err != nil {
		return err
	}
	if auth {
		authHeaders := utils.BuildHttpAuthHeader(req.config.Apikey, req.config.ApiSecret, "POST", url.Path, body)
		getReq.Header.Add("Content-Type", "application/json")
		for k, v := range authHeaders {
			getReq.Header.Add(k, v)
		}
	}
	if req.logger.Enable(logging.DEBUG) {
		req.logger.Debug("-------------begin http headers----------")

		for k, v := range getReq.Header {
			req.logger.Debug("%s=%s", k, strings.Join(v, ","))
		}

		req.logger.Debug("-------------end http headers-------------")
	}
	resp, err := http.DefaultClient.Do(getReq)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)

	if req.logger.Enable(logging.DEBUG) {
		req.logger.Debug("raw res: %s", string(result))
	}
	req.logger.Info("post %s %d ms", urlStr, time.Since(start).Milliseconds())

	err = json.Unmarshal(result, response)
	if err != nil {
		return err
	}
	err = response.CheckRet()
	if err != nil {
		return err
	}

	return nil
}
