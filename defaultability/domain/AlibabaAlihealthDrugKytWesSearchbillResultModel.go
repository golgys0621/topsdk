package domain


type AlibabaAlihealthDrugKytWesSearchbillResultModel struct {
    /*
        返回模型     */
    Model  *AlibabaAlihealthDrugKytWesSearchbillPageInfoDto `json:"model,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        返回信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    Success  *bool `json:"success,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesSearchbillResultModel) SetModel(v AlibabaAlihealthDrugKytWesSearchbillPageInfoDto) *AlibabaAlihealthDrugKytWesSearchbillResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugKytWesSearchbillResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugKytWesSearchbillResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugKytWesSearchbillResultModel {
    s.Success = &v
    return s
}
