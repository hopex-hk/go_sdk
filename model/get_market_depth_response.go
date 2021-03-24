package model

import (
	"github.com/hopex-hk/go_sdk/core/model"
	"github.com/shopspring/decimal"
)

type GetMarketDepthResponse struct {
	*model.ApiResponseModel
	Data *MarketDepth `json:"data"`
}

type MarketDepth struct {
	ContractCode string          `json:"contractCode"`
	Decimalplace string          `json:"decimalplace"`
	Intervals    []string        `json:"intervals"`
	AsksFilter   string          `json:"asksFilter"`
	Asks         []OrderBookItem `json:"asks"`
	BidsFilter   string          `json:"bidsFilter"`
	Bids         []OrderBookItem `json:"bids"`
}

type OrderBookItem struct {
	PriceD       decimal.Decimal `json:"priceD"`
	Price        string          `json:"orderPrice"`
	Quantity     int64           `json:"orderQuantity"`
	HasOpenOrder int             `json:"exist"`
}

func NewGetMarketDepthResponse() *GetMarketDepthResponse {
	return &GetMarketDepthResponse{
		ApiResponseModel: &model.ApiResponseModel{
			Ret: -1,
		},
		Data: &MarketDepth{},
	}
}
