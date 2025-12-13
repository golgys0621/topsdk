package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthDrugCodeKytWesQuerycodeResponse struct {

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
	Result domain.AlibabaAlihealthDrugCodeKytWesQuerycodeResultModel `json:"result,omitempty" `
}
