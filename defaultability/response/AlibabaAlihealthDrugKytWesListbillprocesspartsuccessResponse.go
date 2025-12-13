package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthDrugKytWesListbillprocesspartsuccessResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   是否响应成功
	*/
	ResponseSuccess bool `json:"response_success,omitempty" `
	/*
	   model
	*/
	Model domain.AlibabaAlihealthDrugKytWesListbillprocesspartsuccessPageInfoDTO `json:"model,omitempty" `
	/*
	   msgInfo
	*/
	MsgInfo string `json:"msg_info,omitempty" `
	/*
	   msgCode
	*/
	MsgCode string `json:"msg_code,omitempty" `
	/*
	   httpStatusCode
	*/
	HttpStatusCode int64 `json:"http_status_code,omitempty" `
}
