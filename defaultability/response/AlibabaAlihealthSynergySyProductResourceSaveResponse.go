package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthSynergySyProductResourceSaveResponse struct {

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
	Result domain.AlibabaAlihealthSynergySyProductResourceSaveResultModel `json:"result,omitempty" `
}
