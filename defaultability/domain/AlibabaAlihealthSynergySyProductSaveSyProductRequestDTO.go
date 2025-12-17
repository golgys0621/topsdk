package domain


type AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO struct {
    /*
        批准文号     */
    ApprovalNo  *string `json:"approval_no,omitempty" `

    /*
        供应商refEntId     */
    SupplyRefEntId  *string `json:"supply_ref_ent_id,omitempty" `

    /*
        执行标准     */
    ExecuteStandard  *string `json:"execute_standard,omitempty" `

    /*
        批准日期     */
    ApprovalDate  *string `json:"approval_date,omitempty" `

    /*
        包装规格     */
    PkgSpec  *string `json:"pkg_spec,omitempty" `

    /*
        批准文号有效期     */
    ApprovalEndDate  *string `json:"approval_end_date,omitempty" `

    /*
        商品名     */
    DrugProductName  *string `json:"drug_product_name,omitempty" `

    /*
        供应商产品编码     */
    SupplyProductNo  *string `json:"supply_product_no,omitempty" `

    /*
        来源，0国产，1进口     */
    Source  *int64 `json:"source,omitempty" `

    /*
        产品名称     */
    ProductName  *string `json:"product_name,omitempty" `

    /*
        规格     */
    Spec  *string `json:"spec,omitempty" `

    /*
        器械注册证号/备案号     */
    RegisterNo  *string `json:"register_no,omitempty" `

    /*
        器械适用范围/预期用途     */
    MdApplyScope  *string `json:"md_apply_scope,omitempty" `

    /*
        上市许可持有人     */
    Mah  *string `json:"mah,omitempty" `

    /*
        器械代理人名称     */
    AgentEntName  *string `json:"agent_ent_name,omitempty" `

    /*
        分包装企业地址     */
    PkgEntAddress  *string `json:"pkg_ent_address,omitempty" `

    /*
        生产地址     */
    ProduceAddress  *string `json:"produce_address,omitempty" `

    /*
        器械分类编码     */
    MdClassifyNo  *string `json:"md_classify_no,omitempty" `

    /*
        剂型     */
    PrepnTypeDesc  *string `json:"prepn_type_desc,omitempty" `

    /*
        更新时传递id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        药品本位码     */
    SdcCode  *string `json:"sdc_code,omitempty" `

    /*
        分包装批准文号     */
    PkgApprovalNo  *string `json:"pkg_approval_no,omitempty" `

    /*
        产品编码     */
    ProductNo  *string `json:"product_no,omitempty" `

    /*
        药品类别,0中药，1中药饮片，2化学药，3生物制品，4辅料     */
    DrugTypeCode  *int64 `json:"drug_type_code,omitempty" `

    /*
        分包装文号有效期截止日     */
    PkgApprovalExpireDate  *string `json:"pkg_approval_expire_date,omitempty" `

    /*
        产品描述     */
    MdProductDesc  *string `json:"md_product_desc,omitempty" `

    /*
        启用停用状态,0停用1启用     */
    OpenStatus  *int64 `json:"open_status,omitempty" `

    /*
        码上放心平台药品id     */
    DrugEntBaseInfoId  *string `json:"drug_ent_base_info_id,omitempty" `

    /*
        药品本位码备注     */
    SdcCodeNote  *string `json:"sdc_code_note,omitempty" `

    /*
        生产企业名称     */
    ProduceEntName  *string `json:"produce_ent_name,omitempty" `

    /*
        供应商企业名称     */
    SupplyEntName  *string `json:"supply_ent_name,omitempty" `

    /*
        分包装企业名称     */
    PkgEntName  *string `json:"pkg_ent_name,omitempty" `

    /*
        器械有效期     */
    MdRegisterEndDate  *string `json:"md_register_end_date,omitempty" `

    /*
        器械结构与组成成分     */
    MdComponent  *string `json:"md_component,omitempty" `

    /*
        产品英文名称     */
    ProductNameEn  *string `json:"product_name_en,omitempty" `

    /*
        分包装文号批准日期     */
    PkgApprovalDate  *string `json:"pkg_approval_date,omitempty" `

    /*
        企业id     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        器械管理类别,1第一类，2第二类，3第三类     */
    MdManageTypeCode  *int64 `json:"md_manage_type_code,omitempty" `

    /*
        器械代理人地址     */
    AgentEntAddress  *string `json:"agent_ent_address,omitempty" `

    /*
        产品类别，0药品1器械2其他     */
    ProductTypeCode  *int64 `json:"product_type_code,omitempty" `

    /*
        原批准文号/注册证号     */
    ApprovalNoOld  *string `json:"approval_no_old,omitempty" `

    /*
        是否可被索取，0否1是     */
    RequestStatus  *int64 `json:"request_status,omitempty" `

    /*
        器械注册日期/备案日期     */
    MdRegisterDate  *string `json:"md_register_date,omitempty" `

    /*
        其他产品小类别，1食品、2化妆品、3消毒产品、4其他，其他产品必传     */
    OtherTypeCode  *int64 `json:"other_type_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetApprovalNo(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.ApprovalNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetSupplyRefEntId(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.SupplyRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetExecuteStandard(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.ExecuteStandard = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetApprovalDate(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.ApprovalDate = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetPkgSpec(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.PkgSpec = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetApprovalEndDate(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.ApprovalEndDate = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetDrugProductName(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.DrugProductName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetSupplyProductNo(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.SupplyProductNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetSource(v int64) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.Source = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetProductName(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.ProductName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetSpec(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.Spec = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetRegisterNo(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.RegisterNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetMdApplyScope(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.MdApplyScope = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetMah(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.Mah = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetAgentEntName(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.AgentEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetPkgEntAddress(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.PkgEntAddress = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetProduceAddress(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.ProduceAddress = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetMdClassifyNo(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.MdClassifyNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetPrepnTypeDesc(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.PrepnTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetId(v int64) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.Id = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetSdcCode(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.SdcCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetPkgApprovalNo(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.PkgApprovalNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetProductNo(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.ProductNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetDrugTypeCode(v int64) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.DrugTypeCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetPkgApprovalExpireDate(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.PkgApprovalExpireDate = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetMdProductDesc(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.MdProductDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetOpenStatus(v int64) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.OpenStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetSdcCodeNote(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.SdcCodeNote = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetProduceEntName(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetSupplyEntName(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.SupplyEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetPkgEntName(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.PkgEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetMdRegisterEndDate(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.MdRegisterEndDate = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetMdComponent(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.MdComponent = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetProductNameEn(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.ProductNameEn = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetPkgApprovalDate(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.PkgApprovalDate = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetRefEntId(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetMdManageTypeCode(v int64) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.MdManageTypeCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetAgentEntAddress(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.AgentEntAddress = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetProductTypeCode(v int64) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.ProductTypeCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetApprovalNoOld(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.ApprovalNoOld = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetRequestStatus(v int64) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.RequestStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetMdRegisterDate(v string) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.MdRegisterDate = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO) SetOtherTypeCode(v int64) *AlibabaAlihealthSynergySyProductSaveSyProductRequestDTO {
    s.OtherTypeCode = &v
    return s
}
