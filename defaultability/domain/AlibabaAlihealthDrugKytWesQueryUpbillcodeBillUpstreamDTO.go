package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO struct {
	/*
	   发货企业名称     */
	FromUserName *string `json:"from_user_name,omitempty" `

	/*
	   单据时间     */
	BillTime *util.LocalTime `json:"bill_time,omitempty" `

	/*
	   货主     */
	RefUserId *string `json:"ref_user_id,omitempty" `

	/*
	   发货企业ID     */
	FromUserId *string `json:"from_user_id,omitempty" `

	/*
	   单据类型     */
	BillType *string `json:"bill_type,omitempty" `

	/*
	   收货企业名称     */
	ToUserName *string `json:"to_user_name,omitempty" `

	/*
	   单据号     */
	BillCode *string `json:"bill_code,omitempty" `

	/*
	   收货企业ID     */
	ToUserId *string `json:"to_user_id,omitempty" `

	/*
	   收货企业REF_ENT_ID     */
	ToRefUserId *string `json:"to_ref_user_id,omitempty" `

	/*
	   发货企业REF_ENT_ID     */
	FromRefUserId *string `json:"from_ref_user_id,omitempty" `

	/*
	   委托企业refEntId     */
	AssRefEntId *string `json:"ass_ref_ent_id,omitempty" `

	/*
	   委托企业entId     */
	AssEntId *string `json:"ass_ent_id,omitempty" `

	/*
	   委托企业entId对应的名称     */
	AssEntName *string `json:"ass_ent_name,omitempty" `
}

func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO) SetFromUserName(v string) *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO {
	s.FromUserName = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO) SetBillTime(v util.LocalTime) *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO {
	s.BillTime = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO) SetRefUserId(v string) *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO {
	s.RefUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO) SetFromUserId(v string) *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO {
	s.FromUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO) SetBillType(v string) *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO {
	s.BillType = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO) SetToUserName(v string) *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO {
	s.ToUserName = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO) SetBillCode(v string) *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO {
	s.BillCode = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO) SetToUserId(v string) *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO {
	s.ToUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO) SetToRefUserId(v string) *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO {
	s.ToRefUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO) SetFromRefUserId(v string) *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO {
	s.FromRefUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO) SetAssRefEntId(v string) *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO {
	s.AssRefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO) SetAssEntId(v string) *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO {
	s.AssEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO) SetAssEntName(v string) *AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO {
	s.AssEntName = &v
	return s
}
