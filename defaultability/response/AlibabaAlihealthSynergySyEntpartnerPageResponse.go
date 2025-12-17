package response

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
)

type AlibabaAlihealthSynergySyEntpartnerPageResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   result
	*/
	Result domain.AlibabaAlihealthSynergySyEntpartnerPageResult `json:"result,omitempty" `
}
