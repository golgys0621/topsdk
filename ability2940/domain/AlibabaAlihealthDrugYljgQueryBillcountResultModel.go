package domain


type AlibabaAlihealthDrugYljgQueryBillcountResultModel struct {
    /*
        返回成功失败     */
    Success  *bool `json:"success,omitempty" `

    /*
        model     */
    ModelList  *[]AlibabaAlihealthDrugYljgQueryBillcountBillTypeCountDTO `json:"model_list,omitempty" `

    /*
        msgInfo     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        msgCode     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugYljgQueryBillcountResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugYljgQueryBillcountResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgQueryBillcountResultModel) SetModelList(v []AlibabaAlihealthDrugYljgQueryBillcountBillTypeCountDTO) *AlibabaAlihealthDrugYljgQueryBillcountResultModel {
    s.ModelList = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgQueryBillcountResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugYljgQueryBillcountResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgQueryBillcountResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugYljgQueryBillcountResultModel {
    s.MsgCode = &v
    return s
}
