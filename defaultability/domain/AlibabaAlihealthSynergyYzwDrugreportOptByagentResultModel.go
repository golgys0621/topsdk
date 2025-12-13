package domain


type AlibabaAlihealthSynergyYzwDrugreportOptByagentResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        返回对象     */
    Model  *AlibabaAlihealthSynergyYzwDrugreportOptByagentEntDealResultDTO `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDrugreportOptByagentResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergyYzwDrugreportOptByagentResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptByagentResultModel) SetModel(v AlibabaAlihealthSynergyYzwDrugreportOptByagentEntDealResultDTO) *AlibabaAlihealthSynergyYzwDrugreportOptByagentResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptByagentResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwDrugreportOptByagentResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptByagentResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwDrugreportOptByagentResultModel {
    s.MsgCode = &v
    return s
}
