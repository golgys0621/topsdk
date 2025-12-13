package request


type AlibabaAlihealthDrugKytWesServiceInfoRequest struct {
    /*
        企业的refEntId     */
    RefEntId  *string `json:"ref_ent_id,omitempty" required:"false" `
    /*
        获取的临时licenseToken     */
    LicenseToken  *string `json:"license_token,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugKytWesServiceInfoRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesServiceInfoRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesServiceInfoRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesServiceInfoRequest {
    s.LicenseToken = &v
    return s
}

func (req *AlibabaAlihealthDrugKytWesServiceInfoRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.LicenseToken != nil) {
        paramMap["license_token"] = *req.LicenseToken
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugKytWesServiceInfoRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}