package client

import (
	"encoding/json"
	"errors"
	"strconv"

	"hopex-hk/go_sdk/core"
	"hopex-hk/go_sdk/core/logging/zaplogger"
	coremodel "hopex-hk/go_sdk/core/model"
	"hopex-hk/go_sdk/model"
)

type MarketClient struct {
	requester core.HttpRequester
}

func (c *MarketClient) InitByDefault(cfg *core.Config) *MarketClient {
	c.requester = new(core.DefaultHttpRequester).Init(cfg, new(zaplogger.ZapLogger))

	return c
}

func (c *MarketClient) Init(requester core.HttpRequester) *MarketClient {
	c.requester = requester

	return c
}

//get kline data
//doc: https://hopex-hk.github.io/docs/contract/v1/cn/#hopex-k
func (c *MarketClient) GetKline(contractCode, interval string, startTime, endTime int64) (*model.KLine, error) {
	if len(contractCode) == 0 {
		return nil, errors.New("contractCode must not empty")
	}

	queries := make(map[string]string)

	queries["contractCode"] = contractCode
	queries["interval"] = interval
	queries["startTime"] = strconv.FormatInt(startTime, 10)
	queries["endTime"] = strconv.FormatInt(endTime, 10)

	res := model.NewGetKLineResponse()
	err := c.requester.Get("/api/v1/kline", queries, true, res)

	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

//get contract market ticker
//doc: https://hopex-hk.github.io/docs/contract/v1/cn/#hopex
func (c *MarketClient) GetMarketTicker(contractCode string) (*model.MarketTicker, error) {
	if len(contractCode) == 0 {
		return nil, errors.New("contractCode must not empty")
	}

	queries := map[string]string{
		"contractCode": contractCode,
	}

	res := model.NewGetMarketTickerResponse()
	err := c.requester.Get("/api/v1/ticker", queries, true, res)

	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

//get all contract market ticker
//doc: https://hopex-hk.github.io/docs/contract/v1/cn/#hopex-4
func (c *MarketClient) GetMarkets() ([]*model.MarketTicker1, error) {
	res := model.NewGetAllMarketTickerResponse()
	err := c.requester.Get("/api/v1/markets", nil, true, res)

	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

//get all contract market ticker
//doc: https://hopex-hk.github.io/docs/contract/v1/cn/#hopex-3
func (c *MarketClient) GetTrades(contractCode string, limit int) ([]*model.Trade, error) {

	queries := map[string]string{
		"contractCode": contractCode,
		"pageSize":     strconv.Itoa(limit),
	}

	res := model.NewGetTradesResponse()
	err := c.requester.Get("/api/v1/trades", queries, true, res)

	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

//get market depth
//doc: https://hopex-hk.github.io/docs/contract/v1/cn/#hopex-2
func (c *MarketClient) GetMarketDepth(contractCode string, limit int) (*model.MarketDepth, error) {
	param := &coremodel.ApiRequestModel{
		Param: &struct {
			ContractCode string `json:"contractCode"`
			PageSize     int    `json:"pageSize"`
		}{contractCode, limit},
	}

	body, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	res := model.NewGetMarketDepthResponse()
	err = c.requester.Post("/api/v1/depth", body, nil, true, res)

	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
