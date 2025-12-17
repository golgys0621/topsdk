package domain


type AlibabaAlihealthSynergySyInboxListResultModel struct {
    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        结果     */
    Model  *AlibabaAlihealthSynergySyInboxListPage `json:"model,omitempty" `

    /*
        返回值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyInboxListResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyInboxListResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxListResultModel) SetModel(v AlibabaAlihealthSynergySyInboxListPage) *AlibabaAlihealthSynergySyInboxListResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxListResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyInboxListResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxListResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyInboxListResultModel {
    s.MsgCode = &v
    return s
}
