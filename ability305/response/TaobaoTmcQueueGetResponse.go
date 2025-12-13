package response

import (
	"github.com/golgys0621/topsdk/ability305/domain"
)

type TaobaoTmcQueueGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   队列详细信息
	*/
	Datas []domain.TaobaoTmcQueueGetTmcQueueInfo `json:"datas,omitempty" `
}
