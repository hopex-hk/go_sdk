package model

import (
	"github.com/hopex-hk/go_sdk/core/model"
	"github.com/shopspring/decimal"
)

type GetConditionOrdersResponse struct {
	*model.ApiResponseModel
	Data *ConditionOrderList `json:"data"`
}

type ConditionOrderList struct {
	*model.ListResultModel
	Result []*ConditionOrder `json:"result"`
}

type ConditionOrder struct {
	TaskId           int64           `json:"taskId"`
	TaskType         int             `json:"taskType"`
	TaskTypeD        string          `json:"taskTypeD"`
	ContractCode     string          `json:"contractCode"`
	ContractName     string          `json:"contractName"`
	Action           int             `json:"action"`     //1 open，2 close
	Direct           int             `json:"direct"`     //1 long，2 short
	Side             int             `json:"side"`       //1 sell, 2 buy
	TaskStatus       int             `json:"taskStatus"` //1 created, 2 cancelled, 3 trigged success, 4 trigged failure
	TaskStatusD      string          `json:"taskStatusD"`
	TrigType         int             `json:"trigType"`
	TrigTypeD        string          `json:"trigTypeD"`
	TrigPrice        decimal.Decimal `json:"trigPrice,string"`
	ExpectedQuantity string          `json:"expectedQuantity"`
	ExpectedPrice    string          `json:"expectedPrice"`
	CreateTime       string          `json:"createTime"`
	OrderId          int             `json:"orderId"`
	OrderQuantity    string          `json:"orderQuantity"`
	OrderPrice       string          `json:"orderPrice"`
	FinishTime       string          `json:"finishTime"`
	FailureReason    string          `json:"failureReason"`
	Leverage         string          `json:"leverage"`
}

func NewGetConditionOrdersResponse() *GetConditionOrdersResponse {
	return &GetConditionOrdersResponse{
		ApiResponseModel: &model.ApiResponseModel{
			Ret: -1,
		},
	}
}
