package domain


type AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassResultModel struct {
    /*
        任务id     */
    Model  *int64 `json:"model,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    ResponseSuccess  *string `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassResultModel) SetModel(v int64) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassResultModel) SetResponseSuccess(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassResultModel {
    s.ResponseSuccess = &v
    return s
}
