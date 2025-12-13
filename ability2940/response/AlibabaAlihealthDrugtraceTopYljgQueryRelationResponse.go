package response

import (
	"github.com/golgys0621/topsdk/ability2940/domain"
)

type AlibabaAlihealthDrugtraceTopYljgQueryRelationResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   接口返回model
	*/
	Result domain.AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel `json:"result,omitempty" `
}
