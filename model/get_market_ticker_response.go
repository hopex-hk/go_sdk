package model

import (
	"github.com/hopex-hk/go_sdk/core/model"
	"github.com/shopspring/decimal"
)

type GetMarketTickerResponse struct {
	*model.ApiResponseModel
	Data *MarketTicker `json:"data"`
}

type GetAllMarketTickerResponse struct {
	*model.ApiResponseModel
	Data []*MarketTicker1 `json:"data"`
}

type MarketTicker struct {
	ContractCode    string `json:"contractCode"`    //合约编码
	SpotIndexCode   string `json:"spotIndexCode"`   //现货指数编码
	FairPriceCode   string `json:"fairPriceCode"`   //合理价格编码
	ContractName    string `json:"contractName"`    //合约名称
	CloseCurrency   string `json:"closeCurrency"`   //结算货币
	AllowTrade      bool   `json:"allowTrade"`      //合约是否允许交易
	Pause           bool   `json:"pause"`           //是否暂停交易，暂停：true 启用: false
	LastPrice       string `json:"lastPrice"`       //最新价
	LastPriceLegal  string `json:"lastPriceLegal"`  //最新价 to 法币
	ChangePercent24 string `json:"changePercent24"` //24小时涨跌幅
	MarketPrice     string `json:"marketPrice"`     //现货指数价格
	MarketPriceInfo string `json:"marketPriceInfo"` //现货指数价格-解释
	FairPrice       string `json:"fairPrice"`       //合理价格
	FairePriceInfo  string `json:"fairePriceInfo"`  //合理价格-解释
	Price24Max      string `json:"price24Max"`      //24h最高
	Price24Min      string `json:"price24Min"`      //24h最低
	Amount24h       string `json:"amount24h"`       //24h交易额
	LastPriceToUSD  string `json:"lastPriceToUSD"`  //最新价To USD
	LastPriceToCNY  string `json:"lastPriceToCNY"`  // 最新价To CNY
	Quantity24h     string `json:"quantity24h"`     //24h交易量
	FundRate        string `json:"fundRate"`        //资金费率
}

type MarketTicker1 struct {
	ContractCode     string              `json:"contractCode"`     //合约编码
	ContractName     string              `json:"contractName"`     //合约名称
	AllowTrade       bool                `json:"allowTrade"`       //合约是否允许交易
	HasPosition      bool                `json:"hasPosition "`     //是否有持仓
	CloseCurrency    string              `json:"closeCurrency"`    //结算货币
	QuotedCurrency   string              `json:"quotedCurrency"`   //标价币种
	Precision        int                 `json:"precision"`        //合理价格精度
	MinPriceMovement decimal.Decimal     `json:"minPriceMovement"` //合约最小变动价格
	PricePrecision   int                 `json:"pricePrecision"`   //价格精度
	LastPrice        decimal.NullDecimal `json:"lastestPrice"`     //最新价
	ChangePercent24h decimal.NullDecimal `json:"changePercent24h"` //24小时涨跌幅
	SumAmount24h     decimal.NullDecimal `json:"sumAmount24h"`     //24小时成交额
	SumAmount24hUSDT decimal.NullDecimal `json:"sumAmount24hUSDT"` //24小时成交额-usdt
	PosVauleUSD      string              `json:"posVauleUSD"`      // 合约未平仓量价值(USD)
}

func NewGetMarketTickerResponse() *GetMarketTickerResponse {
	return &GetMarketTickerResponse{
		ApiResponseModel: &model.ApiResponseModel{
			Ret: -1,
		},
		Data: &MarketTicker{},
	}
}

func NewGetAllMarketTickerResponse() *GetAllMarketTickerResponse {
	return &GetAllMarketTickerResponse{
		ApiResponseModel: &model.ApiResponseModel{
			Ret: -1,
		},
		Data: make([]*MarketTicker1, 0),
	}
}
