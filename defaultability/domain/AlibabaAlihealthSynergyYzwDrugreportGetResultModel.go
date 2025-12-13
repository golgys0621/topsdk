package domain


type AlibabaAlihealthSynergyYzwDrugreportGetResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        报告对象     */
    Model  *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDrugreportGetResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergyYzwDrugreportGetResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportGetResultModel) SetModel(v AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO) *AlibabaAlihealthSynergyYzwDrugreportGetResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportGetResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwDrugreportGetResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportGetResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwDrugreportGetResultModel {
    s.MsgCode = &v
    return s
}
