package domain


type AlibabaAlihealthSynergySyInboxCheckResultModel struct {
    /*
        调用是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        查收是否成功     */
    Model  *bool `json:"model,omitempty" `

    /*
        返回值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyInboxCheckResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyInboxCheckResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxCheckResultModel) SetModel(v bool) *AlibabaAlihealthSynergySyInboxCheckResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxCheckResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyInboxCheckResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxCheckResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyInboxCheckResultModel {
    s.MsgCode = &v
    return s
}
