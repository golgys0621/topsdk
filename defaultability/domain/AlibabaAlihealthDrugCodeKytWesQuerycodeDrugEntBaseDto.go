package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto struct {
	/*
	   药品类型描述     */
	PhysicTypeDesc *string `json:"physic_type_desc,omitempty" `

	/*
	   药品名称     */
	PhysicName *string `json:"physic_name,omitempty" `

	/*
	   有效期     */
	Exprie *string `json:"exprie,omitempty" `

	/*
	   药品id     */
	DrugEntBaseInfoId *string `json:"drug_ent_base_info_id,omitempty" `

	/*
	   批准文号     */
	ApprovalLicenceNo *string `json:"approval_licence_no,omitempty" `

	/*
	   包装规格     */
	PkgSpecCrit *string `json:"pkg_spec_crit,omitempty" `

	/*
	   制剂规格     */
	PrepnSpec *string `json:"prepn_spec,omitempty" `

	/*
	   剂型描述     */
	PrepnTypeDesc *string `json:"prepn_type_desc,omitempty" `

	/*
	   小包下的制剂数量     */
	PkgNum *string `json:"pkg_num,omitempty" `

	/*
	   药品原批准文号     */
	ApprovalLicenceNoOld *string `json:"approval_licence_no_old,omitempty" `

	/*
	   原药品批准文号有效期     */
	ApprovalLicenceExpiryOld *util.LocalTime `json:"approval_licence_expiry_old,omitempty" `

	/*
	   批准日期     */
	ApprovalLicenceDate *util.LocalTime `json:"approval_licence_date,omitempty" `

	/*
	   批准文号有效期     */
	ApprovalLicenceExpiry *util.LocalTime `json:"approval_licence_expiry,omitempty" `
}

func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto) SetPhysicTypeDesc(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto {
	s.PhysicTypeDesc = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto) SetPhysicName(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto {
	s.PhysicName = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto) SetExprie(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto {
	s.Exprie = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto {
	s.DrugEntBaseInfoId = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto) SetApprovalLicenceNo(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto {
	s.ApprovalLicenceNo = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto) SetPkgSpecCrit(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto {
	s.PkgSpecCrit = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto) SetPrepnSpec(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto {
	s.PrepnSpec = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto) SetPrepnTypeDesc(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto {
	s.PrepnTypeDesc = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto) SetPkgNum(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto {
	s.PkgNum = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto) SetApprovalLicenceNoOld(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto {
	s.ApprovalLicenceNoOld = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto) SetApprovalLicenceExpiryOld(v util.LocalTime) *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto {
	s.ApprovalLicenceExpiryOld = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto) SetApprovalLicenceDate(v util.LocalTime) *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto {
	s.ApprovalLicenceDate = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto) SetApprovalLicenceExpiry(v util.LocalTime) *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto {
	s.ApprovalLicenceExpiry = &v
	return s
}
