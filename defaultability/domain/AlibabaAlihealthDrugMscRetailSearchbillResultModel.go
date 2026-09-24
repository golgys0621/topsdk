package domain


type AlibabaAlihealthDrugMscRetailSearchbillResultModel struct {
    /*
        返回模型     */
    Model  *AlibabaAlihealthDrugMscRetailSearchbillPageInfoDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugMscRetailSearchbillResultModel) SetModel(v AlibabaAlihealthDrugMscRetailSearchbillPageInfoDto) *AlibabaAlihealthDrugMscRetailSearchbillResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscRetailSearchbillResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscRetailSearchbillResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMscRetailSearchbillResultModel {
    s.ResponseSuccess = &v
    return s
}
