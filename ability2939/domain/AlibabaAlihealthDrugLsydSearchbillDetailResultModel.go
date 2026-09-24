package domain


type AlibabaAlihealthDrugLsydSearchbillDetailResultModel struct {
    /*
        对象模型信息     */
    Model  *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugLsydSearchbillDetailResultModel) SetModel(v AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto) *AlibabaAlihealthDrugLsydSearchbillDetailResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugLsydSearchbillDetailResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugLsydSearchbillDetailResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugLsydSearchbillDetailResultModel {
    s.ResponseSuccess = &v
    return s
}
