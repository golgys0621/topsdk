package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto struct {
	/*
	   药品信息     */
	PhysicInfo *string `json:"physic_info,omitempty" `

	/*
	   企业id     */
	RefEntId *string `json:"ref_ent_id,omitempty" `

	/*
	   药品包装规格     */
	PkgSpec *string `json:"pkg_spec,omitempty" `

	/*
	   药品制剂规格     */
	PrepnSpec *string `json:"prepn_spec,omitempty" `

	/*
	   药品制剂类型     */
	PrepnType *string `json:"prepn_type,omitempty" `

	/*
	   药品通用名称     */
	PhysicName *string `json:"physic_name,omitempty" `

	/*
	   药品包装比例     */
	PkgRatio *string `json:"pkg_ratio,omitempty" `

	/*
	   药品有效期至     */
	ExprieDate *string `json:"exprie_date,omitempty" `

	/*
	   药品生产批次号     */
	ProduceBatchNo *string `json:"produce_batch_no,omitempty" `

	/*
	   药品生产日期     */
	ProduceDate *string `json:"produce_date,omitempty" `

	/*
	   药品自类编码     */
	SubTypeNo *string `json:"sub_type_no,omitempty" `

	/*
	   药品编号     */
	ProductCode *string `json:"product_code,omitempty" `

	/*
	   药品ID     */
	ProdId *string `json:"prod_id,omitempty" `

	/*
	   批准文号     */
	ApproveNo *string `json:"approve_no,omitempty" `

	/*
	   药品类型     */
	PhysicType *string `json:"physic_type,omitempty" `

	/*
	   原批准文号有效期     */
	ApprovalLicenceExpiryOld *util.LocalTime `json:"approval_licence_expiry_old,omitempty" `

	/*
	   批准文号批准日期     */
	ApprovalLicenceDate *util.LocalTime `json:"approval_licence_date,omitempty" `

	/*
	   批准文号有效期至     */
	ApprovalLicenceExpiry *util.LocalTime `json:"approval_licence_expiry,omitempty" `

	/*
	   原批准文号     */
	ApprovalLicenceNoOld *string `json:"approval_licence_no_old,omitempty" `

	/*
	   生产日期     */
	OriginalProduceDate *string `json:"original_produce_date,omitempty" `

	/*
	   有效期至     */
	OriginalExpireDate *string `json:"original_expire_date,omitempty" `
}

func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetPhysicInfo(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.PhysicInfo = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetRefEntId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetPkgSpec(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.PkgSpec = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetPrepnSpec(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.PrepnSpec = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetPrepnType(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.PrepnType = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetPhysicName(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.PhysicName = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetPkgRatio(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.PkgRatio = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetExprieDate(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.ExprieDate = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetProduceBatchNo(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.ProduceBatchNo = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetProduceDate(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.ProduceDate = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetSubTypeNo(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.SubTypeNo = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetProductCode(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.ProductCode = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetProdId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.ProdId = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetApproveNo(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.ApproveNo = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetPhysicType(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.PhysicType = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetApprovalLicenceExpiryOld(v util.LocalTime) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.ApprovalLicenceExpiryOld = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetApprovalLicenceDate(v util.LocalTime) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.ApprovalLicenceDate = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetApprovalLicenceExpiry(v util.LocalTime) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.ApprovalLicenceExpiry = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetApprovalLicenceNoOld(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.ApprovalLicenceNoOld = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetOriginalProduceDate(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.OriginalProduceDate = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) SetOriginalExpireDate(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto {
	s.OriginalExpireDate = &v
	return s
}
