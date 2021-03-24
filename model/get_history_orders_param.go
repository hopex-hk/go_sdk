package model

type GetHistoryOrdersParam struct {
	ContractCodeList []string `json:"contractCodeList"`
	TypeList         []int    `json:"typeList"`
	Side             int      `json:"side"`
	StartTime        int64    `json:"startTime"`
	EndTime          int64    `json:"EndTime"`
}
