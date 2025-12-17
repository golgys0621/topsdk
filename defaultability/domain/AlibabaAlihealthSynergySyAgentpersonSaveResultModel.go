package domain


type AlibabaAlihealthSynergySyAgentpersonSaveResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        委托人id     */
    Model  *int64 `json:"model,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

}

func (s *AlibabaAlihealthSynergySyAgentpersonSaveResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyAgentpersonSaveResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonSaveResultModel) SetModel(v int64) *AlibabaAlihealthSynergySyAgentpersonSaveResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonSaveResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyAgentpersonSaveResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonSaveResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyAgentpersonSaveResultModel {
    s.MsgInfo = &v
    return s
}
