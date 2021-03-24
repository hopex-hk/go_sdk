package model

import "github.com/hopex-hk/go_sdk/core/model"

type GetOrderParameterResponse struct {
	*model.ApiResponseModel
	Data map[string]interface{} `json:"data"`
}

func NewGetOrderParameterResponse() *GetOrderParameterResponse {
	return &GetOrderParameterResponse{
		ApiResponseModel: &model.ApiResponseModel{
			Ret: -1,
		},
	}
}
