package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResponse struct {

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
	Result domain.AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel `json:"result,omitempty" `
}
