package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthDrugCodeKytWesCheckcoderelationResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回的码的结果
	*/
	Model []domain.AlibabaAlihealthDrugCodeKytWesCheckcoderelationWesCodeRelationDTO `json:"model,omitempty" `
	/*
	   返回信息
	*/
	MsgInfo string `json:"msg_info,omitempty" `
	/*
	   返回码
	*/
	MsgCode string `json:"msg_code,omitempty" `
	/*
	   调用是否成功
	*/
	ResponseSuccess bool `json:"response_success,omitempty" `
}
