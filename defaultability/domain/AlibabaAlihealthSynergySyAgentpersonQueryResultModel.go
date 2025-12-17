package domain


type AlibabaAlihealthSynergySyAgentpersonQueryResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        返回对象     */
    Model  *AlibabaAlihealthSynergySyAgentpersonQueryPage `json:"model,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

}

func (s *AlibabaAlihealthSynergySyAgentpersonQueryResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyAgentpersonQueryResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonQueryResultModel) SetModel(v AlibabaAlihealthSynergySyAgentpersonQueryPage) *AlibabaAlihealthSynergySyAgentpersonQueryResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonQueryResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyAgentpersonQueryResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonQueryResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyAgentpersonQueryResultModel {
    s.MsgInfo = &v
    return s
}
