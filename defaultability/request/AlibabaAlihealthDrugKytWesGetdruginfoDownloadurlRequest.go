package request


type AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlRequest struct {
    /*
        调用接口的企业ID     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        服务时间     */
    LicenseToken  *string `json:"license_token" required:"true" `
}

func (s *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlRequest {
    s.LicenseToken = &v
    return s
}

func (req *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.LicenseToken != nil) {
        paramMap["license_token"] = *req.LicenseToken
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}