package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyProductListSyProductDTO struct {
	/*
	   产品id     */
	Id *int64 `json:"id,omitempty" `

	/*
	   创建时间     */
	GmtCreate *util.LocalTime `json:"gmt_create,omitempty" `

	/*
	   更新时间     */
	GmtModified *util.LocalTime `json:"gmt_modified,omitempty" `

	/*
	   产品名称     */
	ProductName *string `json:"product_name,omitempty" `

	/*
	   英文名称     */
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
	   药品分类     */
	DrugTypeCode *int64 `json:"drug_type_code,omitempty" `

	/*
	   药品分类     */
	DrugTypeDesc *string `json:"drug_type_desc,omitempty" `

	/*
	   来源     */
	Source *int64 `json:"source,omitempty" `

	/*
	   产品分类     */
	ProductTypeCode *int64 `json:"product_type_code,omitempty" `

	/*
	   产品分类     */
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
	   原批准文号/注册证号     */
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
	   器械管理分类     */
	MdManageTypeCode *int64 `json:"md_manage_type_code,omitempty" `

	/*
	   器械管理分类     */
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
	   是否有过期资料     */
	ExpireStatus *int64 `json:"expire_status,omitempty" `

	/*
	   是否有归档资料     */
	ArchiveStatus *int64 `json:"archive_status,omitempty" `

	/*
	   启用状态     */
	OpenStatus *int64 `json:"open_status,omitempty" `

	/*
	   生成方式     */
	CreateType *int64 `json:"create_type,omitempty" `

	/*
	   是否可被索取     */
	RequestStatus *int64 `json:"request_status,omitempty" `

	/*
	   其他产品小类别，1食品、2化妆品、3消毒产品、4其他     */
	OtherTypeCode *int64 `json:"other_type_code,omitempty" `
}

func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetId(v int64) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.Id = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetGmtCreate(v util.LocalTime) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.GmtCreate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetGmtModified(v util.LocalTime) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.GmtModified = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetProductName(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.ProductName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetProductNameEn(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.ProductNameEn = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetApprovalNo(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.ApprovalNo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetApprovalDate(v util.LocalTime) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.ApprovalDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetApprovalEndDate(v util.LocalTime) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.ApprovalEndDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetDrugProductName(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.DrugProductName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetPrepnTypeDesc(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.PrepnTypeDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetPrepnSpec(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.PrepnSpec = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetPkgSpec(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.PkgSpec = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetSpec(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.Spec = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetDrugTypeCode(v int64) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.DrugTypeCode = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetDrugTypeDesc(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.DrugTypeDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetSource(v int64) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.Source = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetProductTypeCode(v int64) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.ProductTypeCode = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetProductTypeDesc(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.ProductTypeDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetProduceEntName(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.ProduceEntName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetMah(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.Mah = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetProduceAddress(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.ProduceAddress = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetSupplyEntName(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.SupplyEntName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetSupplyRefEntId(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.SupplyRefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetSdcCode(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.SdcCode = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetSdcCodeNote(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.SdcCodeNote = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetProductNo(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.ProductNo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetSupplyProductNo(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.SupplyProductNo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetApprovalNoOld(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.ApprovalNoOld = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetExecuteStandard(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.ExecuteStandard = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetPkgApprovalNo(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.PkgApprovalNo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetPkgApprovalDate(v util.LocalTime) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.PkgApprovalDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetPkgApprovalExpireDate(v util.LocalTime) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.PkgApprovalExpireDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetPkgEntName(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.PkgEntName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetPkgEntAddress(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.PkgEntAddress = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetRegisterNo(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.RegisterNo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetMdRegisterDate(v util.LocalTime) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.MdRegisterDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetMdRegisterEndDate(v util.LocalTime) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.MdRegisterEndDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetMdManageTypeCode(v int64) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.MdManageTypeCode = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetMdManageTypeDesc(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.MdManageTypeDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetMdClassifyNo(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.MdClassifyNo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetMdComponent(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.MdComponent = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetMdApplyScope(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.MdApplyScope = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetMdProductDesc(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.MdProductDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetAgentEntName(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.AgentEntName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetAgentEntAddress(v string) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.AgentEntAddress = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetExpireStatus(v int64) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.ExpireStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetArchiveStatus(v int64) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.ArchiveStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetOpenStatus(v int64) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.OpenStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetCreateType(v int64) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.CreateType = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetRequestStatus(v int64) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.RequestStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductListSyProductDTO) SetOtherTypeCode(v int64) *AlibabaAlihealthSynergySyProductListSyProductDTO {
	s.OtherTypeCode = &v
	return s
}
