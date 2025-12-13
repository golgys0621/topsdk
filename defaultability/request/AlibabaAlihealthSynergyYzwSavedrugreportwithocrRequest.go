package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergyYzwSavedrugreportwithocrRequest struct {
	/*
	   企业ID     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   报告数据 最大100个文件     */
	ReportData *[]domain.AlibabaAlihealthSynergyYzwSavedrugreportwithocrTopFileDTO `json:"report_data" required:"true" `
}

func (s *AlibabaAlihealthSynergyYzwSavedrugreportwithocrRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwSavedrugreportwithocrRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportwithocrRequest) SetReportData(v []domain.AlibabaAlihealthSynergyYzwSavedrugreportwithocrTopFileDTO) *AlibabaAlihealthSynergyYzwSavedrugreportwithocrRequest {
	s.ReportData = &v
	return s
}

func (req *AlibabaAlihealthSynergyYzwSavedrugreportwithocrRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.ReportData != nil {
		paramMap["report_data"] = util.ConvertStructList(*req.ReportData)
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergyYzwSavedrugreportwithocrRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
