package domain


type AlibabaAlihealthDrugKytWesRetailSearchbillResultModel struct {
    /*
        返回模型     */
    Model  *AlibabaAlihealthDrugKytWesRetailSearchbillPageInfoDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugKytWesRetailSearchbillResultModel) SetModel(v AlibabaAlihealthDrugKytWesRetailSearchbillPageInfoDto) *AlibabaAlihealthDrugKytWesRetailSearchbillResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugKytWesRetailSearchbillResultModel {
    s.Success = &v
    return s
}
