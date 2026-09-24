package domain


type AlibabaAlihealthDrugLsydRetailSearchbillResultModel struct {
    /*
        返回模型     */
    Model  *AlibabaAlihealthDrugLsydRetailSearchbillPageInfoDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugLsydRetailSearchbillResultModel) SetModel(v AlibabaAlihealthDrugLsydRetailSearchbillPageInfoDto) *AlibabaAlihealthDrugLsydRetailSearchbillResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugLsydRetailSearchbillResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugLsydRetailSearchbillResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugLsydRetailSearchbillResultModel {
    s.ResponseSuccess = &v
    return s
}
