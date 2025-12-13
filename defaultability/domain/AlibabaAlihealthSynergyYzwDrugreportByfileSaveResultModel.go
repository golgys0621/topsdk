package domain


type AlibabaAlihealthSynergyYzwDrugreportByfileSaveResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        返回对象类型是Map，属性如下 id（报告id）和 sealed_report_url（盖章后的链接）     */
    Model  *AlibabaAlihealthSynergyYzwDrugreportByfileSaveResData `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveResultModel) SetModel(v AlibabaAlihealthSynergyYzwDrugreportByfileSaveResData) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveResultModel {
    s.MsgCode = &v
    return s
}
