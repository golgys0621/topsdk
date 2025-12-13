package response

import (
	"github.com/golgys0621/topsdk/ability2940/domain"
)

type AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   监控宝推送网站监控信息，返回结果
	*/
	Result domain.AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoResultModel `json:"result,omitempty" `
}
