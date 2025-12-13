package response

import (
	"github.com/golgys0621/topsdk/ability2940/domain"
)

type AlibabaAlihealthDrugtraceTopYljgQueryCodedetailResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   最外层结果
	*/
	Result domain.AlibabaAlihealthDrugtraceTopYljgQueryCodedetailResultModel `json:"result,omitempty" `
}
