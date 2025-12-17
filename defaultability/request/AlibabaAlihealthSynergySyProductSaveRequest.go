package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyProductSaveRequest struct {
	/*
	   产品参数     */
	SyProductRequestDto *domain.AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO `json:"sy_product_request_dto" required:"true" `
}

func (s *AlibabaAlihealthSynergySyProductSaveRequest) SetSyProductRequestDto(v domain.AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) *AlibabaAlihealthSynergySyProductSaveRequest {
	s.SyProductRequestDto = &v
	return s
}

func (req *AlibabaAlihealthSynergySyProductSaveRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.SyProductRequestDto != nil {
		paramMap["sy_product_request_dto"] = util.ConvertStruct(*req.SyProductRequestDto)
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergySyProductSaveRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
