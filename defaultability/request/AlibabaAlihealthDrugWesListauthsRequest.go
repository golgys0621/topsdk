package request


type AlibabaAlihealthDrugWesListauthsRequest struct {
    /*
        企业id     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.getlicense     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" required:"false" `
    /*
        页码     */
    Page  *int64 `json:"page" required:"true" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *AlibabaAlihealthDrugWesListauthsRequest) SetRefEntId(v string) *AlibabaAlihealthDrugWesListauthsRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugWesListauthsRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugWesListauthsRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugWesListauthsRequest) SetEntName(v string) *AlibabaAlihealthDrugWesListauthsRequest {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugWesListauthsRequest) SetPage(v int64) *AlibabaAlihealthDrugWesListauthsRequest {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthDrugWesListauthsRequest) SetPageSize(v int64) *AlibabaAlihealthDrugWesListauthsRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthDrugWesListauthsRequest) ToMap() map[string]interface{} {
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
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugWesListauthsRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}