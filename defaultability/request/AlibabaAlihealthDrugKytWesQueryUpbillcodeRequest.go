package request


type AlibabaAlihealthDrugKytWesQueryUpbillcodeRequest struct {
    /*
        企业REF_ENT_ID （当前企业的唯一标识）     */
    RefEntId  *string `json:"ref_ent_id,omitempty" required:"false" `
    /*
        获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.getlicense     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        追溯码     */
    Code  *string `json:"code,omitempty" required:"false" `
    /*
        物流委托单位     */
    AgentRefEntId  *string `json:"agent_ref_ent_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesQueryUpbillcodeRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesQueryUpbillcodeRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeRequest) SetCode(v string) *AlibabaAlihealthDrugKytWesQueryUpbillcodeRequest {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeRequest) SetAgentRefEntId(v string) *AlibabaAlihealthDrugKytWesQueryUpbillcodeRequest {
    s.AgentRefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugKytWesQueryUpbillcodeRequest) ToMap() map[string]interface{} {
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
    if(req.AgentRefEntId != nil) {
        paramMap["agent_ref_ent_id"] = *req.AgentRefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugKytWesQueryUpbillcodeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}