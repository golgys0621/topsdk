package request


type AlibabaAlihealthDrugKytWesGetbyentidRequest struct {
    /*
        接口调用企业的唯一标识（接口调用者）     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.getlicense     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        准备要查询的企业ID（返回该企业ID的详细信息）     */
    EntId  *string `json:"ent_id" required:"true" `
}

func (s *AlibabaAlihealthDrugKytWesGetbyentidRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesGetbyentidRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesGetbyentidRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidRequest) SetEntId(v string) *AlibabaAlihealthDrugKytWesGetbyentidRequest {
    s.EntId = &v
    return s
}

func (req *AlibabaAlihealthDrugKytWesGetbyentidRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.LicenseToken != nil) {
        paramMap["license_token"] = *req.LicenseToken
    }
    if(req.EntId != nil) {
        paramMap["ent_id"] = *req.EntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugKytWesGetbyentidRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}