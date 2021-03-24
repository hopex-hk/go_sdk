package model

import (
	"github.com/hopex-hk/go_sdk/core/model"
	"github.com/shopspring/decimal"
)

type GetPositionsResponse struct {
	*model.ApiResponseModel
	Data []*Position `json:"data"`
}

type Position struct {
	ContractCode              string          `json:"contractCode"`
	ContractName              string          `json:"contractName"`
	AllowFullClose            bool            `json:"allowFullClose"`
	Leverage                  string          `json:"leverage"`
	ContractValue             string          `json:"contractValue"`
	ContractDirection         string          `json:"contractDirection"` //1 forward, 2 reverse
	MaintMarginRate           string          `json:"maintMarginRate"`
	TakerFee                  string          `json:"takerFee"`
	PositionQuantity          string          `json:"positionQuantity"`
	PositionQuantityD         int             `json:"positionQuantityD"`
	Direct                    int             `json:"direct"` //1 long, 2 short
	EntryPrice                string          `json:"entryPrice"`
	EntryPriceD               decimal.Decimal `json:"entryPriceD"`
	PositionMargin            string          `json:"positionMargin"`
	PositionMarginD           decimal.Decimal `json:"positionMarginD"`
	LiquidationPrice          string          `json:"liquidationPrice"`
	MaintMargin               string          `json:"maintMargin"`
	UnrealisedPnl             string          `json:"unrealisedPnl"`
	UnrealisedPnlPcnt         string          `json:"unrealisedPnlPcnt"`
	FairPrice                 string          `json:"fairPrice"`
	FairPriceD                decimal.Decimal `json:"fairPriceD"`
	LastPrice                 string          `json:"lastPrice"`
	Sequence                  int             `json:"sequence"`
	Rank                      int             `json:"rank"`
	MinPriceMovement          decimal.Decimal `json:"minPriceMovement"`
	MinPriceMovementPrecision int             `json:"minPriceMovementPrecision"`
	PositionQuantityFreeze    string          `json:"positionQuantityFreeze"`
	CloseablePositionQuantity string          `json:"closeablePositionQuantity"`
	IsAddMargin               bool            `json:"isAddMargin"`
	Sort                      int             `json:"sort"`
	CloseCurrency             string          `json:"closeCurrency"`
}

func NewGetPositionsResponse() *GetPositionsResponse {
	return &GetPositionsResponse{
		ApiResponseModel: &model.ApiResponseModel{
			Ret: -1,
		},
		Data: make([]*Position, 0),
	}
}
