package domain


type AlibabaAlihealthDrugWesQueryUpoutbillcountResultModel struct {
    /*
        返回成功失败     */
    Success  *bool `json:"success,omitempty" `

    /*
        model     */
    ModelList  *[]AlibabaAlihealthDrugWesQueryUpoutbillcountBillTypeCountDTO `json:"model_list,omitempty" `

    /*
        msgInfo     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        msgCode     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugWesQueryUpoutbillcountResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugWesQueryUpoutbillcountResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthDrugWesQueryUpoutbillcountResultModel) SetModelList(v []AlibabaAlihealthDrugWesQueryUpoutbillcountBillTypeCountDTO) *AlibabaAlihealthDrugWesQueryUpoutbillcountResultModel {
    s.ModelList = &v
    return s
}
func (s *AlibabaAlihealthDrugWesQueryUpoutbillcountResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugWesQueryUpoutbillcountResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugWesQueryUpoutbillcountResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugWesQueryUpoutbillcountResultModel {
    s.MsgCode = &v
    return s
}
