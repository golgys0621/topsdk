package domain


type AlibabaAlihealthDrugCodeKytWesGetlicenseResultModel struct {
    /*
        licenseToken，有效期1天，有效期内返回值不变     */
    LicenseToken  *string `json:"license_token,omitempty" `

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

func (s *AlibabaAlihealthDrugCodeKytWesGetlicenseResultModel) SetLicenseToken(v string) *AlibabaAlihealthDrugCodeKytWesGetlicenseResultModel {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesGetlicenseResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugCodeKytWesGetlicenseResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesGetlicenseResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugCodeKytWesGetlicenseResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesGetlicenseResultModel) SetResponseSuccess(v string) *AlibabaAlihealthDrugCodeKytWesGetlicenseResultModel {
    s.ResponseSuccess = &v
    return s
}
