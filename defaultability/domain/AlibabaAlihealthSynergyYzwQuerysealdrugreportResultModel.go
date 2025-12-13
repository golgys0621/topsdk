package domain


type AlibabaAlihealthSynergyYzwQuerysealdrugreportResultModel struct {
    /*
        返回模型     */
    Model  *AlibabaAlihealthSynergyYzwQuerysealdrugreportPage `json:"model,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportResultModel) SetModel(v AlibabaAlihealthSynergyYzwQuerysealdrugreportPage) *AlibabaAlihealthSynergyYzwQuerysealdrugreportResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwQuerysealdrugreportResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwQuerysealdrugreportResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergyYzwQuerysealdrugreportResultModel {
    s.ResponseSuccess = &v
    return s
}
