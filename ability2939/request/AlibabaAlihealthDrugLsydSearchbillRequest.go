package request


type AlibabaAlihealthDrugLsydSearchbillRequest struct {
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
        发货企业entId     */
    PartnerIdSend  *string `json:"partner_id_send,omitempty" required:"false" `
    /*
        收货企业entId     */
    PartnerIdRecv  *string `json:"partner_id_recv,omitempty" required:"false" `
    /*
        单据号码     */
    BillCode  *string `json:"bill_code,omitempty" required:"false" `
    /*
        单据类型  A : 所有  AI :入库    AO:出库     */
    BillType  *string `json:"bill_type" required:"true" `
    /*
        当前页     */
    CurPage  *int64 `json:"cur_page" required:"true" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *AlibabaAlihealthDrugLsydSearchbillRequest) SetRefEntId(v string) *AlibabaAlihealthDrugLsydSearchbillRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillRequest) SetAuthRefUserId(v string) *AlibabaAlihealthDrugLsydSearchbillRequest {
    s.AuthRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillRequest) SetBeginDate(v string) *AlibabaAlihealthDrugLsydSearchbillRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillRequest) SetEndDate(v string) *AlibabaAlihealthDrugLsydSearchbillRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillRequest) SetPartnerIdSend(v string) *AlibabaAlihealthDrugLsydSearchbillRequest {
    s.PartnerIdSend = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillRequest) SetPartnerIdRecv(v string) *AlibabaAlihealthDrugLsydSearchbillRequest {
    s.PartnerIdRecv = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillRequest) SetBillCode(v string) *AlibabaAlihealthDrugLsydSearchbillRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillRequest) SetBillType(v string) *AlibabaAlihealthDrugLsydSearchbillRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillRequest) SetCurPage(v int64) *AlibabaAlihealthDrugLsydSearchbillRequest {
    s.CurPage = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillRequest) SetPageSize(v int64) *AlibabaAlihealthDrugLsydSearchbillRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthDrugLsydSearchbillRequest) ToMap() map[string]interface{} {
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
    if(req.PartnerIdSend != nil) {
        paramMap["partner_id_send"] = *req.PartnerIdSend
    }
    if(req.PartnerIdRecv != nil) {
        paramMap["partner_id_recv"] = *req.PartnerIdRecv
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

func (req *AlibabaAlihealthDrugLsydSearchbillRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}