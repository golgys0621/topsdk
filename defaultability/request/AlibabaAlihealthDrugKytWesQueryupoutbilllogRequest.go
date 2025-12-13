package request


type AlibabaAlihealthDrugKytWesQueryupoutbilllogRequest struct {
    /*
        企业id     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        获取的临时licenseToken     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        开始查询时间（日期格式）     */
    BeginOperDate  *string `json:"begin_oper_date" required:"true" `
    /*
        截止查询时间（日期格式）     */
    EndOperDate  *string `json:"end_oper_date" required:"true" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页码     */
    Page  *int64 `json:"page" required:"true" `
}

func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesQueryupoutbilllogRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesQueryupoutbilllogRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogRequest) SetBeginOperDate(v string) *AlibabaAlihealthDrugKytWesQueryupoutbilllogRequest {
    s.BeginOperDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogRequest) SetEndOperDate(v string) *AlibabaAlihealthDrugKytWesQueryupoutbilllogRequest {
    s.EndOperDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogRequest) SetPageSize(v int64) *AlibabaAlihealthDrugKytWesQueryupoutbilllogRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogRequest) SetPage(v int64) *AlibabaAlihealthDrugKytWesQueryupoutbilllogRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugKytWesQueryupoutbilllogRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.LicenseToken != nil) {
        paramMap["license_token"] = *req.LicenseToken
    }
    if(req.BeginOperDate != nil) {
        paramMap["begin_oper_date"] = *req.BeginOperDate
    }
    if(req.EndOperDate != nil) {
        paramMap["end_oper_date"] = *req.EndOperDate
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugKytWesQueryupoutbilllogRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}