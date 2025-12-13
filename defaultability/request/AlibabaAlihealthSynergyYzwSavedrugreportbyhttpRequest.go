package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest struct {
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
	Files *[]domain.AlibabaAlihealthSynergyYzwSavedrugreportbyhttpOnenetOcrPdfDTO `json:"files" required:"true" `
	/*
	   用户id(三方系统标识)     */
	UserId *string `json:"user_id" required:"true" `
	/*
	   印章名称     */
	SealName *string `json:"seal_name,omitempty" required:"false" `
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
	   X坐标     */
	PositionX *int64 `json:"position_x,omitempty" required:"false" `
	/*
	   Y坐标     */
	PositionY *int64 `json:"position_y,omitempty" required:"false" `
	/*
	   是否要显示日期 defalutValue��false    */
	NoDate *bool `json:"no_date,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest {
	s.DrugEntBaseInfoId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest) SetBatchNo(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest {
	s.BatchNo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest) SetFiles(v []domain.AlibabaAlihealthSynergyYzwSavedrugreportbyhttpOnenetOcrPdfDTO) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest {
	s.Files = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest) SetUserId(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest {
	s.UserId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest) SetSealName(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest {
	s.SealName = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest) SetReportNo(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest {
	s.ReportNo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest) SetReportDate(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest {
	s.ReportDate = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest) SetProduceDate(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest {
	s.ProduceDate = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest) SetPositionX(v int64) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest {
	s.PositionX = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest) SetPositionY(v int64) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest {
	s.PositionY = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest) SetNoDate(v bool) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest {
	s.NoDate = &v
	return s
}

func (req *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest) ToMap() map[string]interface{} {
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
	if req.SealName != nil {
		paramMap["seal_name"] = *req.SealName
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
	if req.PositionX != nil {
		paramMap["position_x"] = *req.PositionX
	}
	if req.PositionY != nil {
		paramMap["position_y"] = *req.PositionY
	}
	if req.NoDate != nil {
		paramMap["no_date"] = *req.NoDate
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
