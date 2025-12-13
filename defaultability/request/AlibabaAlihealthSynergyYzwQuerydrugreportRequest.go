package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergyYzwQuerydrugreportRequest struct {
	/*
	   企业（本企业id）     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   发货企业     */
	FromRefEntId *string `json:"from_ref_ent_id,omitempty" required:"false" `
	/*
	   单据编号     */
	BillCode *string `json:"bill_code,omitempty" required:"false" `
	/*
	   开始日期（最大查询区间一月）     */
	BeginTime *string `json:"begin_time" required:"true" `
	/*
	   结束日期（最大查询区间一月）     */
	EndTime *string `json:"end_time" required:"true" `
	/*
	   药品ID     */
	DrugId *string `json:"drug_id,omitempty" required:"false" `
	/*
	   批号     */
	BatchNo *string `json:"batch_no,omitempty" required:"false" `
	/*
	   分页对象     */
	Page *domain.AlibabaAlihealthSynergyYzwQuerydrugreportPage `json:"page" required:"true" `
}

func (s *AlibabaAlihealthSynergyYzwQuerydrugreportRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportRequest) SetFromRefEntId(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportRequest {
	s.FromRefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportRequest) SetBillCode(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportRequest {
	s.BillCode = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportRequest) SetBeginTime(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportRequest {
	s.BeginTime = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportRequest) SetEndTime(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportRequest {
	s.EndTime = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportRequest) SetDrugId(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportRequest {
	s.DrugId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportRequest) SetBatchNo(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportRequest {
	s.BatchNo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportRequest) SetPage(v domain.AlibabaAlihealthSynergyYzwQuerydrugreportPage) *AlibabaAlihealthSynergyYzwQuerydrugreportRequest {
	s.Page = &v
	return s
}

func (req *AlibabaAlihealthSynergyYzwQuerydrugreportRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.FromRefEntId != nil {
		paramMap["from_ref_ent_id"] = *req.FromRefEntId
	}
	if req.BillCode != nil {
		paramMap["bill_code"] = *req.BillCode
	}
	if req.BeginTime != nil {
		paramMap["begin_time"] = *req.BeginTime
	}
	if req.EndTime != nil {
		paramMap["end_time"] = *req.EndTime
	}
	if req.DrugId != nil {
		paramMap["drug_id"] = *req.DrugId
	}
	if req.BatchNo != nil {
		paramMap["batch_no"] = *req.BatchNo
	}
	if req.Page != nil {
		paramMap["page"] = util.ConvertStruct(*req.Page)
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergyYzwQuerydrugreportRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
