package model

import "github.com/hopex-hk/go_sdk/core/model"

type GetAccountHistoryResponse struct {
	*model.ApiResponseModel
	Data *AccountHistoryList `json:"data"`
}

type AccountHistoryList struct {
	*model.ListResultModel
	Result []*AccountHistory `json:"result"`
}

type AccountHistory struct {
	Id           int    `json:"id"`
	Asset        string `json:"asset"`
	OrderType    int    `json:"OrderType"` //1 OTC入金，2 OTC出金，3 链上入金，4 链上出金，5 内部转账-入金，6 内部转账-出金, 7 人工入金, 8 人工出金,9 快速入金,10 快速出金
	OrderTypeD   string `json:"orderTypeD"`
	Amount       string `json:"amount"`
	RMBAmount    string `json:"rmbAmount"`
	BankName     string `json:"bankName"`
	Addr         string `json:"addr"`
	OrderStatus  int    `json:"orderStatus"`
	OrderStatusD string `json:"orderStatusD"`
	CreatedTime  string `json:"createdTime"`
}

func NewGetAccountHistoryResponse() *GetAccountHistoryResponse {
	return &GetAccountHistoryResponse{
		ApiResponseModel: &model.ApiResponseModel{
			Ret: -1,
		},
		Data: &AccountHistoryList{},
	}
}
