package request


type AlibabaAlihealthDrugMscListpartByagentRequest struct {
    /*
        企业唯一标识（货主企业）     */
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
        代理企业唯一标识（物流企业）     */
    AgentRefEntId  *string `json:"agent_ref_ent_id,omitempty" required:"false" `
    /*
        1审核通过、2审核不通过     */
    AuditFlag  *int64 `json:"audit_flag,omitempty" required:"false" `
    /*
        1:待替换 2:待更新     */
    Shared  *string `json:"shared,omitempty" required:"false" `
    /*
        唯一认证代码     */
    OrgCode  *string `json:"org_code,omitempty" required:"false" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页码     */
    Page  *int64 `json:"page" required:"true" `
}

func (s *AlibabaAlihealthDrugMscListpartByagentRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMscListpartByagentRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentRequest) SetEntName(v string) *AlibabaAlihealthDrugMscListpartByagentRequest {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentRequest) SetRefPartnerId(v string) *AlibabaAlihealthDrugMscListpartByagentRequest {
    s.RefPartnerId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentRequest) SetBeginDate(v string) *AlibabaAlihealthDrugMscListpartByagentRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentRequest) SetEndDate(v string) *AlibabaAlihealthDrugMscListpartByagentRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentRequest) SetAgentRefEntId(v string) *AlibabaAlihealthDrugMscListpartByagentRequest {
    s.AgentRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentRequest) SetAuditFlag(v int64) *AlibabaAlihealthDrugMscListpartByagentRequest {
    s.AuditFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentRequest) SetShared(v string) *AlibabaAlihealthDrugMscListpartByagentRequest {
    s.Shared = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentRequest) SetOrgCode(v string) *AlibabaAlihealthDrugMscListpartByagentRequest {
    s.OrgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentRequest) SetPageSize(v int64) *AlibabaAlihealthDrugMscListpartByagentRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentRequest) SetPage(v int64) *AlibabaAlihealthDrugMscListpartByagentRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugMscListpartByagentRequest) ToMap() map[string]interface{} {
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
    if(req.AgentRefEntId != nil) {
        paramMap["agent_ref_ent_id"] = *req.AgentRefEntId
    }
    if(req.AuditFlag != nil) {
        paramMap["audit_flag"] = *req.AuditFlag
    }
    if(req.Shared != nil) {
        paramMap["shared"] = *req.Shared
    }
    if(req.OrgCode != nil) {
        paramMap["org_code"] = *req.OrgCode
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugMscListpartByagentRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}