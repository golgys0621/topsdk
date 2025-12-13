package request


type AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest struct {
    /*
        企业唯一标识     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" required:"false" `
    /*
        企业自定义编号     */
    RefPartnerId  *string `json:"ref_partner_id,omitempty" required:"false" `
    /*
        开始时间：往来单位最后修改时间（不推荐使用：因为往来单位是共用的，任意企业提交了信息变更都会引起这个值的变更）     */
    BeginDate  *string `json:"begin_date,omitempty" required:"false" `
    /*
        结束时间：往来单位最后修改时间（不推荐使用：因为往来单位是共用的，任意企业提交了信息变更都会引起这个值的变更）     */
    EndDate  *string `json:"end_date,omitempty" required:"false" `
    /*
        1:审核通过，2：审核不通过     */
    AuditFlag  *int64 `json:"audit_flag,omitempty" required:"false" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页码     */
    Page  *int64 `json:"page" required:"true" `
}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) SetEntName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) SetRefPartnerId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest {
    s.RefPartnerId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) SetBeginDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) SetEndDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) SetAuditFlag(v int64) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest {
    s.AuditFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) SetPageSize(v int64) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) SetPage(v int64) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.EntName != nil) {
        paramMap["ent_name"] = *req.EntName
    }
    if(req.RefPartnerId != nil) {
        paramMap["ref_partner_id"] = *req.RefPartnerId
    }
    if(req.BeginDate != nil) {
        paramMap["begin_date"] = *req.BeginDate
    }
    if(req.EndDate != nil) {
        paramMap["end_date"] = *req.EndDate
    }
    if(req.AuditFlag != nil) {
        paramMap["audit_flag"] = *req.AuditFlag
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}