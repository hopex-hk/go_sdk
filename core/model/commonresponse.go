package model

import "github.com/shopspring/decimal"

type BoolResponse struct {
	*ApiResponseModel
	Data bool `json:"data"`
}

type Int64Response struct {
	*ApiResponseModel
	Data int64 `json:"data"`
}

type DecimalResponse struct {
	*ApiResponseModel
	Data decimal.Decimal `json:"data"`
}

func NewBoolResponse() *BoolResponse {
	return &BoolResponse{
		ApiResponseModel: &ApiResponseModel{
			Ret: -1,
		},
		Data: false,
	}
}

func NewInt64Response() *Int64Response {
	return &Int64Response{
		ApiResponseModel: &ApiResponseModel{
			Ret: -1,
		},
	}
}

func NewDecimalResponse() *DecimalResponse {
	return &DecimalResponse{
		ApiResponseModel: &ApiResponseModel{
			Ret: -1,
		},
	}
}
