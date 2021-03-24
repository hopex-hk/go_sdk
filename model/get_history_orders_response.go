package model

import "hopex-hk/go_sdk/core/model"

type GetHistoryOrdersResponse struct {
	*model.ApiResponseModel
	Data *HistoryOrderList `json:"data"`
}

type HistoryOrderList struct {
	*model.ListResultModel
	Result []*HistoryOrder `json:"result"`
}

type HistoryOrder struct {
	OrderId            int64   `json:"orderId"`
	ContractCode       string  `json:"contractCode"`
	ContractName       string  `json:"contractName"`
	Type               string  `json:"type"` //1 limit open 2 market open 3 limit close 4 market close
	Side               string  `json:"side"` //1 sell, 2 buy
	SideDisplay        string  `json:"sideDisplay"`
	Direct             int     `json:"direct"` //1 long, 2 short
	Ctime              string  `json:"ctime"`
	Ftime              string  `json:"ftime"`
	OrderQuantity      string  `json:"orderQuantity"`
	FillQuantity       string  `json:"fillQuantity"`
	OrderStatus        string  `json:"orderStatus"`
	OrderStatusDisplay string  `json:"orderStatusDisplay"`
	OrderPrice         string  `json:"orderPrice"`
	Leverage           string  `json:"leverage"`
	Fee                string  `json:"fee"`
	AvgFillMoney       string  `json:"avgFillMoney"`
	ClosePosPNL        string  `json:"closePosPNL"`
	Timestamp          int64   `json:"timestamp"`
	OrderTypeVal       int     `json:"orderTypeVal"`
	OrderType          string  `json:"orderType"`
	CancelReason       *string `json:"cancelReason"`
}

func NewGetHistoryOrdersResponse() *GetHistoryOrdersResponse {
	return &GetHistoryOrdersResponse{
		ApiResponseModel: &model.ApiResponseModel{
			Ret: -1,
		},
		Data: &HistoryOrderList{},
	}
}
