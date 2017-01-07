package utils

import (
	"github.com/YanshuoH/youkonger/consts"
)

type JSONResponse struct {
	ResultCode        string      `json:"resultCode"`
	ResultDescription string      `json:"resultDescription"`
	Data              interface{} `json:"data"`
}

func NewJSONResponse(resultCode string, datas ...interface{}) JSONResponse {
	r := JSONResponse{
		ResultCode:        resultCode,
		ResultDescription: consts.Messenger.Get(resultCode),
	}
	if len(datas) > 0 {
		r.Data = datas[0]
	}

	return r
}

func NewOKJSONResponse(data interface{}) JSONResponse {
	return JSONResponse{
		ResultCode: consts.OK,
		ResultDescription: consts.Messenger.Get(consts.OK),
		Data: data,
	}
}