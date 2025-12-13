package domain


type AlibabaAlihealthDrugWesListauthsResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugWesListauthsPage `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugWesListauthsResultModel) SetModel(v AlibabaAlihealthDrugWesListauthsPage) *AlibabaAlihealthDrugWesListauthsResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugWesListauthsResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugWesListauthsResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugWesListauthsResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugWesListauthsResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugWesListauthsResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugWesListauthsResultModel {
    s.ResponseSuccess = &v
    return s
}
