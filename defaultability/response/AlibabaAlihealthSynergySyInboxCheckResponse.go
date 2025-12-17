package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthSynergySyInboxCheckResponse struct {

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
	Result domain.AlibabaAlihealthSynergySyInboxCheckResultModel `json:"result,omitempty" `
}
