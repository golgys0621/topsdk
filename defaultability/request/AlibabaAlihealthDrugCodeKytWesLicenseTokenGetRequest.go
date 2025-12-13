package request


type AlibabaAlihealthDrugCodeKytWesLicenseTokenGetRequest struct {
    /*
        企业id     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        license     */
    License  *string `json:"license" required:"true" `
}

func (s *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetRequest) SetRefEntId(v string) *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetRequest) SetLicense(v string) *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetRequest {
    s.License = &v
    return s
}

func (req *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.License != nil) {
        paramMap["license"] = *req.License
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}