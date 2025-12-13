package ability305

import (
	"errors"
	"log"

	"github.com/golgys0621/topsdk"
	"github.com/golgys0621/topsdk/ability305/request"
	"github.com/golgys0621/topsdk/ability305/response"
	"github.com/golgys0621/topsdk/util"
)

type Ability305 struct {
	Client *topsdk.TopClient
}

func NewAbility305(client *topsdk.TopClient) *Ability305 {
	return &Ability305{client}
}

/*
获取消息队列积压情况
*/
func (ability *Ability305) TaobaoTmcQueueGet(req *request.TaobaoTmcQueueGetRequest) (*response.TaobaoTmcQueueGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability305 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tmc.queue.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTmcQueueGetResponse{}
	if err != nil {
		log.Println("taobaoTmcQueueGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取用户开通的topic列表
*/
func (ability *Ability305) TaobaoTmcUserTopicsGet(req *request.TaobaoTmcUserTopicsGetRequest) (*response.TaobaoTmcUserTopicsGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability305 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tmc.user.topics.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTmcUserTopicsGetResponse{}
	if err != nil {
		log.Println("taobaoTmcUserTopicsGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
