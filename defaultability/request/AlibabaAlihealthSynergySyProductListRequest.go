package request


type AlibabaAlihealthSynergySyProductListRequest struct {
    /*
        批准文号     */
    ApprovalNo  *string `json:"approval_no,omitempty" required:"false" `
    /*
        器械管理类别     */
    MdManageType  *string `json:"md_manage_type,omitempty" required:"false" `
    /*
        启用状态，0停用1启用     */
    OpenStatus  *int64 `json:"open_status,omitempty" required:"false" `
    /*
        包装规格     */
    PkgSpec  *string `json:"pkg_spec,omitempty" required:"false" `
    /*
        创建方式，1手动创建2引入生成3在线接收     */
    CreateType  *string `json:"create_type,omitempty" required:"false" `
    /*
        来源，0国产1进口     */
    Source  *string `json:"source,omitempty" required:"false" `
    /*
        产品名称     */
    ProductName  *string `json:"product_name,omitempty" required:"false" `
    /*
        规格     */
    Spec  *string `json:"spec,omitempty" required:"false" `
    /*
        器械注册证号     */
    RegisterNo  *string `json:"register_no,omitempty" required:"false" `
    /*
        生产企业名称     */
    ProduceEntName  *string `json:"produce_ent_name,omitempty" required:"false" `
    /*
        供应商名称     */
    SupplyEntName  *string `json:"supply_ent_name,omitempty" required:"false" `
    /*
        归档状态，0未归档1已归档     */
    ArchiveStatus  *string `json:"archive_status,omitempty" required:"false" `
    /*
        剂型     */
    PrepnTypeDesc  *string `json:"prepn_type_desc,omitempty" required:"false" `
    /*
        是否有过期资料，0否1是     */
    ExpireStatus  *string `json:"expire_status,omitempty" required:"false" `
    /*
        企业id     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        类别，0药品1器械2其他     */
    ProductTypeCode  *int64 `json:"product_type_code" required:"true" `
    /*
        产品编码     */
    ProductNo  *string `json:"product_no,omitempty" required:"false" `
    /*
        药品类别，0中药1中药饮片2化学药3生物制品4辅料     */
    DrugTypeCode  *string `json:"drug_type_code,omitempty" required:"false" `
    /*
        其他产品小类别，1食品、2化妆品、3消毒产品、4其他     */
    OtherTypeCode  *int64 `json:"other_type_code,omitempty" required:"false" `
    /*
        页数，1开始 defalutValue��1    */
    Page  *int64 `json:"page" required:"true" `
    /*
        页码 defalutValue��10    */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *AlibabaAlihealthSynergySyProductListRequest) SetApprovalNo(v string) *AlibabaAlihealthSynergySyProductListRequest {
    s.ApprovalNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetMdManageType(v string) *AlibabaAlihealthSynergySyProductListRequest {
    s.MdManageType = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetOpenStatus(v int64) *AlibabaAlihealthSynergySyProductListRequest {
    s.OpenStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetPkgSpec(v string) *AlibabaAlihealthSynergySyProductListRequest {
    s.PkgSpec = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetCreateType(v string) *AlibabaAlihealthSynergySyProductListRequest {
    s.CreateType = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetSource(v string) *AlibabaAlihealthSynergySyProductListRequest {
    s.Source = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetProductName(v string) *AlibabaAlihealthSynergySyProductListRequest {
    s.ProductName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetSpec(v string) *AlibabaAlihealthSynergySyProductListRequest {
    s.Spec = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetRegisterNo(v string) *AlibabaAlihealthSynergySyProductListRequest {
    s.RegisterNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetProduceEntName(v string) *AlibabaAlihealthSynergySyProductListRequest {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetSupplyEntName(v string) *AlibabaAlihealthSynergySyProductListRequest {
    s.SupplyEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetArchiveStatus(v string) *AlibabaAlihealthSynergySyProductListRequest {
    s.ArchiveStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetPrepnTypeDesc(v string) *AlibabaAlihealthSynergySyProductListRequest {
    s.PrepnTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetExpireStatus(v string) *AlibabaAlihealthSynergySyProductListRequest {
    s.ExpireStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyProductListRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetProductTypeCode(v int64) *AlibabaAlihealthSynergySyProductListRequest {
    s.ProductTypeCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetProductNo(v string) *AlibabaAlihealthSynergySyProductListRequest {
    s.ProductNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetDrugTypeCode(v string) *AlibabaAlihealthSynergySyProductListRequest {
    s.DrugTypeCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetOtherTypeCode(v int64) *AlibabaAlihealthSynergySyProductListRequest {
    s.OtherTypeCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetPage(v int64) *AlibabaAlihealthSynergySyProductListRequest {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListRequest) SetPageSize(v int64) *AlibabaAlihealthSynergySyProductListRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthSynergySyProductListRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ApprovalNo != nil) {
        paramMap["approval_no"] = *req.ApprovalNo
    }
    if(req.MdManageType != nil) {
        paramMap["md_manage_type"] = *req.MdManageType
    }
    if(req.OpenStatus != nil) {
        paramMap["open_status"] = *req.OpenStatus
    }
    if(req.PkgSpec != nil) {
        paramMap["pkg_spec"] = *req.PkgSpec
    }
    if(req.CreateType != nil) {
        paramMap["create_type"] = *req.CreateType
    }
    if(req.Source != nil) {
        paramMap["source"] = *req.Source
    }
    if(req.ProductName != nil) {
        paramMap["product_name"] = *req.ProductName
    }
    if(req.Spec != nil) {
        paramMap["spec"] = *req.Spec
    }
    if(req.RegisterNo != nil) {
        paramMap["register_no"] = *req.RegisterNo
    }
    if(req.ProduceEntName != nil) {
        paramMap["produce_ent_name"] = *req.ProduceEntName
    }
    if(req.SupplyEntName != nil) {
        paramMap["supply_ent_name"] = *req.SupplyEntName
    }
    if(req.ArchiveStatus != nil) {
        paramMap["archive_status"] = *req.ArchiveStatus
    }
    if(req.PrepnTypeDesc != nil) {
        paramMap["prepn_type_desc"] = *req.PrepnTypeDesc
    }
    if(req.ExpireStatus != nil) {
        paramMap["expire_status"] = *req.ExpireStatus
    }
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.ProductTypeCode != nil) {
        paramMap["product_type_code"] = *req.ProductTypeCode
    }
    if(req.ProductNo != nil) {
        paramMap["product_no"] = *req.ProductNo
    }
    if(req.DrugTypeCode != nil) {
        paramMap["drug_type_code"] = *req.DrugTypeCode
    }
    if(req.OtherTypeCode != nil) {
        paramMap["other_type_code"] = *req.OtherTypeCode
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyProductListRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}