package domain


type AlibabaAlihealthSynergySyAgentpersonResourceSaveResultModel struct {
    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        任务id     */
    Model  *int64 `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyAgentpersonResourceSaveResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyAgentpersonResourceSaveResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceSaveResultModel) SetModel(v int64) *AlibabaAlihealthSynergySyAgentpersonResourceSaveResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceSaveResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyAgentpersonResourceSaveResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceSaveResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyAgentpersonResourceSaveResultModel {
    s.MsgCode = &v
    return s
}
