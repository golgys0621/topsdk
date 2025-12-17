package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySySignatureResourceSealRequest struct {
	/*
	   签章id     */
	SySignatureId *int64 `json:"sy_signature_id" required:"true" `
	/*
	   发送企业refEntId     */
	RefEntId *string `json:"ref_ent_id,omitempty" required:"false" `
	/*
	   资料的盖章信息     */
	SealInfos *[]domain.AlibabaAlihealthSynergySySignatureResourceSealSealInfo `json:"seal_infos" required:"true" `
	/*
	   签章资料id     */
	SySignatureResourceId *int64 `json:"sy_signature_resource_id" required:"true" `
	/*
	   操作人     */
	OptName *string `json:"opt_name" required:"true" `
}

func (s *AlibabaAlihealthSynergySySignatureResourceSealRequest) SetSySignatureId(v int64) *AlibabaAlihealthSynergySySignatureResourceSealRequest {
	s.SySignatureId = &v
	return s
}
func (s *AlibabaAlihealthSynergySySignatureResourceSealRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySySignatureResourceSealRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySySignatureResourceSealRequest) SetSealInfos(v []domain.AlibabaAlihealthSynergySySignatureResourceSealSealInfo) *AlibabaAlihealthSynergySySignatureResourceSealRequest {
	s.SealInfos = &v
	return s
}
func (s *AlibabaAlihealthSynergySySignatureResourceSealRequest) SetSySignatureResourceId(v int64) *AlibabaAlihealthSynergySySignatureResourceSealRequest {
	s.SySignatureResourceId = &v
	return s
}
func (s *AlibabaAlihealthSynergySySignatureResourceSealRequest) SetOptName(v string) *AlibabaAlihealthSynergySySignatureResourceSealRequest {
	s.OptName = &v
	return s
}

func (req *AlibabaAlihealthSynergySySignatureResourceSealRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.SySignatureId != nil {
		paramMap["sy_signature_id"] = *req.SySignatureId
	}
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.SealInfos != nil {
		paramMap["seal_infos"] = util.ConvertStructList(*req.SealInfos)
	}
	if req.SySignatureResourceId != nil {
		paramMap["sy_signature_resource_id"] = *req.SySignatureResourceId
	}
	if req.OptName != nil {
		paramMap["opt_name"] = *req.OptName
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergySySignatureResourceSealRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
