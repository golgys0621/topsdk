package domain


type AlibabaAlihealthDrugCodeKytWesLicenseTokenGetResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        token信息     */
    Model  *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetTokenInfo `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetResultModel) SetModel(v AlibabaAlihealthDrugCodeKytWesLicenseTokenGetTokenInfo) *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetResultModel {
    s.MsgCode = &v
    return s
}
