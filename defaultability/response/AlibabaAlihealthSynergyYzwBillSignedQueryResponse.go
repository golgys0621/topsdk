package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthSynergyYzwBillSignedQueryResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   结果对象
	*/
	ResResult domain.AlibabaAlihealthSynergyYzwBillSignedQueryResultModel `json:"res_result,omitempty" `
}
