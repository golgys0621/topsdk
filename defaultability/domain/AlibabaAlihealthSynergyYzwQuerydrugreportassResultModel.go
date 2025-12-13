package domain


type AlibabaAlihealthSynergyYzwQuerydrugreportassResultModel struct {
    /*
        是否响应成功(true |false)     */
    ResSuccess  *bool `json:"res_success,omitempty" `

    /*
        结果     */
    Model  *AlibabaAlihealthSynergyYzwQuerydrugreportassPage `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassResultModel) SetResSuccess(v bool) *AlibabaAlihealthSynergyYzwQuerydrugreportassResultModel {
    s.ResSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassResultModel) SetModel(v AlibabaAlihealthSynergyYzwQuerydrugreportassPage) *AlibabaAlihealthSynergyYzwQuerydrugreportassResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassResultModel {
    s.MsgCode = &v
    return s
}
