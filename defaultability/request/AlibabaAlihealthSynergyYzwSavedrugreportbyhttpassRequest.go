package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest struct {
	/*
	   企业ID     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   药品ID     */
	DrugEntBaseInfoId *string `json:"drug_ent_base_info_id" required:"true" `
	/*
	   批次号     */
	BatchNo *string `json:"batch_no" required:"true" `
	/*
	   报告文件     */
	Files *[]domain.AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassOnenetOcrPdfDTO `json:"files" required:"true" `
	/*
	   用户id(三方系统标识)     */
	UserId *string `json:"user_id" required:"true" `
	/*
	   委托企业签章     */
	AssSealInfo *domain.AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassDrugReportSealInfo `json:"ass_seal_info,omitempty" required:"false" `
	/*
	   药检报告号     */
	ReportNo *string `json:"report_no,omitempty" required:"false" `
	/*
	   报告日期     */
	ReportDate *string `json:"report_date,omitempty" required:"false" `
	/*
	   生产日期     */
	ProduceDate *string `json:"produce_date,omitempty" required:"false" `
	/*
	   本企业签章     */
	SelfSealInfo *domain.AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassDrugReportSealInfo `json:"self_seal_info,omitempty" required:"false" `
	/*
	   委托企业     */
	AssRefEntId *string `json:"ass_ref_ent_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest {
	s.DrugEntBaseInfoId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest) SetBatchNo(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest {
	s.BatchNo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest) SetFiles(v []domain.AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassOnenetOcrPdfDTO) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest {
	s.Files = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest) SetUserId(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest {
	s.UserId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest) SetAssSealInfo(v domain.AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassDrugReportSealInfo) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest {
	s.AssSealInfo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest) SetReportNo(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest {
	s.ReportNo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest) SetReportDate(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest {
	s.ReportDate = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest) SetProduceDate(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest {
	s.ProduceDate = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest) SetSelfSealInfo(v domain.AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassDrugReportSealInfo) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest {
	s.SelfSealInfo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest) SetAssRefEntId(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest {
	s.AssRefEntId = &v
	return s
}

func (req *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.DrugEntBaseInfoId != nil {
		paramMap["drug_ent_base_info_id"] = *req.DrugEntBaseInfoId
	}
	if req.BatchNo != nil {
		paramMap["batch_no"] = *req.BatchNo
	}
	if req.Files != nil {
		paramMap["files"] = util.ConvertStructList(*req.Files)
	}
	if req.UserId != nil {
		paramMap["user_id"] = *req.UserId
	}
	if req.AssSealInfo != nil {
		paramMap["ass_seal_info"] = util.ConvertStruct(*req.AssSealInfo)
	}
	if req.ReportNo != nil {
		paramMap["report_no"] = *req.ReportNo
	}
	if req.ReportDate != nil {
		paramMap["report_date"] = *req.ReportDate
	}
	if req.ProduceDate != nil {
		paramMap["produce_date"] = *req.ProduceDate
	}
	if req.SelfSealInfo != nil {
		paramMap["self_seal_info"] = util.ConvertStruct(*req.SelfSealInfo)
	}
	if req.AssRefEntId != nil {
		paramMap["ass_ref_ent_id"] = *req.AssRefEntId
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
