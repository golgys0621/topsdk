package request


type AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoRequest struct {
    /*
        企业标识     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.getlicense     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        码,如有多个英文逗号隔开。一次最多支持10个追溯码     */
    Code  *string `json:"code" required:"true" `
    /*
        货主企业（在三方物流场景中，当物流企业查询货主是否上传过单据时，auth_ref_user_id 对应货主企业，ref_ent_id 对应当前开通 WES 权限的物流企业。）     */
    AuthRefUserId  *string `json:"auth_ref_user_id,omitempty" required:"false" `
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
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoRequest) SetAuthRefUserId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoRequest {
    s.AuthRefUserId = &v
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
    if(req.AuthRefUserId != nil) {
        paramMap["auth_ref_user_id"] = *req.AuthRefUserId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}