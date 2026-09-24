package domain


type AlibabaAlihealthDrugYljgSearchbillDetailResultModel struct {
    /*
        对象模型信息     */
    Model  *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugYljgSearchbillDetailResultModel) SetModel(v AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto) *AlibabaAlihealthDrugYljgSearchbillDetailResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugYljgSearchbillDetailResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugYljgSearchbillDetailResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugYljgSearchbillDetailResultModel {
    s.ResponseSuccess = &v
    return s
}
