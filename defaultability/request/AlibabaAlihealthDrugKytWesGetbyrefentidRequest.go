package request


type AlibabaAlihealthDrugKytWesGetbyrefentidRequest struct {
    /*
        接口调用企业的唯一标识（接口调用者）     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.getlicense     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）     */
    DestRefEntId  *string `json:"dest_ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthDrugKytWesGetbyrefentidRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidRequest) SetDestRefEntId(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidRequest {
    s.DestRefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugKytWesGetbyrefentidRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.LicenseToken != nil) {
        paramMap["license_token"] = *req.LicenseToken
    }
    if(req.DestRefEntId != nil) {
        paramMap["dest_ref_ent_id"] = *req.DestRefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugKytWesGetbyrefentidRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}