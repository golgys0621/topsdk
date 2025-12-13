package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthSynergyYzwGetentinfolistResponse struct {

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
	Result domain.AlibabaAlihealthSynergyYzwGetentinfolistResultModel `json:"result,omitempty" `
}
