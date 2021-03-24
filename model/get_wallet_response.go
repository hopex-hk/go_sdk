package model

import (
	"github.com/hopex-hk/go_sdk/core/model"
	"github.com/shopspring/decimal"
)

type GetWalletResponse struct {
	*model.ApiResponseModel
	Data *WalletInfo `json:"data"`
}

type WalletInfo struct {
	Summary *AllAssetSummary `json:"summary"`
	Detail  []*AssetSummary  `json:"detail"`
}

type AllAssetSummary struct {
	ConversionCurrency string          `json:"conversionCurrency"` //计价货币
	TotalWealth        decimal.Decimal `json:"totalWealth,string"` //总权益
	FloatProfit        string          `json:"floatProfit"`        //浮动盈亏
	AvailableBalance   decimal.Decimal `json:"availableBalance"`   //总可用余额
}

type AssetSummary struct {
	AssetName             string          `json:"assetName"`
	FloatProfit           decimal.Decimal `json:"floatProfit,string"`
	FloatProfitLegal      string          `json:"floatProfitLegal"`
	ProfitRate            string          `json:"profitRate"`
	TotalWealth           decimal.Decimal `json:"totalWealth,string"`
	TotalWealthLegal      string          `json:"totalWealthLegal"`
	TotalWealthInfo       string          `json:"totalWealthInfo"`
	AvailableBalance      decimal.Decimal `json:"availableBalance,string"`
	AvailableBalanceLegal string          `json:"availableBalanceLegal"`
	WalletBalance         decimal.Decimal `json:"walletBalance,string"`
	WalletBalanceLegal    string          `json:"walletBalanceLegal"`
	PositionMargin        decimal.Decimal `json:"positionMargin,string"`
	PositionMarginLegal   string          `json:"positionMarginLegal"`
	DelegateMargin        decimal.Decimal `json:"delegateMargin,string"` //order margin
	DelegateMarginLegal   string          `json:"delegateMarginLegal"`
	WithdrawFreeze        decimal.Decimal `json:"withdrawFreeze,string"`
	WithdrawFreezeLegal   string          `json:"withdrawFreezeLegal"`
	DepositAmount         decimal.Decimal `json:"depositAmount,string"`
	DepositAmountLegal    string          `json:"depositAmountLegal"`
	WithdrawAmount        decimal.Decimal `json:"withdrawAmount,string"`
	WithdrawAmountLegal   string          `json:"withdrawAmountLegal"`
}

func NewGetWalletResponse() *GetWalletResponse {
	return &GetWalletResponse{
		ApiResponseModel: &model.ApiResponseModel{
			Ret: -1,
		},
		Data: &WalletInfo{},
	}
}
