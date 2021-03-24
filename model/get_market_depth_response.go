package model

import (
	"github.com/shopspring/decimal"
	"hopex-hk/go_sdk/core/model"
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
	HasOpenOrder int             `json:"Exist"`
}

func NewGetMarketDepthResponse() *GetMarketDepthResponse {
	return &GetMarketDepthResponse{
		ApiResponseModel: &model.ApiResponseModel{
			Ret: -1,
		},
		Data: &MarketDepth{},
	}
}
