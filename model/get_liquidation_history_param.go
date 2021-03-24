package model

type GetLiquidationHistoryParam struct {
	Side             *int     `json:"side,omitempty"`
	ContractCodeList []string `json:"contractCodeList"`
}
