package request


type AlibabaAlihealthDrugYzwDrugtableRequest struct {
    /*
        调用企业id     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        药品通用名     */
    PhysicName  *string `json:"physic_name" required:"true" `
    /*
        批准文号     */
    ApprovalLicenceNo  *string `json:"approval_licence_no" required:"true" `
    /*
        mah企业名称     */
    MahEntName  *string `json:"mah_ent_name,omitempty" required:"false" `
    /*
        生产企业名称     */
    ProductEntName  *string `json:"product_ent_name,omitempty" required:"false" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页码     */
    Page  *int64 `json:"page" required:"true" `
}

func (s *AlibabaAlihealthDrugYzwDrugtableRequest) SetRefEntId(v string) *AlibabaAlihealthDrugYzwDrugtableRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugYzwDrugtableRequest) SetPhysicName(v string) *AlibabaAlihealthDrugYzwDrugtableRequest {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugYzwDrugtableRequest) SetApprovalLicenceNo(v string) *AlibabaAlihealthDrugYzwDrugtableRequest {
    s.ApprovalLicenceNo = &v
    return s
}
func (s *AlibabaAlihealthDrugYzwDrugtableRequest) SetMahEntName(v string) *AlibabaAlihealthDrugYzwDrugtableRequest {
    s.MahEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugYzwDrugtableRequest) SetProductEntName(v string) *AlibabaAlihealthDrugYzwDrugtableRequest {
    s.ProductEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugYzwDrugtableRequest) SetPageSize(v int64) *AlibabaAlihealthDrugYzwDrugtableRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugYzwDrugtableRequest) SetPage(v int64) *AlibabaAlihealthDrugYzwDrugtableRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugYzwDrugtableRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.PhysicName != nil) {
        paramMap["physic_name"] = *req.PhysicName
    }
    if(req.ApprovalLicenceNo != nil) {
        paramMap["approval_licence_no"] = *req.ApprovalLicenceNo
    }
    if(req.MahEntName != nil) {
        paramMap["mah_ent_name"] = *req.MahEntName
    }
    if(req.ProductEntName != nil) {
        paramMap["product_ent_name"] = *req.ProductEntName
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugYzwDrugtableRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}