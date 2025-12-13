package request


type AlibabaAlihealthDrugCodeKytWesCheckcoderelationRequest struct {
    /*
        企业refentid     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.getlicense     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        多个码用英文逗号分隔     */
    Codes  *string `json:"codes" required:"true" `
}

func (s *AlibabaAlihealthDrugCodeKytWesCheckcoderelationRequest) SetRefEntId(v string) *AlibabaAlihealthDrugCodeKytWesCheckcoderelationRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesCheckcoderelationRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugCodeKytWesCheckcoderelationRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesCheckcoderelationRequest) SetCodes(v string) *AlibabaAlihealthDrugCodeKytWesCheckcoderelationRequest {
    s.Codes = &v
    return s
}

func (req *AlibabaAlihealthDrugCodeKytWesCheckcoderelationRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.LicenseToken != nil) {
        paramMap["license_token"] = *req.LicenseToken
    }
    if(req.Codes != nil) {
        paramMap["codes"] = *req.Codes
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugCodeKytWesCheckcoderelationRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}