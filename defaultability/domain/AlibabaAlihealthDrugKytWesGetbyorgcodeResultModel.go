package domain


type AlibabaAlihealthDrugKytWesGetbyorgcodeResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugKytWesGetbyorgcodePUserEntInfoDto `json:"model,omitempty" `

    /*
        返回值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesGetbyorgcodeResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugKytWesGetbyorgcodeResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyorgcodeResultModel) SetModel(v AlibabaAlihealthDrugKytWesGetbyorgcodePUserEntInfoDto) *AlibabaAlihealthDrugKytWesGetbyorgcodeResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyorgcodeResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugKytWesGetbyorgcodeResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyorgcodeResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugKytWesGetbyorgcodeResultModel {
    s.MsgCode = &v
    return s
}
