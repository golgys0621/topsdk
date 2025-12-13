package request


type AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessRequest struct {
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

func (s *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessRequest) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessRequest) SetErrorCode(v string) *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessRequest {
    s.ErrorCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessRequest) SetCode(v string) *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessRequest {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessRequest) SetPage(v int64) *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessRequest {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessRequest) SetPageSize(v int64) *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessRequest) ToMap() map[string]interface{} {
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

func (req *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}