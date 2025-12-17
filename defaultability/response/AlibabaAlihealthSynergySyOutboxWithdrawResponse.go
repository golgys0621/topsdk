package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthSynergySyOutboxWithdrawResponse struct {

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
	Result domain.AlibabaAlihealthSynergySyOutboxWithdrawResultModel `json:"result,omitempty" `
}
