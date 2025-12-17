package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyContractSealRequest struct {
	/*
	   参数     */
	SyContractSealRequest *domain.AlibabaAlihealthSynergySyContractSealSyContractSealRequest `json:"sy_contract_seal_request" required:"true" `
}

func (s *AlibabaAlihealthSynergySyContractSealRequest) SetSyContractSealRequest(v domain.AlibabaAlihealthSynergySyContractSealSyContractSealRequest) *AlibabaAlihealthSynergySyContractSealRequest {
	s.SyContractSealRequest = &v
	return s
}

func (req *AlibabaAlihealthSynergySyContractSealRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.SyContractSealRequest != nil {
		paramMap["sy_contract_seal_request"] = util.ConvertStruct(*req.SyContractSealRequest)
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergySyContractSealRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
