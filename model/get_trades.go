package model

import "github.com/hopex-hk/go_sdk/core/model"

type GetTradesResponse struct {
	*model.ApiResponseModel
	Data []*Trade `json:"data"`
}

type Trade struct {
	Id           int64   `json:"id"`           //deal id
	Timestamp    float64 `json:"timestamp"`    //timestamp(ms)
	FillPrice    string  `json:"fillPrice"`    //fill price
	FillQuantity string  `json:"fillQuantity"` //fill quantity
	Side         string  `json:"side"`         //side, 1 sell 2 buy
}

func NewGetTradesResponse() *GetTradesResponse {
	return &GetTradesResponse{
		ApiResponseModel: &model.ApiResponseModel{
			Ret: -1,
		},
		Data: make([]*Trade, 0),
	}
}
