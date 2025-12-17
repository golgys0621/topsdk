package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthSynergySyCreaterequestResponse struct {

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
	ResResult domain.AlibabaAlihealthSynergySyCreaterequestResultModel `json:"res_result,omitempty" `
}
