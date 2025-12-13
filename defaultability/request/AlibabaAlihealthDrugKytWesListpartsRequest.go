package request


type AlibabaAlihealthDrugKytWesListpartsRequest struct {
    /*
        企业ID：查谁的往来单位数据就传谁     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.getlicense     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" required:"false" `
    /*
        企业自定义编号     */
    RefPartnerId  *string `json:"ref_partner_id,omitempty" required:"false" `
    /*
        开始时间：往来单位最后修改时间（不推荐使用：因为往来单位是共用的，任意企业提交了信息变更都会引起这个值的变更）     */
    BeginDate  *string `json:"begin_date,omitempty" required:"false" `
    /*
        结束时间：往来单位最后修改时间（不推荐使用：因为往来单位是共用的，任意企业提交了信息变更都会引起这个值的变更）     */
    EndDate  *string `json:"end_date,omitempty" required:"false" `
    /*
        1审核通过、2审核不通过     */
    AuditFlag  *int64 `json:"audit_flag,omitempty" required:"false" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页码     */
    Page  *int64 `json:"page" required:"true" `
}

func (s *AlibabaAlihealthDrugKytWesListpartsRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesListpartsRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesListpartsRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsRequest) SetEntName(v string) *AlibabaAlihealthDrugKytWesListpartsRequest {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsRequest) SetRefPartnerId(v string) *AlibabaAlihealthDrugKytWesListpartsRequest {
    s.RefPartnerId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsRequest) SetBeginDate(v string) *AlibabaAlihealthDrugKytWesListpartsRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsRequest) SetEndDate(v string) *AlibabaAlihealthDrugKytWesListpartsRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsRequest) SetAuditFlag(v int64) *AlibabaAlihealthDrugKytWesListpartsRequest {
    s.AuditFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsRequest) SetPageSize(v int64) *AlibabaAlihealthDrugKytWesListpartsRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsRequest) SetPage(v int64) *AlibabaAlihealthDrugKytWesListpartsRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugKytWesListpartsRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.LicenseToken != nil) {
        paramMap["license_token"] = *req.LicenseToken
    }
    if(req.EntName != nil) {
        paramMap["ent_name"] = *req.EntName
    }
    if(req.RefPartnerId != nil) {
        paramMap["ref_partner_id"] = *req.RefPartnerId
    }
    if(req.BeginDate != nil) {
        paramMap["begin_date"] = *req.BeginDate
    }
    if(req.EndDate != nil) {
        paramMap["end_date"] = *req.EndDate
    }
    if(req.AuditFlag != nil) {
        paramMap["audit_flag"] = *req.AuditFlag
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugKytWesListpartsRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}