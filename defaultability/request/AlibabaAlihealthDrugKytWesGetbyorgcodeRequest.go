package request


type AlibabaAlihealthDrugKytWesGetbyorgcodeRequest struct {
    /*
        接口调用企业的唯一标识（接口调用者）     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        licenseToken     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        统一信用代码     */
    OrgCode  *string `json:"org_code" required:"true" `
}

func (s *AlibabaAlihealthDrugKytWesGetbyorgcodeRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesGetbyorgcodeRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyorgcodeRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesGetbyorgcodeRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyorgcodeRequest) SetOrgCode(v string) *AlibabaAlihealthDrugKytWesGetbyorgcodeRequest {
    s.OrgCode = &v
    return s
}

func (req *AlibabaAlihealthDrugKytWesGetbyorgcodeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.LicenseToken != nil) {
        paramMap["license_token"] = *req.LicenseToken
    }
    if(req.OrgCode != nil) {
        paramMap["org_code"] = *req.OrgCode
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugKytWesGetbyorgcodeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}