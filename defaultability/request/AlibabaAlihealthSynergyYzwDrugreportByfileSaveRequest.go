package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest struct {
	/*
	   批次号     */
	BatchNo *string `json:"batch_no" required:"true" `
	/*
	   药品ID     */
	DrugEntBaseInfoId *string `json:"drug_ent_base_info_id" required:"true" `
	/*
	   报告文件二进制字节数组     */
	ReportData *[]byte `json:"report_data" required:"true" `
	/*
	   三方用户标识     */
	UserId *string `json:"user_id" required:"true" `
	/*
	   代理操作企业(调用企业)的章信息     */
	SelfSealInfo *domain.AlibabaAlihealthSynergyYzwDrugreportByfileSaveDrugReportSealInfo `json:"self_seal_info,omitempty" required:"false" `
	/*
	   生产日期     */
	ProduceDate *string `json:"produce_date,omitempty" required:"false" `
	/*
	   空表示给自己上传报告，报告属于调用企业。填授权企业refEntId,表示调用企业替授权企业上传报告，报告属于授权企业。     */
	AssRefEntId *string `json:"ass_ref_ent_id,omitempty" required:"false" `
	/*
	   报告日期     */
	ReportDate *string `json:"report_date,omitempty" required:"false" `
	/*
	   授权单位的章信息     */
	AssSealInfo *domain.AlibabaAlihealthSynergyYzwDrugreportByfileSaveDrugReportSealInfo `json:"ass_seal_info,omitempty" required:"false" `
	/*
	   调用企业的refEntId     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   文件类型支持:pdf,jpg,jpeg, png     */
	FileType *string `json:"file_type" required:"true" `
	/*
	   报告编号     */
	ReportNo *string `json:"report_no,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest) SetBatchNo(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest {
	s.BatchNo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest {
	s.DrugEntBaseInfoId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest) SetReportData(v []byte) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest {
	s.ReportData = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest) SetUserId(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest {
	s.UserId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest) SetSelfSealInfo(v domain.AlibabaAlihealthSynergyYzwDrugreportByfileSaveDrugReportSealInfo) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest {
	s.SelfSealInfo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest) SetProduceDate(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest {
	s.ProduceDate = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest) SetAssRefEntId(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest {
	s.AssRefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest) SetReportDate(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest {
	s.ReportDate = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest) SetAssSealInfo(v domain.AlibabaAlihealthSynergyYzwDrugreportByfileSaveDrugReportSealInfo) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest {
	s.AssSealInfo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest) SetFileType(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest {
	s.FileType = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest) SetReportNo(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest {
	s.ReportNo = &v
	return s
}

func (req *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.BatchNo != nil {
		paramMap["batch_no"] = *req.BatchNo
	}
	if req.DrugEntBaseInfoId != nil {
		paramMap["drug_ent_base_info_id"] = *req.DrugEntBaseInfoId
	}
	if req.UserId != nil {
		paramMap["user_id"] = *req.UserId
	}
	if req.SelfSealInfo != nil {
		paramMap["self_seal_info"] = util.ConvertStruct(*req.SelfSealInfo)
	}
	if req.ProduceDate != nil {
		paramMap["produce_date"] = *req.ProduceDate
	}
	if req.AssRefEntId != nil {
		paramMap["ass_ref_ent_id"] = *req.AssRefEntId
	}
	if req.ReportDate != nil {
		paramMap["report_date"] = *req.ReportDate
	}
	if req.AssSealInfo != nil {
		paramMap["ass_seal_info"] = util.ConvertStruct(*req.AssSealInfo)
	}
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.FileType != nil {
		paramMap["file_type"] = *req.FileType
	}
	if req.ReportNo != nil {
		paramMap["report_no"] = *req.ReportNo
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	if req.ReportData != nil {
		fileMap["report_data"] = *req.ReportData
	}
	return fileMap
}
