package domain


type AlibabaAlihealthSynergyYzwSignednotstampedbillQueryResultModel struct {
    /*
        是否响应成功(true |false)     */
    ResSuccess  *bool `json:"res_success,omitempty" `

    /*
        结果     */
    Model  *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryPage `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryResultModel) SetResSuccess(v bool) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryResultModel {
    s.ResSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryResultModel) SetModel(v AlibabaAlihealthSynergyYzwSignednotstampedbillQueryPage) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryResultModel {
    s.MsgCode = &v
    return s
}
