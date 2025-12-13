package domain


type AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        返回对象     */
    Model  *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllPage `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllResultModel) SetModel(v AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllPage) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllResultModel {
    s.MsgCode = &v
    return s
}
