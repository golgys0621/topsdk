package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergyYzwSignednotstampedbillQueryRequest struct {
	/*
	   企业（本企业id）     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   上游出库单开始日期（最大查询区间一月）     */
	BeginTime *string `json:"begin_time" required:"true" `
	/*
	   上游出库单结束日期（最大查询区间一月）     */
	EndTime *string `json:"end_time" required:"true" `
	/*
	   分页对象     */
	Page *domain.AlibabaAlihealthSynergyYzwSignednotstampedbillQueryPage `json:"page" required:"true" `
}

func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryRequest) SetBeginTime(v string) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryRequest {
	s.BeginTime = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryRequest) SetEndTime(v string) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryRequest {
	s.EndTime = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryRequest) SetPage(v domain.AlibabaAlihealthSynergyYzwSignednotstampedbillQueryPage) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryRequest {
	s.Page = &v
	return s
}

func (req *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryRequest) ToMap() map[string]interface{} {
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

func (req *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
