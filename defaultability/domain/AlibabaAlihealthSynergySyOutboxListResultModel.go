package domain


type AlibabaAlihealthSynergySyOutboxListResultModel struct {
    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        结果     */
    Model  *AlibabaAlihealthSynergySyOutboxListPage `json:"model,omitempty" `

    /*
        返回值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyOutboxListResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyOutboxListResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListResultModel) SetModel(v AlibabaAlihealthSynergySyOutboxListPage) *AlibabaAlihealthSynergySyOutboxListResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyOutboxListResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyOutboxListResultModel {
    s.MsgCode = &v
    return s
}
