package domain


type AlibabaAlihealthDrugKytWesSearchbillDetailResultModel struct {
    /*
        对象模型信息     */
    Model  *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto `json:"model,omitempty" `

    /*
        消息码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        消息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        成功失败     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesSearchbillDetailResultModel) SetModel(v AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto) *AlibabaAlihealthDrugKytWesSearchbillDetailResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugKytWesSearchbillDetailResultModel {
    s.ResponseSuccess = &v
    return s
}
