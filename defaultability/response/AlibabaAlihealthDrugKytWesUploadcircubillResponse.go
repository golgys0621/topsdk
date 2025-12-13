package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthDrugKytWesUploadcircubillResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回结果
	*/
	Model string `json:"model,omitempty" `
	/*
	   返回结果
	*/
	MsgCode string `json:"msg_code,omitempty" `
	/*
	   返回结果
	*/
	MsgInfo string `json:"msg_info,omitempty" `
	/*
	   是否成功(true 成功 ,false失败)
	*/
	ResponseSuccess bool `json:"response_success,omitempty" `
	/*
	   结果子编码
	*/
	SubMsgCode string `json:"sub_msg_code,omitempty" `
	/*
	   错误信息明细,便于后续操作
	*/
	OperationInfoMap domain.AlibabaAlihealthDrugKytWesUploadcircubillOperationInfoMap `json:"operation_info_map,omitempty" `
}
