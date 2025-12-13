package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthSynergyYzwSignednotstampedbillQueryResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   结果
	*/
	ResResult domain.AlibabaAlihealthSynergyYzwSignednotstampedbillQueryResultModel `json:"res_result,omitempty" `
}
