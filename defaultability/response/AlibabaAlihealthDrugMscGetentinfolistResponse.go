package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthDrugMscGetentinfolistResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   resultModel
	*/
	Result domain.AlibabaAlihealthDrugMscGetentinfolistResultModel `json:"result,omitempty" `
}
