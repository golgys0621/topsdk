package domain


type AlibabaAlihealthDrugYljgRetailSearchbillResultModel struct {
    /*
        返回模型     */
    Model  *AlibabaAlihealthDrugYljgRetailSearchbillPageInfoDTO `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugYljgRetailSearchbillResultModel) SetModel(v AlibabaAlihealthDrugYljgRetailSearchbillPageInfoDTO) *AlibabaAlihealthDrugYljgRetailSearchbillResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugYljgRetailSearchbillResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugYljgRetailSearchbillResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugYljgRetailSearchbillResultModel {
    s.ResponseSuccess = &v
    return s
}
