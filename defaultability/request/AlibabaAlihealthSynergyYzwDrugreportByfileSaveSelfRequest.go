package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest struct {
	/*
	   药检报告批次号     */
	BatchNo *string `json:"batch_no" required:"true" `
	/*
	   药品生产日期     */
	ProduceDate *string `json:"produce_date,omitempty" required:"false" `
	/*
	   报告日期     */
	ReportDate *string `json:"report_date,omitempty" required:"false" `
	/*
	   药品ID     */
	DrugEntBaseInfoId *string `json:"drug_ent_base_info_id" required:"true" `
	/*
	   报告文件二进制字节数组     */
	ReportData *[]byte `json:"report_data" required:"true" `
	/*
	   授权企业的相关印章信息     */
	AuthSealInfo *domain.AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfDrugReportSealInfo `json:"auth_seal_info,omitempty" required:"false" `
	/*
	   上传报告企业的refEntId     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   三方用户标识     */
	UserId *string `json:"user_id" required:"true" `
	/*
	   文件类型支持pdf,jpg,jpeg,png     */
	FileType *string `json:"file_type" required:"true" `
	/*
	   报告编号     */
	ReportNo *string `json:"report_no,omitempty" required:"false" `
	/*
	   调用企业的章信息     */
	SelfSealInfo *domain.AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfDrugReportSealInfo `json:"self_seal_info,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest) SetBatchNo(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest {
	s.BatchNo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest) SetProduceDate(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest {
	s.ProduceDate = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest) SetReportDate(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest {
	s.ReportDate = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest {
	s.DrugEntBaseInfoId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest) SetReportData(v []byte) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest {
	s.ReportData = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest) SetAuthSealInfo(v domain.AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfDrugReportSealInfo) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest {
	s.AuthSealInfo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest) SetUserId(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest {
	s.UserId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest) SetFileType(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest {
	s.FileType = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest) SetReportNo(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest {
	s.ReportNo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest) SetSelfSealInfo(v domain.AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfDrugReportSealInfo) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest {
	s.SelfSealInfo = &v
	return s
}

func (req *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.BatchNo != nil {
		paramMap["batch_no"] = *req.BatchNo
	}
	if req.ProduceDate != nil {
		paramMap["produce_date"] = *req.ProduceDate
	}
	if req.ReportDate != nil {
		paramMap["report_date"] = *req.ReportDate
	}
	if req.DrugEntBaseInfoId != nil {
		paramMap["drug_ent_base_info_id"] = *req.DrugEntBaseInfoId
	}
	if req.AuthSealInfo != nil {
		paramMap["auth_seal_info"] = util.ConvertStruct(*req.AuthSealInfo)
	}
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.UserId != nil {
		paramMap["user_id"] = *req.UserId
	}
	if req.FileType != nil {
		paramMap["file_type"] = *req.FileType
	}
	if req.ReportNo != nil {
		paramMap["report_no"] = *req.ReportNo
	}
	if req.SelfSealInfo != nil {
		paramMap["self_seal_info"] = util.ConvertStruct(*req.SelfSealInfo)
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	if req.ReportData != nil {
		fileMap["report_data"] = *req.ReportData
	}
	return fileMap
}
