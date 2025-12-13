package domain


type AlibabaAlihealthSynergyYzwDrugreportSealResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        true盖章成功、false盖章失败     */
    Model  *bool `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDrugreportSealResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergyYzwDrugreportSealResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportSealResultModel) SetModel(v bool) *AlibabaAlihealthSynergyYzwDrugreportSealResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportSealResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwDrugreportSealResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportSealResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwDrugreportSealResultModel {
    s.MsgCode = &v
    return s
}
