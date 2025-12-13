package domain


type AlibabaAlihealthDrugMyjCodewarnGetwarninfoResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        返回结果，空表示没有异常     */
    Model  *[]AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO `json:"model,omitempty" `

    /*
        返回信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugMyjCodewarnGetwarninfoResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMyjCodewarnGetwarninfoResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnGetwarninfoResultModel) SetModel(v []AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) *AlibabaAlihealthDrugMyjCodewarnGetwarninfoResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnGetwarninfoResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMyjCodewarnGetwarninfoResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnGetwarninfoResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMyjCodewarnGetwarninfoResultModel {
    s.MsgCode = &v
    return s
}
