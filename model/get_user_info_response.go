package model

import "hopex-hk/go_sdk/core/model"

type GetUserInfoResponse struct {
	*model.ApiResponseModel
	Data *UserInfo `json:"data"`
}

type UserInfo struct {
	ConversionCurrency string `json:"conversionCurrency"` //计价货币
	ProfitRate         string `json:"profitRate"`         //当前持仓收益率(浮动盈亏/持仓占用保证金)
	TotalWealth        string `json:"totalWealth"`        //账户总权益估值（法币)
	FloatProfit        string `json:"floatProfit"`        //总浮动盈亏估值（法币)
	Positions          int    `json:"position"`           //持仓数量
	OpenOrders         int    `json:"activeOrder"`        //活跃委托数量
}

func NewGetUserInfoResponse() *GetUserInfoResponse {
	return &GetUserInfoResponse{
		ApiResponseModel: &model.ApiResponseModel{
			Ret: -1,
		},
		Data: &UserInfo{},
	}
}
