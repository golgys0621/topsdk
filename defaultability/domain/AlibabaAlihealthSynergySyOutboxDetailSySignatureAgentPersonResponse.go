package domain


type AlibabaAlihealthSynergySyOutboxDetailSySignatureAgentPersonResponse struct {
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
    AgentPersonResources  *[]AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse `json:"agent_person_resources,omitempty" `

}

func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureAgentPersonResponse) SetAgentPersonId(v int64) *AlibabaAlihealthSynergySyOutboxDetailSySignatureAgentPersonResponse {
    s.AgentPersonId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureAgentPersonResponse) SetUsername(v string) *AlibabaAlihealthSynergySyOutboxDetailSySignatureAgentPersonResponse {
    s.Username = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureAgentPersonResponse) SetUserNo(v string) *AlibabaAlihealthSynergySyOutboxDetailSySignatureAgentPersonResponse {
    s.UserNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureAgentPersonResponse) SetTel(v string) *AlibabaAlihealthSynergySyOutboxDetailSySignatureAgentPersonResponse {
    s.Tel = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureAgentPersonResponse) SetResourceCount(v int64) *AlibabaAlihealthSynergySyOutboxDetailSySignatureAgentPersonResponse {
    s.ResourceCount = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureAgentPersonResponse) SetAgentPersonResources(v []AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse) *AlibabaAlihealthSynergySyOutboxDetailSySignatureAgentPersonResponse {
    s.AgentPersonResources = &v
    return s
}
