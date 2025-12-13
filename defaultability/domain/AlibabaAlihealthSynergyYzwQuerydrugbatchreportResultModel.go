package domain


type AlibabaAlihealthSynergyYzwQuerydrugbatchreportResultModel struct {
    /*
        返回模型     */
    Model  *AlibabaAlihealthSynergyYzwQuerydrugbatchreportPage `json:"model,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportResultModel) SetModel(v AlibabaAlihealthSynergyYzwQuerydrugbatchreportPage) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportResultModel {
    s.ResponseSuccess = &v
    return s
}
