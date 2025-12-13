package request


type AlibabaAlihealthDrugCodeKytWesQuerycoderelationRequest struct {
    /*
        接口调用企业的唯一标识（接口调用者）     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        license_token     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        追溯码     */
    Code  *string `json:"code" required:"true" `
    /*
        目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）     */
    DesRefEntId  *string `json:"des_ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationRequest) SetRefEntId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationRequest) SetCode(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationRequest {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationRequest) SetDesRefEntId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationRequest {
    s.DesRefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugCodeKytWesQuerycoderelationRequest) ToMap() map[string]interface{} {
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
    if(req.DesRefEntId != nil) {
        paramMap["des_ref_ent_id"] = *req.DesRefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugCodeKytWesQuerycoderelationRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}