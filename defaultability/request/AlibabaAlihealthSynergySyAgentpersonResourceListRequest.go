package request


type AlibabaAlihealthSynergySyAgentpersonResourceListRequest struct {
    /*
        refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        委托人id，通过本企业委托人接口获取ID或者通过合作企业委托人查询接口获取ID     */
    SyAgentPersonId  *string `json:"sy_agent_person_id" required:"true" `
    /*
        使用状态，0已过期1使用中     */
    UseStatus  *string `json:"use_status,omitempty" required:"false" `
    /*
        归档状态，0未归档1已归档     */
    ArchiveStatus  *string `json:"archive_status,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergySyAgentpersonResourceListRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyAgentpersonResourceListRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceListRequest) SetSyAgentPersonId(v string) *AlibabaAlihealthSynergySyAgentpersonResourceListRequest {
    s.SyAgentPersonId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceListRequest) SetUseStatus(v string) *AlibabaAlihealthSynergySyAgentpersonResourceListRequest {
    s.UseStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceListRequest) SetArchiveStatus(v string) *AlibabaAlihealthSynergySyAgentpersonResourceListRequest {
    s.ArchiveStatus = &v
    return s
}

func (req *AlibabaAlihealthSynergySyAgentpersonResourceListRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.SyAgentPersonId != nil) {
        paramMap["sy_agent_person_id"] = *req.SyAgentPersonId
    }
    if(req.UseStatus != nil) {
        paramMap["use_status"] = *req.UseStatus
    }
    if(req.ArchiveStatus != nil) {
        paramMap["archive_status"] = *req.ArchiveStatus
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyAgentpersonResourceListRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}