package response

import (
	"github.com/golgys0621/topsdk/ability2940/domain"
)

type AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   接口返回
	*/
	Result domain.AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel `json:"result,omitempty" `
}
