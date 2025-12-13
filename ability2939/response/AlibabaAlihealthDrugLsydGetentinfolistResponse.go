package response

import (
	"github.com/golgys0621/topsdk/ability2939/domain"
)

type AlibabaAlihealthDrugLsydGetentinfolistResponse struct {

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
	Result domain.AlibabaAlihealthDrugLsydGetentinfolistResultModel `json:"result,omitempty" `
}
