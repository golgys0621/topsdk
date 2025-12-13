package domain


type AlibabaAlihealthSynergyYzwQuerydrugreportResultModel struct {
    /*
        是否响应成功(true |false)     */
    ResSuccess  *bool `json:"res_success,omitempty" `

    /*
        结果     */
    Model  *AlibabaAlihealthSynergyYzwQuerydrugreportPage `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQuerydrugreportResultModel) SetResSuccess(v bool) *AlibabaAlihealthSynergyYzwQuerydrugreportResultModel {
    s.ResSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportResultModel) SetModel(v AlibabaAlihealthSynergyYzwQuerydrugreportPage) *AlibabaAlihealthSynergyYzwQuerydrugreportResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportResultModel {
    s.MsgCode = &v
    return s
}
