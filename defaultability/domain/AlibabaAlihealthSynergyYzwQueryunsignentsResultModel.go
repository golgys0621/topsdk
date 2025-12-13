package domain


type AlibabaAlihealthSynergyYzwQueryunsignentsResultModel struct {
    /*
        结果     */
    Model  *[]AlibabaAlihealthSynergyYzwQueryunsignentsHoloBillSearchCommonShowDTO `json:"model,omitempty" `

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

func (s *AlibabaAlihealthSynergyYzwQueryunsignentsResultModel) SetModel(v []AlibabaAlihealthSynergyYzwQueryunsignentsHoloBillSearchCommonShowDTO) *AlibabaAlihealthSynergyYzwQueryunsignentsResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryunsignentsResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwQueryunsignentsResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryunsignentsResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwQueryunsignentsResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryunsignentsResultModel) SetResponseSuccess(v string) *AlibabaAlihealthSynergyYzwQueryunsignentsResultModel {
    s.ResponseSuccess = &v
    return s
}
