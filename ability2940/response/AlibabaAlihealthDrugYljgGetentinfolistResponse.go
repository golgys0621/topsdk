package response

import (
	"github.com/golgys0621/topsdk/ability2940/domain"
)

type AlibabaAlihealthDrugYljgGetentinfolistResponse struct {

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
	Result domain.AlibabaAlihealthDrugYljgGetentinfolistResultModel `json:"result,omitempty" `
}
