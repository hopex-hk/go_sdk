package model

type GetConditionOrdersParam struct {
	ContractCodeList []string `json:"contractCodeList"`
	TaskTypeList     []int    `json:"taskTypeList"`
	TrigTypeList     []int    `json:"trigTypeList"`
	TaskStatusList   []int    `json:"taskStatusList"`
	Direct           int      `json:"direct"`
	Side             int      `json:"side"`
	StartTime        float64  `json:"startTime"`
	EndTime          float64  `json:"endTime"`
}
