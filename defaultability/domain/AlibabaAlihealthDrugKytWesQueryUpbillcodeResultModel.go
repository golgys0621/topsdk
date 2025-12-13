package domain


type AlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel struct {
    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        model     */
    ModelList  *[]AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO `json:"model_list,omitempty" `

    /*
        msgInfo     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        msgCode     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel) SetModelList(v []AlibabaAlihealthDrugKytWesQueryUpbillcodeBillUpstreamDTO) *AlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel {
    s.ModelList = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugKytWesQueryUpbillcodeResultModel {
    s.MsgCode = &v
    return s
}
