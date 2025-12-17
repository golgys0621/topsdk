package domain


type AlibabaAlihealthSynergySySignatureDetailGetSySignatureAgentPersonResponse struct {
    /*
        委托人id     */
    AgentPersonId  *int64 `json:"agent_person_id,omitempty" `

    /*
        委托人名字     */
    Username  *string `json:"username,omitempty" `

    /*
        委托人编号     */
    UserNo  *string `json:"user_no,omitempty" `

    /*
        电话     */
    Tel  *string `json:"tel,omitempty" `

    /*
        委托人所属资料的总页数     */
    ResourceCount  *int64 `json:"resource_count,omitempty" `

    /*
        委托人资料集合     */
    AgentPersonResourcesList  *[]AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse `json:"agent_person_resources_list,omitempty" `

}

func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureAgentPersonResponse) SetAgentPersonId(v int64) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureAgentPersonResponse {
    s.AgentPersonId = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureAgentPersonResponse) SetUsername(v string) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureAgentPersonResponse {
    s.Username = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureAgentPersonResponse) SetUserNo(v string) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureAgentPersonResponse {
    s.UserNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureAgentPersonResponse) SetTel(v string) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureAgentPersonResponse {
    s.Tel = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureAgentPersonResponse) SetResourceCount(v int64) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureAgentPersonResponse {
    s.ResourceCount = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureAgentPersonResponse) SetAgentPersonResourcesList(v []AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureAgentPersonResponse {
    s.AgentPersonResourcesList = &v
    return s
}
