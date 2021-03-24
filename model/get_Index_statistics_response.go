package model

import "github.com/hopex-hk/go_sdk/core/model"

type GetIndexStatisticsResponse struct {
	*model.ApiResponseModel
	Data *IndexStatistics `json:"data"`
}

type IndexStatistics struct {
	PosVauleUSD   string `json:"posVauleUSD"`   //未平仓合约价值(USD)
	PosVauleCNY   string `json:"posVauleCNY"`   //未平仓合约价值(CNY)
	Amount24hUSD  string `json:"amount24hUSD"`  //24h交易额-usd
	Amount24hCNY  string `json:"amount24hCNY"`  //24h交易额-cny
	Amount7dayUSD string `json:"amount7dayUSD"` //7day交易额-usd
	Amount7dayCNY string `json:"amount7dayCNY"` //7day交易额-cny
	UserCount     string `json:"userCount"`     //用户数
	DealCountUSD  string `json:"dealCountUSD"`  //总交易额-usd
	DealCountCNY  string `json:"dealCountCNY"`  //总交易额-cny
}

func NewGetIndexStatisticsResponse() *GetIndexStatisticsResponse {
	return &GetIndexStatisticsResponse{
		ApiResponseModel: &model.ApiResponseModel{
			Ret: -1,
		},
		Data: &IndexStatistics{},
	}
}
