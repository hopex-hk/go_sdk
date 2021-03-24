package model

import "errors"

type ApiResponseModel struct {
	Ret     int    `json:"ret"`     //success:0, fail:not 0
	ErrCode string `json:"errCode"` //err code
	ErrStr  string `json:"errStr"`  //err msg
	Message string `json:"message"` //error message ("API rate limit exceeded || Hmac signature no match")
}

func (response *ApiResponseModel) CheckRet() error {
	if response.Ret != 0 {
		errMsg := response.ErrStr
		if len(errMsg) == 0 {
			errMsg = response.Message
		}
		return errors.New(errMsg)
	}
	return nil
}
