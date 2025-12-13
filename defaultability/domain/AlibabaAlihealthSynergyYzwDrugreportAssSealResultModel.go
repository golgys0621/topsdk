package domain


type AlibabaAlihealthSynergyYzwDrugreportAssSealResultModel struct {
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

func (s *AlibabaAlihealthSynergyYzwDrugreportAssSealResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergyYzwDrugreportAssSealResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportAssSealResultModel) SetModel(v bool) *AlibabaAlihealthSynergyYzwDrugreportAssSealResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportAssSealResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwDrugreportAssSealResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportAssSealResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwDrugreportAssSealResultModel {
    s.MsgCode = &v
    return s
}
