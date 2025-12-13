package domain


type AlibabaAlihealthSynergyYzwDrugreportOptResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        返回对象     */
    Model  *AlibabaAlihealthSynergyYzwDrugreportOptEntDealResultDTO `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDrugreportOptResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergyYzwDrugreportOptResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptResultModel) SetModel(v AlibabaAlihealthSynergyYzwDrugreportOptEntDealResultDTO) *AlibabaAlihealthSynergyYzwDrugreportOptResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwDrugreportOptResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwDrugreportOptResultModel {
    s.MsgCode = &v
    return s
}
