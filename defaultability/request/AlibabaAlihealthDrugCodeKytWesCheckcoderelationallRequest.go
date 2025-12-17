package request


type AlibabaAlihealthDrugCodeKytWesCheckcoderelationallRequest struct {
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

func (s *AlibabaAlihealthDrugCodeKytWesCheckcoderelationallRequest) SetRefEntId(v string) *AlibabaAlihealthDrugCodeKytWesCheckcoderelationallRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesCheckcoderelationallRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugCodeKytWesCheckcoderelationallRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesCheckcoderelationallRequest) SetCodes(v string) *AlibabaAlihealthDrugCodeKytWesCheckcoderelationallRequest {
    s.Codes = &v
    return s
}

func (req *AlibabaAlihealthDrugCodeKytWesCheckcoderelationallRequest) ToMap() map[string]interface{} {
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

func (req *AlibabaAlihealthDrugCodeKytWesCheckcoderelationallRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}