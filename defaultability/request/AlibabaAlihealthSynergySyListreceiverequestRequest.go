package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyListreceiverequestRequest struct {
	/*
	   入参     */
	SyRequestDto *domain.AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO `json:"sy_request_dto" required:"true" `
}

func (s *AlibabaAlihealthSynergySyListreceiverequestRequest) SetSyRequestDto(v domain.AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO) *AlibabaAlihealthSynergySyListreceiverequestRequest {
	s.SyRequestDto = &v
	return s
}

func (req *AlibabaAlihealthSynergySyListreceiverequestRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.SyRequestDto != nil {
		paramMap["sy_request_dto"] = util.ConvertStruct(*req.SyRequestDto)
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergySyListreceiverequestRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
