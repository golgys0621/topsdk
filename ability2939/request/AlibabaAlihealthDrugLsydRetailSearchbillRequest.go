package request


type AlibabaAlihealthDrugLsydRetailSearchbillRequest struct {
    /*
        企业标识     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        单据所有者     */
    AuthRefUserId  *string `json:"auth_ref_user_id,omitempty" required:"false" `
    /*
        开始日期     */
    BeginDate  *string `json:"begin_date" required:"true" `
    /*
        结束日期     */
    EndDate  *string `json:"end_date" required:"true" `
    /*
        单据号码     */
    BillCode  *string `json:"bill_code,omitempty" required:"false" `
    /*
        单据类型  A : 所有     */
    BillType  *string `json:"bill_type" required:"true" `
    /*
        当前页     */
    CurPage  *int64 `json:"cur_page" required:"true" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *AlibabaAlihealthDrugLsydRetailSearchbillRequest) SetRefEntId(v string) *AlibabaAlihealthDrugLsydRetailSearchbillRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillRequest) SetAuthRefUserId(v string) *AlibabaAlihealthDrugLsydRetailSearchbillRequest {
    s.AuthRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillRequest) SetBeginDate(v string) *AlibabaAlihealthDrugLsydRetailSearchbillRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillRequest) SetEndDate(v string) *AlibabaAlihealthDrugLsydRetailSearchbillRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillRequest) SetBillCode(v string) *AlibabaAlihealthDrugLsydRetailSearchbillRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillRequest) SetBillType(v string) *AlibabaAlihealthDrugLsydRetailSearchbillRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillRequest) SetCurPage(v int64) *AlibabaAlihealthDrugLsydRetailSearchbillRequest {
    s.CurPage = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillRequest) SetPageSize(v int64) *AlibabaAlihealthDrugLsydRetailSearchbillRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthDrugLsydRetailSearchbillRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
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
    return paramMap
}

func (req *AlibabaAlihealthDrugLsydRetailSearchbillRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}