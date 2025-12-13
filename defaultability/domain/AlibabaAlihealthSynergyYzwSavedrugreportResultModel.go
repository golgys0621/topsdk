package domain


type AlibabaAlihealthSynergyYzwSavedrugreportResultModel struct {
    /*
        返回对象类型是Map，属性如下 id（报告id）和 sealed_report_url（盖章后的链接）     */
    Model  *AlibabaAlihealthSynergyYzwSavedrugreportResData `json:"model,omitempty" `

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

func (s *AlibabaAlihealthSynergyYzwSavedrugreportResultModel) SetModel(v AlibabaAlihealthSynergyYzwSavedrugreportResData) *AlibabaAlihealthSynergyYzwSavedrugreportResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwSavedrugreportResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwSavedrugreportResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportResultModel) SetResponseSuccess(v string) *AlibabaAlihealthSynergyYzwSavedrugreportResultModel {
    s.ResponseSuccess = &v
    return s
}
