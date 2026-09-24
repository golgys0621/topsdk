package domain


type AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        返回结果，空表示没有异常     */
    Model  *[]AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO `json:"model,omitempty" `

    /*
        返回信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoResultModel) SetModel(v []AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoResultModel {
    s.MsgCode = &v
    return s
}
