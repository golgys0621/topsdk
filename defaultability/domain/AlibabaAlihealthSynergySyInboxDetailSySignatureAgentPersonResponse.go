package domain


type AlibabaAlihealthSynergySyInboxDetailSySignatureAgentPersonResponse struct {
    /*
        委托人id     */
    AgentPersonId  *int64 `json:"agent_person_id,omitempty" `

    /*
        委托人姓名     */
    Username  *string `json:"username,omitempty" `

    /*
        委托人编号     */
    UserNo  *string `json:"user_no,omitempty" `

    /*
        电话     */
    Tel  *string `json:"tel,omitempty" `

    /*
        资料数     */
    ResourceCount  *int64 `json:"resource_count,omitempty" `

    /*
        委托人资料列表     */
    AgentPersonResources  *[]AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse `json:"agent_person_resources,omitempty" `

    /*
        委托人详情     */
    AgentDetail  *AlibabaAlihealthSynergySyInboxDetailSyAgentPersonDTO `json:"agent_detail,omitempty" `

}

func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureAgentPersonResponse) SetAgentPersonId(v int64) *AlibabaAlihealthSynergySyInboxDetailSySignatureAgentPersonResponse {
    s.AgentPersonId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureAgentPersonResponse) SetUsername(v string) *AlibabaAlihealthSynergySyInboxDetailSySignatureAgentPersonResponse {
    s.Username = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureAgentPersonResponse) SetUserNo(v string) *AlibabaAlihealthSynergySyInboxDetailSySignatureAgentPersonResponse {
    s.UserNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureAgentPersonResponse) SetTel(v string) *AlibabaAlihealthSynergySyInboxDetailSySignatureAgentPersonResponse {
    s.Tel = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureAgentPersonResponse) SetResourceCount(v int64) *AlibabaAlihealthSynergySyInboxDetailSySignatureAgentPersonResponse {
    s.ResourceCount = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureAgentPersonResponse) SetAgentPersonResources(v []AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse) *AlibabaAlihealthSynergySyInboxDetailSySignatureAgentPersonResponse {
    s.AgentPersonResources = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureAgentPersonResponse) SetAgentDetail(v AlibabaAlihealthSynergySyInboxDetailSyAgentPersonDTO) *AlibabaAlihealthSynergySyInboxDetailSySignatureAgentPersonResponse {
    s.AgentDetail = &v
    return s
}
