package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergyYzwDrugreportAssSealRequest struct {
	/*
	   授权企业refEntId，报告属于该企业。不传表示没有委托场景     */
	AssRefEntId *string `json:"ass_ref_ent_id,omitempty" required:"false" `
	/*
	   药检报告id     */
	ReportId *int64 `json:"report_id" required:"true" `
	/*
	   授权单位的章信息     */
	AssSealInfo *domain.AlibabaAlihealthSynergyYzwDrugreportAssSealDrugReportSealInfo `json:"ass_seal_info,omitempty" required:"false" `
	/*
	   调用企业的refEntId     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   三方用户标识     */
	UserId *string `json:"user_id" required:"true" `
	/*
	   代理操作企业(调用企业)的章信息     */
	SelfSealInfo *domain.AlibabaAlihealthSynergyYzwDrugreportAssSealDrugReportSealInfo `json:"self_seal_info,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergyYzwDrugreportAssSealRequest) SetAssRefEntId(v string) *AlibabaAlihealthSynergyYzwDrugreportAssSealRequest {
	s.AssRefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportAssSealRequest) SetReportId(v int64) *AlibabaAlihealthSynergyYzwDrugreportAssSealRequest {
	s.ReportId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportAssSealRequest) SetAssSealInfo(v domain.AlibabaAlihealthSynergyYzwDrugreportAssSealDrugReportSealInfo) *AlibabaAlihealthSynergyYzwDrugreportAssSealRequest {
	s.AssSealInfo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportAssSealRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwDrugreportAssSealRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportAssSealRequest) SetUserId(v string) *AlibabaAlihealthSynergyYzwDrugreportAssSealRequest {
	s.UserId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportAssSealRequest) SetSelfSealInfo(v domain.AlibabaAlihealthSynergyYzwDrugreportAssSealDrugReportSealInfo) *AlibabaAlihealthSynergyYzwDrugreportAssSealRequest {
	s.SelfSealInfo = &v
	return s
}

func (req *AlibabaAlihealthSynergyYzwDrugreportAssSealRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.AssRefEntId != nil {
		paramMap["ass_ref_ent_id"] = *req.AssRefEntId
	}
	if req.ReportId != nil {
		paramMap["report_id"] = *req.ReportId
	}
	if req.AssSealInfo != nil {
		paramMap["ass_seal_info"] = util.ConvertStruct(*req.AssSealInfo)
	}
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.UserId != nil {
		paramMap["user_id"] = *req.UserId
	}
	if req.SelfSealInfo != nil {
		paramMap["self_seal_info"] = util.ConvertStruct(*req.SelfSealInfo)
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergyYzwDrugreportAssSealRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
