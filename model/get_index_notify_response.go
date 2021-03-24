package model

import "github.com/hopex-hk/go_sdk/core/model"

type GetIndexNotifyResponse struct {
	*model.ApiResponseModel
	Data *IndexNotifyList `json:"data"`
}

type IndexNotifyList struct {
	*model.ListResultModel
	Result []IndexNotify `json:"result"`
}

type IndexNotify struct {
	Id        int    `json:"id"`        //id
	Title     string `json:"title"`     //title
	Link      string `json:"link"`      //link
	Timestamp int32  `json:"timestamp"` //timestamp(s)
}

func NewGetIndexNotifyResponse() *GetIndexNotifyResponse {
	return &GetIndexNotifyResponse{
		ApiResponseModel: &model.ApiResponseModel{
			Ret: -1,
		},
		Data: &IndexNotifyList{
			ListResultModel: &model.ListResultModel{
				TotalCount: 0,
				Page:       -1,
				PageSize:   0,
			},
			Result: make([]IndexNotify, 0),
		},
	}
}
