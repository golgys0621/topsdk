package request


type AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoRequest struct {
    /*
        企业标识     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.getlicense     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        码     */
    Code  *string `json:"code" required:"true" `
}

func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoRequest) SetRefEntId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoRequest) SetCode(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoRequest {
    s.Code = &v
    return s
}

func (req *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.LicenseToken != nil) {
        paramMap["license_token"] = *req.LicenseToken
    }
    if(req.Code != nil) {
        paramMap["code"] = *req.Code
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}