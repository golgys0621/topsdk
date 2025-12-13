package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthSynergyYzwQuerydrugreportResponse struct {

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
	ResResult domain.AlibabaAlihealthSynergyYzwQuerydrugreportResultModel `json:"res_result,omitempty" `
}
