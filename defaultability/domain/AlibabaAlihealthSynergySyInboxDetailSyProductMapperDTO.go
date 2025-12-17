package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO struct {
	/*
	   产品名称     */
	ProductName *string `json:"product_name,omitempty" `

	/*
	   产品英文名     */
	ProductNameEn *string `json:"product_name_en,omitempty" `

	/*
	   批准文号     */
	ApprovalNo *string `json:"approval_no,omitempty" `

	/*
	   批准日期     */
	ApprovalDate *util.LocalTime `json:"approval_date,omitempty" `

	/*
	   批准有效期截止日期     */
	ApprovalEndDate *util.LocalTime `json:"approval_end_date,omitempty" `

	/*
	   商品名称(一般指药品的商品名称)     */
	DrugProductName *string `json:"drug_product_name,omitempty" `

	/*
	   剂型     */
	PrepnTypeDesc *string `json:"prepn_type_desc,omitempty" `

	/*
	   制剂规格     */
	PrepnSpec *string `json:"prepn_spec,omitempty" `

	/*
	   包装规格     */
	PkgSpec *string `json:"pkg_spec,omitempty" `

	/*
	   规格     */
	Spec *string `json:"spec,omitempty" `

	/*
	   药品类别，0中药、1中药饮片、2化学药、3生物制剂、4辅料     */
	DrugTypeCode *int64 `json:"drug_type_code,omitempty" `

	/*
	   药品类别     */
	DrugTypeDesc *string `json:"drug_type_desc,omitempty" `

	/*
	   0国产、1进口     */
	Source *int64 `json:"source,omitempty" `

	/*
	   产品类型，0药品、1器械、2其他     */
	ProductTypeCode *int64 `json:"product_type_code,omitempty" `

	/*
	   产品类型     */
	ProductTypeDesc *string `json:"product_type_desc,omitempty" `

	/*
	   生产企业名称     */
	ProduceEntName *string `json:"produce_ent_name,omitempty" `

	/*
	   上市许可持有人     */
	Mah *string `json:"mah,omitempty" `

	/*
	   生产企业地址     */
	ProduceAddress *string `json:"produce_address,omitempty" `

	/*
	   供应商企业名称     */
	SupplyEntName *string `json:"supply_ent_name,omitempty" `

	/*
	   供应商企业id     */
	SupplyRefEntId *string `json:"supply_ref_ent_id,omitempty" `

	/*
	   药品本位码     */
	SdcCode *string `json:"sdc_code,omitempty" `

	/*
	   药品本位码备注     */
	SdcCodeNote *string `json:"sdc_code_note,omitempty" `

	/*
	   产品编码     */
	ProductNo *string `json:"product_no,omitempty" `

	/*
	   供应商产品编码     */
	SupplyProductNo *string `json:"supply_product_no,omitempty" `

	/*
	   原批准文号/ 注册证号     */
	ApprovalNoOld *string `json:"approval_no_old,omitempty" `

	/*
	   执行标准     */
	ExecuteStandard *string `json:"execute_standard,omitempty" `

	/*
	   分包装批准文号     */
	PkgApprovalNo *string `json:"pkg_approval_no,omitempty" `

	/*
	   分包装批准文号批准日期     */
	PkgApprovalDate *util.LocalTime `json:"pkg_approval_date,omitempty" `

	/*
	   分包装批准文号有效期至     */
	PkgApprovalExpireDate *util.LocalTime `json:"pkg_approval_expire_date,omitempty" `

	/*
	   分包装企业名称     */
	PkgEntName *string `json:"pkg_ent_name,omitempty" `

	/*
	   分包装企业地址     */
	PkgEntAddress *string `json:"pkg_ent_address,omitempty" `

	/*
	   器械注册证号     */
	RegisterNo *string `json:"register_no,omitempty" `

	/*
	   器械注册日期     */
	MdRegisterDate *util.LocalTime `json:"md_register_date,omitempty" `

	/*
	   器械注册有效期截止日期     */
	MdRegisterEndDate *util.LocalTime `json:"md_register_end_date,omitempty" `

	/*
	   器械管理类别, 1第一类2第二类3第三类，器械必传     */
	MdManageTypeCode *int64 `json:"md_manage_type_code,omitempty" `

	/*
	   器械管理类别     */
	MdManageTypeDesc *string `json:"md_manage_type_desc,omitempty" `

	/*
	   器械分类编码     */
	MdClassifyNo *string `json:"md_classify_no,omitempty" `

	/*
	   器械结构以及组成成分     */
	MdComponent *string `json:"md_component,omitempty" `

	/*
	   器械适用范围     */
	MdApplyScope *string `json:"md_apply_scope,omitempty" `

	/*
	   器械产品描述     */
	MdProductDesc *string `json:"md_product_desc,omitempty" `

	/*
	   代理人名称     */
	AgentEntName *string `json:"agent_ent_name,omitempty" `

	/*
	   代理人地址     */
	AgentEntAddress *string `json:"agent_ent_address,omitempty" `

	/*
	   其他产品小类别，1食品、2化妆品、3消毒产品、4其他     */
	OtherTypeCode *int64 `json:"other_type_code,omitempty" `
}

func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetProductName(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.ProductName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetProductNameEn(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.ProductNameEn = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetApprovalNo(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.ApprovalNo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetApprovalDate(v util.LocalTime) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.ApprovalDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetApprovalEndDate(v util.LocalTime) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.ApprovalEndDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetDrugProductName(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.DrugProductName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetPrepnTypeDesc(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.PrepnTypeDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetPrepnSpec(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.PrepnSpec = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetPkgSpec(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.PkgSpec = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetSpec(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.Spec = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetDrugTypeCode(v int64) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.DrugTypeCode = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetDrugTypeDesc(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.DrugTypeDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetSource(v int64) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.Source = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetProductTypeCode(v int64) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.ProductTypeCode = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetProductTypeDesc(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.ProductTypeDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetProduceEntName(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.ProduceEntName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetMah(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.Mah = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetProduceAddress(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.ProduceAddress = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetSupplyEntName(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.SupplyEntName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetSupplyRefEntId(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.SupplyRefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetSdcCode(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.SdcCode = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetSdcCodeNote(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.SdcCodeNote = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetProductNo(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.ProductNo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetSupplyProductNo(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.SupplyProductNo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetApprovalNoOld(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.ApprovalNoOld = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetExecuteStandard(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.ExecuteStandard = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetPkgApprovalNo(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.PkgApprovalNo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetPkgApprovalDate(v util.LocalTime) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.PkgApprovalDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetPkgApprovalExpireDate(v util.LocalTime) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.PkgApprovalExpireDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetPkgEntName(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.PkgEntName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetPkgEntAddress(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.PkgEntAddress = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetRegisterNo(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.RegisterNo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetMdRegisterDate(v util.LocalTime) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.MdRegisterDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetMdRegisterEndDate(v util.LocalTime) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.MdRegisterEndDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetMdManageTypeCode(v int64) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.MdManageTypeCode = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetMdManageTypeDesc(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.MdManageTypeDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetMdClassifyNo(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.MdClassifyNo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetMdComponent(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.MdComponent = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetMdApplyScope(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.MdApplyScope = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetMdProductDesc(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.MdProductDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetAgentEntName(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.AgentEntName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetAgentEntAddress(v string) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.AgentEntAddress = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) SetOtherTypeCode(v int64) *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO {
	s.OtherTypeCode = &v
	return s
}
