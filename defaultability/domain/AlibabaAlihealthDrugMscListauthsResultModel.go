package domain


type AlibabaAlihealthDrugMscListauthsResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugMscListauthsPage `json:"model,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        返回值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugMscListauthsResultModel) SetModel(v AlibabaAlihealthDrugMscListauthsPage) *AlibabaAlihealthDrugMscListauthsResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListauthsResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscListauthsResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListauthsResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscListauthsResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListauthsResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMscListauthsResultModel {
    s.ResponseSuccess = &v
    return s
}
