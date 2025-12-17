package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO struct {
	/*
	   批准文号     */
	ApprovalNo *string `json:"approval_no,omitempty" `

	/*
	   gmtModified     */
	GmtModified *util.LocalTime `json:"gmt_modified,omitempty" `

	/*
	   sy_request主键id     */
	SyRequestId *int64 `json:"sy_request_id,omitempty" `

	/*
	   gmtCreate     */
	GmtCreate *util.LocalTime `json:"gmt_create,omitempty" `

	/*
	   产品名称     */
	ProductName *string `json:"product_name,omitempty" `

	/*
	   产品大类     */
	ProductTypeDesc *string `json:"product_type_desc,omitempty" `

	/*
	   供应商     */
	SupplyEntName *string `json:"supply_ent_name,omitempty" `

	/*
	   资料数     */
	ResourceCount *int64 `json:"resource_count,omitempty" `

	/*
	   syProductId     */
	SyProductId *int64 `json:"sy_product_id,omitempty" `

	/*
	   本企业     */
	RefEntId *string `json:"ref_ent_id,omitempty" `

	/*
	   id     */
	Id *int64 `json:"id,omitempty" `

	/*
	   productTypeCode     */
	ProductTypeCode *int64 `json:"product_type_code,omitempty" `
}

func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO) SetApprovalNo(v string) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO {
	s.ApprovalNo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO) SetGmtModified(v util.LocalTime) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO {
	s.GmtModified = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO) SetSyRequestId(v int64) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO {
	s.SyRequestId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO) SetGmtCreate(v util.LocalTime) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO {
	s.GmtCreate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO) SetProductName(v string) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO {
	s.ProductName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO) SetProductTypeDesc(v string) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO {
	s.ProductTypeDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO) SetSupplyEntName(v string) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO {
	s.SupplyEntName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO) SetResourceCount(v int64) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO {
	s.ResourceCount = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO) SetSyProductId(v int64) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO {
	s.SyProductId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO) SetRefEntId(v string) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO) SetId(v int64) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO {
	s.Id = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO) SetProductTypeCode(v int64) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO {
	s.ProductTypeCode = &v
	return s
}
