package domain


type AlibabaAlihealthSynergyYzwQueryassdrugreportinfoResultModel struct {
    /*
        外层大对象     */
    Model  *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoPageInfoDTO `json:"model,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    Success  *bool `json:"success,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoResultModel) SetModel(v AlibabaAlihealthSynergyYzwQueryassdrugreportinfoPageInfoDTO) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoResultModel) SetSuccess(v bool) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoResultModel {
    s.Success = &v
    return s
}
