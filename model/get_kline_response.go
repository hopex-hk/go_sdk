package model

import "github.com/hopex-hk/go_sdk/core/model"

type GetKLineResponse struct {
	*model.ApiResponseModel
	Data *KLine `json:"data"`
}

type KLine struct {
	Decimalplace string      `json:"decimalplace"` //小数点位数
	TimeData     []KLineItem `json:"timeData"`     //数据
}

type KLineItem struct {
	Time          int64  `json:"time"`          //timestamp(s)
	Open          string `json:"open"`          //开市值
	Close         string `json:"close"`         //闭市值
	High          string `json:"high"`          //最高价
	Low           string `json:"low"`           //最低价
	Vol           string `json:"vol"`           //成交量
	Val           string `json:"val"`           //成交额
	PrevClose     string `json:"prevClose"`     //上一笔的闭市价
	UpDown        string `json:"upDown"`        //涨跌额
	UpDownRate    string `json:"upDownRate"`    //涨跌率
	Direct        int    `json:"Direct"`        //合约方向, 1:Forward,-1 Reverse
	ContractValue string `json:"ContractValue"` //合约价值
}

func NewGetKLineResponse() *GetKLineResponse {
	return &GetKLineResponse{
		ApiResponseModel: &model.ApiResponseModel{
			Ret: -1,
		},
		Data: &KLine{},
	}
}
