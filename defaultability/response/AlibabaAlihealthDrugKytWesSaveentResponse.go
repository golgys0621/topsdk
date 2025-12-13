package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthDrugKytWesSaveentResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   往来单位新增接口返回
	*/
	Result domain.AlibabaAlihealthDrugKytWesSaveentResultModel `json:"result,omitempty" `
}
