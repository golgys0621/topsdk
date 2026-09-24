package request


type AlibabaAlihealthDrugKytWesRetailSearchbillRequest struct {
    /*
        企业标识     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.getlicense     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        单据所有者     */
    AuthRefUserId  *string `json:"auth_ref_user_id" required:"true" `
    /*
        单据时间的开始日期（不写时分秒），格式：yyyy-MM-dd     */
    BeginDate  *string `json:"begin_date,omitempty" required:"false" `
    /*
        单据时间的结束日期（不写时分秒），格式：yyyy-MM-dd     */
    EndDate  *string `json:"end_date,omitempty" required:"false" `
    /*
        代理企业     */
    AgentRefUserId  *string `json:"agent_ref_user_id,omitempty" required:"false" `
    /*
        单据号码     */
    BillCode  *string `json:"bill_code,omitempty" required:"false" `
    /*
        单据类型  A : 所有      */
    BillType  *string `json:"bill_type" required:"true" `
    /*
        当前页     */
    CurPage  *int64 `json:"cur_page" required:"true" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        单据上传的开始日期，格式：yyyy-MM-dd HH:mm:ss     */
    UploadTimeBegin  *string `json:"upload_time_begin,omitempty" required:"false" `
    /*
        单据上传的结束日期，格式：yyyy-MM-dd HH:mm:ss     */
    UploadTimeEnd  *string `json:"upload_time_end,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugKytWesRetailSearchbillRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillRequest) SetAuthRefUserId(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillRequest {
    s.AuthRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillRequest) SetBeginDate(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillRequest) SetEndDate(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillRequest) SetAgentRefUserId(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillRequest {
    s.AgentRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillRequest) SetBillCode(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillRequest) SetBillType(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillRequest) SetCurPage(v int64) *AlibabaAlihealthDrugKytWesRetailSearchbillRequest {
    s.CurPage = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillRequest) SetPageSize(v int64) *AlibabaAlihealthDrugKytWesRetailSearchbillRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillRequest) SetUploadTimeBegin(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillRequest {
    s.UploadTimeBegin = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillRequest) SetUploadTimeEnd(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillRequest {
    s.UploadTimeEnd = &v
    return s
}

func (req *AlibabaAlihealthDrugKytWesRetailSearchbillRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.LicenseToken != nil) {
        paramMap["license_token"] = *req.LicenseToken
    }
    if(req.AuthRefUserId != nil) {
        paramMap["auth_ref_user_id"] = *req.AuthRefUserId
    }
    if(req.BeginDate != nil) {
        paramMap["begin_date"] = *req.BeginDate
    }
    if(req.EndDate != nil) {
        paramMap["end_date"] = *req.EndDate
    }
    if(req.AgentRefUserId != nil) {
        paramMap["agent_ref_user_id"] = *req.AgentRefUserId
    }
    if(req.BillCode != nil) {
        paramMap["bill_code"] = *req.BillCode
    }
    if(req.BillType != nil) {
        paramMap["bill_type"] = *req.BillType
    }
    if(req.CurPage != nil) {
        paramMap["cur_page"] = *req.CurPage
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.UploadTimeBegin != nil) {
        paramMap["upload_time_begin"] = *req.UploadTimeBegin
    }
    if(req.UploadTimeEnd != nil) {
        paramMap["upload_time_end"] = *req.UploadTimeEnd
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugKytWesRetailSearchbillRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}