package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergyYzwBillSignedQueryRequest struct {
	/*
	   企业（本企业id）     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   单据签收开始时间（最大查询区间一月）     */
	BeginTime *string `json:"begin_time" required:"true" `
	/*
	   单据签收截止时间（最大查询区间一月）     */
	EndTime *string `json:"end_time" required:"true" `
	/*
	   分页对象     */
	Page *domain.AlibabaAlihealthSynergyYzwBillSignedQueryPage `json:"page,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergyYzwBillSignedQueryRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryRequest) SetBeginTime(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryRequest {
	s.BeginTime = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryRequest) SetEndTime(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryRequest {
	s.EndTime = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryRequest) SetPage(v domain.AlibabaAlihealthSynergyYzwBillSignedQueryPage) *AlibabaAlihealthSynergyYzwBillSignedQueryRequest {
	s.Page = &v
	return s
}

func (req *AlibabaAlihealthSynergyYzwBillSignedQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.BeginTime != nil {
		paramMap["begin_time"] = *req.BeginTime
	}
	if req.EndTime != nil {
		paramMap["end_time"] = *req.EndTime
	}
	if req.Page != nil {
		paramMap["page"] = util.ConvertStruct(*req.Page)
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergyYzwBillSignedQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
