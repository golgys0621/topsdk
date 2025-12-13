package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthDrugMscServiceinfoResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回的结果对象
	*/
	Result domain.AlibabaAlihealthDrugMscServiceinfoResultModel `json:"result,omitempty" `
}
