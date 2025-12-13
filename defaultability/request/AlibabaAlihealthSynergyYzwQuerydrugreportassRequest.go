package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergyYzwQuerydrugreportassRequest struct {
	/*
	   本企业     */
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
	Page *domain.AlibabaAlihealthSynergyYzwQuerydrugreportassPage `json:"page" required:"true" `
	/*
	   委托企业/货主企业     */
	AssRefEntId *string `json:"ass_ref_ent_id,omitempty" required:"false" `
	/*
	   0:按照单据时间 1:按照单据上传时间 2：按照单据签收时间 defalutValue��0    */
	TimeType *string `json:"time_type" required:"true" `
	/*
	   0:待盖章 5:已盖章 6:同批号已盖章 注意：status_type需要填3      */
	SealStatus *string `json:"seal_status,omitempty" required:"false" `
	/*
	   2:待签收 3:已签收 4:已拒绝     */
	StatusType *string `json:"status_type,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest) SetFromRefEntId(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest {
	s.FromRefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest) SetBillCode(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest {
	s.BillCode = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest) SetBeginTime(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest {
	s.BeginTime = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest) SetEndTime(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest {
	s.EndTime = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest) SetDrugId(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest {
	s.DrugId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest) SetBatchNo(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest {
	s.BatchNo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest) SetPage(v domain.AlibabaAlihealthSynergyYzwQuerydrugreportassPage) *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest {
	s.Page = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest) SetAssRefEntId(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest {
	s.AssRefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest) SetTimeType(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest {
	s.TimeType = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest) SetSealStatus(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest {
	s.SealStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest) SetStatusType(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest {
	s.StatusType = &v
	return s
}

func (req *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest) ToMap() map[string]interface{} {
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
	if req.AssRefEntId != nil {
		paramMap["ass_ref_ent_id"] = *req.AssRefEntId
	}
	if req.TimeType != nil {
		paramMap["time_type"] = *req.TimeType
	}
	if req.SealStatus != nil {
		paramMap["seal_status"] = *req.SealStatus
	}
	if req.StatusType != nil {
		paramMap["status_type"] = *req.StatusType
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergyYzwQuerydrugreportassRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
