package model

import "hopex-hk/go_sdk/core/model"

type GetLiquidationHistoryResponse struct {
	*model.ApiResponseModel
	Data *LiquidationHistoryList `json:"data"`
}

type LiquidationHistoryList struct {
	*model.ListResultModel
	Result []*LiquidationHistory `json:"result"`
}

type LiquidationHistory struct {
	OrderId              int64  `json:"orderId"`
	ContractCode         string `json:"contractCode"`
	ContractName         string `json:"ContractName"`
	Side                 string `json:"side"`
	SideDisplay          string `json:"sideDisplay"`
	OrderType            string `json:"orderType"`
	OrderTypeVal         int    `json:"orderTypeVal"`
	Direct               int    `json:"direct"`
	Leverage             string `json:"leverage"`
	OrderQuantity        string `json:"orderQuantity"`
	OrderPrice           string `json:"orderPrice"`
	ClosePosPNL          string `json:"closePosPNL"`
	Fee                  string `json:"fee"`
	Ctime                string `json:"ctime"`
	Timestamp            int64  `json:"timestamp"`
	Direction            int    `json:"direction"`
	DirectionDisplay     string `json:"directionDisplay"`
	PositionMargin       string `json:"positionMargin"`
	OpenPrice            string `json:"openPrice"`
	LiquidationPriceReal string `json:"liquidationPriceReal"`
}

func NewGetLiquidationHistoryResponse() *GetLiquidationHistoryResponse {
	return &GetLiquidationHistoryResponse{
		ApiResponseModel: &model.ApiResponseModel{
			Ret: -1,
		},
		Data: &LiquidationHistoryList{},
	}
}
