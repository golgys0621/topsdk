package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyCreaterequestRequest struct {
	/*
	   入参     */
	SyRequestDto *domain.AlibabaAlihealthSynergySyCreaterequestSyRequestDTO `json:"sy_request_dto" required:"true" `
}

func (s *AlibabaAlihealthSynergySyCreaterequestRequest) SetSyRequestDto(v domain.AlibabaAlihealthSynergySyCreaterequestSyRequestDTO) *AlibabaAlihealthSynergySyCreaterequestRequest {
	s.SyRequestDto = &v
	return s
}

func (req *AlibabaAlihealthSynergySyCreaterequestRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.SyRequestDto != nil {
		paramMap["sy_request_dto"] = util.ConvertStruct(*req.SyRequestDto)
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergySyCreaterequestRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
