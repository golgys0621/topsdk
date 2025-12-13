package domain


type AlibabaAlihealthDrugKytWesDrugrescodeResultModel struct {
    /*
        返回模型     */
    Model  *AlibabaAlihealthDrugKytWesDrugrescodePage `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugKytWesDrugrescodeResultModel) SetModel(v AlibabaAlihealthDrugKytWesDrugrescodePage) *AlibabaAlihealthDrugKytWesDrugrescodeResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugKytWesDrugrescodeResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugKytWesDrugrescodeResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugKytWesDrugrescodeResultModel {
    s.ResponseSuccess = &v
    return s
}
