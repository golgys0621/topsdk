package domain


type AlibabaAlihealthSynergySyResourceOptArchiveResultModel struct {
    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        true成功，false失败     */
    Model  *bool `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyResourceOptArchiveResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyResourceOptArchiveResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyResourceOptArchiveResultModel) SetModel(v bool) *AlibabaAlihealthSynergySyResourceOptArchiveResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyResourceOptArchiveResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyResourceOptArchiveResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyResourceOptArchiveResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyResourceOptArchiveResultModel {
    s.MsgCode = &v
    return s
}
