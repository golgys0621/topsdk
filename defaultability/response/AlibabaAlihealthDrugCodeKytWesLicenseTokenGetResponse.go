package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthDrugCodeKytWesLicenseTokenGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回值
	*/
	Result domain.AlibabaAlihealthDrugCodeKytWesLicenseTokenGetResultModel `json:"result,omitempty" `
}
