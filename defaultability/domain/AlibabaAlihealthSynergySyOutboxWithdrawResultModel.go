package domain


type AlibabaAlihealthSynergySyOutboxWithdrawResultModel struct {
    /*
        是否调用成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        是否撤回成功     */
    Model  *bool `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyOutboxWithdrawResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyOutboxWithdrawResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxWithdrawResultModel) SetModel(v bool) *AlibabaAlihealthSynergySyOutboxWithdrawResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxWithdrawResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyOutboxWithdrawResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxWithdrawResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyOutboxWithdrawResultModel {
    s.MsgCode = &v
    return s
}
