package domain


type AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        返回对象类型是Map，属性如下 id（报告id）和 sealed_report_url（盖章后的链接）     */
    Model  *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfResData `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfResultModel) SetModel(v AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfResData) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfResultModel {
    s.MsgCode = &v
    return s
}
