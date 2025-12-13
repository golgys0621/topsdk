package domain


type AlibabaAlihealthDrugYljgQueryUpoutbillcountResultModel struct {
    /*
        返回成功失败     */
    Success  *bool `json:"success,omitempty" `

    /*
        model     */
    ModelList  *[]AlibabaAlihealthDrugYljgQueryUpoutbillcountBillTypeCountDTO `json:"model_list,omitempty" `

    /*
        msgInfo     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        msgCode     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugYljgQueryUpoutbillcountResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugYljgQueryUpoutbillcountResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgQueryUpoutbillcountResultModel) SetModelList(v []AlibabaAlihealthDrugYljgQueryUpoutbillcountBillTypeCountDTO) *AlibabaAlihealthDrugYljgQueryUpoutbillcountResultModel {
    s.ModelList = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgQueryUpoutbillcountResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugYljgQueryUpoutbillcountResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgQueryUpoutbillcountResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugYljgQueryUpoutbillcountResultModel {
    s.MsgCode = &v
    return s
}
