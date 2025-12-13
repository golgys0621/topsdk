package request


type AlibabaAlihealthDrugMscListbillprocesspartsuccessRequest struct {
    /*
        企业标识     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        单据号     */
    BillCode  *string `json:"bill_code,omitempty" required:"false" `
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

func (s *AlibabaAlihealthDrugMscListbillprocesspartsuccessRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMscListbillprocesspartsuccessRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListbillprocesspartsuccessRequest) SetBillCode(v string) *AlibabaAlihealthDrugMscListbillprocesspartsuccessRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListbillprocesspartsuccessRequest) SetErrorCode(v string) *AlibabaAlihealthDrugMscListbillprocesspartsuccessRequest {
    s.ErrorCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListbillprocesspartsuccessRequest) SetCode(v string) *AlibabaAlihealthDrugMscListbillprocesspartsuccessRequest {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListbillprocesspartsuccessRequest) SetPage(v int64) *AlibabaAlihealthDrugMscListbillprocesspartsuccessRequest {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListbillprocesspartsuccessRequest) SetPageSize(v int64) *AlibabaAlihealthDrugMscListbillprocesspartsuccessRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthDrugMscListbillprocesspartsuccessRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
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

func (req *AlibabaAlihealthDrugMscListbillprocesspartsuccessRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}