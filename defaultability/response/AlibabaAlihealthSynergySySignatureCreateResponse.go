package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthSynergySySignatureCreateResponse struct {

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
	Result domain.AlibabaAlihealthSynergySySignatureCreateResultModel `json:"result,omitempty" `
}
