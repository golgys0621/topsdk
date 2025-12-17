package domain


type AlibabaAlihealthSynergySyPartneragentpersonQueryResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        返回对象     */
    Model  *AlibabaAlihealthSynergySyPartneragentpersonQueryPage `json:"model,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

}

func (s *AlibabaAlihealthSynergySyPartneragentpersonQueryResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyPartneragentpersonQueryResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQueryResultModel) SetModel(v AlibabaAlihealthSynergySyPartneragentpersonQueryPage) *AlibabaAlihealthSynergySyPartneragentpersonQueryResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQueryResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyPartneragentpersonQueryResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQueryResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyPartneragentpersonQueryResultModel {
    s.MsgInfo = &v
    return s
}
