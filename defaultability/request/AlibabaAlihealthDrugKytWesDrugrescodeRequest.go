package request


type AlibabaAlihealthDrugKytWesDrugrescodeRequest struct {
    /*
        企业ID     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.getlicense     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        药品通用名     */
    PhysicName  *string `json:"physic_name" required:"true" `
    /*
        批准文号     */
    ApprovalLicenceNo  *string `json:"approval_licence_no,omitempty" required:"false" `
    /*
        开始日期     */
    StartDate  *string `json:"start_date,omitempty" required:"false" `
    /*
        结束日期     */
    EndDate  *string `json:"end_date,omitempty" required:"false" `
    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" required:"false" `
    /*
        包装规格     */
    PackageSpec  *string `json:"package_spec,omitempty" required:"false" `
    /*
        制剂规格     */
    PrepnSpec  *string `json:"prepn_spec,omitempty" required:"false" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页码     */
    Page  *int64 `json:"page" required:"true" `
}

func (s *AlibabaAlihealthDrugKytWesDrugrescodeRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesDrugrescodeRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesDrugrescodeRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeRequest) SetPhysicName(v string) *AlibabaAlihealthDrugKytWesDrugrescodeRequest {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeRequest) SetApprovalLicenceNo(v string) *AlibabaAlihealthDrugKytWesDrugrescodeRequest {
    s.ApprovalLicenceNo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeRequest) SetStartDate(v string) *AlibabaAlihealthDrugKytWesDrugrescodeRequest {
    s.StartDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeRequest) SetEndDate(v string) *AlibabaAlihealthDrugKytWesDrugrescodeRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeRequest) SetEntName(v string) *AlibabaAlihealthDrugKytWesDrugrescodeRequest {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeRequest) SetPackageSpec(v string) *AlibabaAlihealthDrugKytWesDrugrescodeRequest {
    s.PackageSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeRequest) SetPrepnSpec(v string) *AlibabaAlihealthDrugKytWesDrugrescodeRequest {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeRequest) SetPageSize(v int64) *AlibabaAlihealthDrugKytWesDrugrescodeRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeRequest) SetPage(v int64) *AlibabaAlihealthDrugKytWesDrugrescodeRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugKytWesDrugrescodeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.LicenseToken != nil) {
        paramMap["license_token"] = *req.LicenseToken
    }
    if(req.PhysicName != nil) {
        paramMap["physic_name"] = *req.PhysicName
    }
    if(req.ApprovalLicenceNo != nil) {
        paramMap["approval_licence_no"] = *req.ApprovalLicenceNo
    }
    if(req.StartDate != nil) {
        paramMap["start_date"] = *req.StartDate
    }
    if(req.EndDate != nil) {
        paramMap["end_date"] = *req.EndDate
    }
    if(req.EntName != nil) {
        paramMap["ent_name"] = *req.EntName
    }
    if(req.PackageSpec != nil) {
        paramMap["package_spec"] = *req.PackageSpec
    }
    if(req.PrepnSpec != nil) {
        paramMap["prepn_spec"] = *req.PrepnSpec
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugKytWesDrugrescodeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}