package domain


type AlibabaAlihealthSynergySyAgentpersonResourceListResultModel struct {
    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        资料列表     */
    Model  *[]AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel `json:"model,omitempty" `

    /*
        错误值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        错误码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyAgentpersonResourceListResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyAgentpersonResourceListResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceListResultModel) SetModel(v []AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel) *AlibabaAlihealthSynergySyAgentpersonResourceListResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceListResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyAgentpersonResourceListResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceListResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyAgentpersonResourceListResultModel {
    s.MsgCode = &v
    return s
}
