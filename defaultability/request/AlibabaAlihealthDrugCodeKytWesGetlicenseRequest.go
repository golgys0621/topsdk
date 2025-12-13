package request


type AlibabaAlihealthDrugCodeKytWesGetlicenseRequest struct {
    /*
        企业id     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        license     */
    License  *string `json:"license" required:"true" `
}

func (s *AlibabaAlihealthDrugCodeKytWesGetlicenseRequest) SetRefEntId(v string) *AlibabaAlihealthDrugCodeKytWesGetlicenseRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesGetlicenseRequest) SetLicense(v string) *AlibabaAlihealthDrugCodeKytWesGetlicenseRequest {
    s.License = &v
    return s
}

func (req *AlibabaAlihealthDrugCodeKytWesGetlicenseRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.License != nil) {
        paramMap["license"] = *req.License
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugCodeKytWesGetlicenseRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}