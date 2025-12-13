package domain


type AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel struct {
    /*
        最外层对象     */
    Model  *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto `json:"model,omitempty" `

    /*
        提示信息编码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        提示信息内容     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        成功失败标记     */
    Success  *bool `json:"success,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel) SetModel(v AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResultModel {
    s.Success = &v
    return s
}
