package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyContractCreateRequest struct {
	/*
	   合同参数     */
	SyContractCreateTopRequest *domain.AlibabaAlihealthSynergySyContractCreateSyContractCreateTopRequest `json:"sy_contract_create_top_request" required:"true" `
}

func (s *AlibabaAlihealthSynergySyContractCreateRequest) SetSyContractCreateTopRequest(v domain.AlibabaAlihealthSynergySyContractCreateSyContractCreateTopRequest) *AlibabaAlihealthSynergySyContractCreateRequest {
	s.SyContractCreateTopRequest = &v
	return s
}

func (req *AlibabaAlihealthSynergySyContractCreateRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.SyContractCreateTopRequest != nil {
		paramMap["sy_contract_create_top_request"] = util.ConvertStruct(*req.SyContractCreateTopRequest)
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergySyContractCreateRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
