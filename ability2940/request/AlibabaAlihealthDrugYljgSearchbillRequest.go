package request


type AlibabaAlihealthDrugYljgSearchbillRequest struct {
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

func (s *AlibabaAlihealthDrugYljgSearchbillRequest) SetRefEntId(v string) *AlibabaAlihealthDrugYljgSearchbillRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillRequest) SetAuthRefUserId(v string) *AlibabaAlihealthDrugYljgSearchbillRequest {
    s.AuthRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillRequest) SetBeginDate(v string) *AlibabaAlihealthDrugYljgSearchbillRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillRequest) SetEndDate(v string) *AlibabaAlihealthDrugYljgSearchbillRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillRequest) SetPartnerIdSend(v string) *AlibabaAlihealthDrugYljgSearchbillRequest {
    s.PartnerIdSend = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillRequest) SetPartnerIdRecv(v string) *AlibabaAlihealthDrugYljgSearchbillRequest {
    s.PartnerIdRecv = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillRequest) SetBillCode(v string) *AlibabaAlihealthDrugYljgSearchbillRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillRequest) SetBillType(v string) *AlibabaAlihealthDrugYljgSearchbillRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillRequest) SetCurPage(v int64) *AlibabaAlihealthDrugYljgSearchbillRequest {
    s.CurPage = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillRequest) SetPageSize(v int64) *AlibabaAlihealthDrugYljgSearchbillRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthDrugYljgSearchbillRequest) ToMap() map[string]interface{} {
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

func (req *AlibabaAlihealthDrugYljgSearchbillRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}