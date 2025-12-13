package domain


type AlibabaAlihealthDrugLsydSearchbillResultModel struct {
    /*
        返回模型     */
    Model  *AlibabaAlihealthDrugLsydSearchbillPageInfoDto `json:"model,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        返回信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugLsydSearchbillResultModel) SetModel(v AlibabaAlihealthDrugLsydSearchbillPageInfoDto) *AlibabaAlihealthDrugLsydSearchbillResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugLsydSearchbillResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugLsydSearchbillResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugLsydSearchbillResultModel {
    s.ResponseSuccess = &v
    return s
}
