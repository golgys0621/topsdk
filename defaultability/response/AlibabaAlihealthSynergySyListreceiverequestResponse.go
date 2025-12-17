package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthSynergySyListreceiverequestResponse struct {

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
	ResResult domain.AlibabaAlihealthSynergySyListreceiverequestResultModel `json:"res_result,omitempty" `
}
