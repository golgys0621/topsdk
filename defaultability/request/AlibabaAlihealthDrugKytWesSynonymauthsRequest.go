package request


type AlibabaAlihealthDrugKytWesSynonymauthsRequest struct {
    /*
        企业ID     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.getlicense     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" required:"false" `
    /*
        货主自定义编号     */
    SynOwnEntId  *string `json:"syn_own_ent_id,omitempty" required:"false" `
    /*
        页码     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页面大小     */
    Page  *int64 `json:"page" required:"true" `
}

func (s *AlibabaAlihealthDrugKytWesSynonymauthsRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesSynonymauthsRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSynonymauthsRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesSynonymauthsRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSynonymauthsRequest) SetEntName(v string) *AlibabaAlihealthDrugKytWesSynonymauthsRequest {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSynonymauthsRequest) SetSynOwnEntId(v string) *AlibabaAlihealthDrugKytWesSynonymauthsRequest {
    s.SynOwnEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSynonymauthsRequest) SetPageSize(v int64) *AlibabaAlihealthDrugKytWesSynonymauthsRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSynonymauthsRequest) SetPage(v int64) *AlibabaAlihealthDrugKytWesSynonymauthsRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugKytWesSynonymauthsRequest) ToMap() map[string]interface{} {
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
    if(req.SynOwnEntId != nil) {
        paramMap["syn_own_ent_id"] = *req.SynOwnEntId
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugKytWesSynonymauthsRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}