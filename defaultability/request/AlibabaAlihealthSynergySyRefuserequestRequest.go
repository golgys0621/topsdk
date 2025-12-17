package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyRefuserequestRequest struct {
	/*
	   入参     */
	SyRequestDto *domain.AlibabaAlihealthSynergySyRefuserequestSyRequestDTO `json:"sy_request_dto" required:"true" `
}

func (s *AlibabaAlihealthSynergySyRefuserequestRequest) SetSyRequestDto(v domain.AlibabaAlihealthSynergySyRefuserequestSyRequestDTO) *AlibabaAlihealthSynergySyRefuserequestRequest {
	s.SyRequestDto = &v
	return s
}

func (req *AlibabaAlihealthSynergySyRefuserequestRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.SyRequestDto != nil {
		paramMap["sy_request_dto"] = util.ConvertStruct(*req.SyRequestDto)
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergySyRefuserequestRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
