package request


type AlibabaAlihealthDrugKytWesListbillprocesspartsuccessRequest struct {
    /*
        企业标识     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        licenseToken     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        单据号     */
    BillCode  *string `json:"bill_code" required:"true" `
    /*
        错误码类型     */
    ErrorCode  *string `json:"error_code,omitempty" required:"false" `
    /*
        错误码     */
    Code  *string `json:"code,omitempty" required:"false" `
    /*
        当前页     */
    Page  *int64 `json:"page" required:"true" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessRequest) SetBillCode(v string) *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessRequest) SetErrorCode(v string) *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessRequest {
    s.ErrorCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessRequest) SetCode(v string) *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessRequest {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessRequest) SetPage(v int64) *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessRequest {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessRequest) SetPageSize(v int64) *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.LicenseToken != nil) {
        paramMap["license_token"] = *req.LicenseToken
    }
    if(req.BillCode != nil) {
        paramMap["bill_code"] = *req.BillCode
    }
    if(req.ErrorCode != nil) {
        paramMap["error_code"] = *req.ErrorCode
    }
    if(req.Code != nil) {
        paramMap["code"] = *req.Code
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}