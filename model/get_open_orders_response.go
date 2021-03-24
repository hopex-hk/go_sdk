package model

import "github.com/hopex-hk/go_sdk/core/model"

type GetOpenOrdersResponse struct {
	*model.ApiResponseModel
	Data []*OpenOrder `json:"data"`
}

type OpenOrder struct {
	OrderId            int64  `json:"orderId"`
	OrderType          string `json:"orderType"`
	OrderTypeVal       int    `json:"orderTypeVal"`
	Direct             int    `json:"direct"` //1 long, 2 short
	ContractCode       string `json:"contractCode"`
	ContractName       string `json:"contractName"`
	Type               string `json:"type"` //1 limit open 2 market open 3 limit close 4 market close
	Side               string `json:"side"` //1 sell, 2 buy
	SideDisplay        string `json:"sideDisplay"`
	Ctime              string `json:"ctime"`
	Mtime              string `json:"mtime"`
	OrderQuantity      string `json:"orderQuantity"`
	LeftQuantity       string `json:"leftQuantity"`
	FillQuantity       string `json:"fillQuantity"`
	OrderStatus        string `json:"orderStatus"` //1 partial deal, 2 will deal
	OrderStatusDisplay string `json:"orderStatusDisplay"`
	OrderPrice         string `json:"orderPrice"`
	Leverage           string `json:"leverage"`
	Fee                string `json:"fee"`
	AvgFillMoney       string `json:"avgFillMoney"`
	OrderMargin        string `json:"orderMargin"`
	ExpireTime         string `json:"expireTime"`
}

func NewGetOpenOrdersResponse() *GetOpenOrdersResponse {
	return &GetOpenOrdersResponse{
		ApiResponseModel: &model.ApiResponseModel{
			Ret: -1,
		},
		Data: make([]*OpenOrder, 0),
	}
}
