package model

type ListResultModel struct {
	TotalCount int `json:"totalCount"`
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
}
