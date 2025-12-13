package request


type AlibabaAlihealthDrugKytWesSearchstatusRequest struct {
    /*
        企业ref_ent_id（货主企业的ref_ent_id）     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.getlicense     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        开始日期（没有时分秒，【单据创建时间】）     */
    BeginDate  *string `json:"begin_date" required:"true" `
    /*
        结束日期（没有时分秒，【单据创建时间】）     */
    EndDate  *string `json:"end_date" required:"true" `
    /*
        单据类型 A：全部 AI：全部入库 AO：全部出库     */
    BillType  *string `json:"bill_type" required:"true" `
    /*
        单据号（精确值，不支持模糊查询）     */
    BillCode  *string `json:"bill_code,omitempty" required:"false" `
    /*
        药品类型     */
    DrugType  *string `json:"drug_type,omitempty" required:"false" `
    /*
        状态  0, 处理中     3, 处理成功     4, 处理失败     */
    DealStatus  *string `json:"deal_status,omitempty" required:"false" `
    /*
        发货商     */
    FromUserId  *string `json:"from_user_id,omitempty" required:"false" `
    /*
        收货商     */
    ToUserId  *string `json:"to_user_id,omitempty" required:"false" `
    /*
        代理商（第三方物流企业）     */
    AgentRefUserId  *string `json:"agent_ref_user_id,omitempty" required:"false" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页码     */
    Page  *int64 `json:"page" required:"true" `
}

func (s *AlibabaAlihealthDrugKytWesSearchstatusRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesSearchstatusRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesSearchstatusRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusRequest) SetBeginDate(v string) *AlibabaAlihealthDrugKytWesSearchstatusRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusRequest) SetEndDate(v string) *AlibabaAlihealthDrugKytWesSearchstatusRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusRequest) SetBillType(v string) *AlibabaAlihealthDrugKytWesSearchstatusRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusRequest) SetBillCode(v string) *AlibabaAlihealthDrugKytWesSearchstatusRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusRequest) SetDrugType(v string) *AlibabaAlihealthDrugKytWesSearchstatusRequest {
    s.DrugType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusRequest) SetDealStatus(v string) *AlibabaAlihealthDrugKytWesSearchstatusRequest {
    s.DealStatus = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusRequest) SetFromUserId(v string) *AlibabaAlihealthDrugKytWesSearchstatusRequest {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusRequest) SetToUserId(v string) *AlibabaAlihealthDrugKytWesSearchstatusRequest {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusRequest) SetAgentRefUserId(v string) *AlibabaAlihealthDrugKytWesSearchstatusRequest {
    s.AgentRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusRequest) SetPageSize(v int64) *AlibabaAlihealthDrugKytWesSearchstatusRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusRequest) SetPage(v int64) *AlibabaAlihealthDrugKytWesSearchstatusRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugKytWesSearchstatusRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.LicenseToken != nil) {
        paramMap["license_token"] = *req.LicenseToken
    }
    if(req.BeginDate != nil) {
        paramMap["begin_date"] = *req.BeginDate
    }
    if(req.EndDate != nil) {
        paramMap["end_date"] = *req.EndDate
    }
    if(req.BillType != nil) {
        paramMap["bill_type"] = *req.BillType
    }
    if(req.BillCode != nil) {
        paramMap["bill_code"] = *req.BillCode
    }
    if(req.DrugType != nil) {
        paramMap["drug_type"] = *req.DrugType
    }
    if(req.DealStatus != nil) {
        paramMap["deal_status"] = *req.DealStatus
    }
    if(req.FromUserId != nil) {
        paramMap["from_user_id"] = *req.FromUserId
    }
    if(req.ToUserId != nil) {
        paramMap["to_user_id"] = *req.ToUserId
    }
    if(req.AgentRefUserId != nil) {
        paramMap["agent_ref_user_id"] = *req.AgentRefUserId
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugKytWesSearchstatusRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}