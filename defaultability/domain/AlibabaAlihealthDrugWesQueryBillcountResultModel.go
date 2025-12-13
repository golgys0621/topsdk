package domain


type AlibabaAlihealthDrugWesQueryBillcountResultModel struct {
    /*
        返回成功失败     */
    Success  *bool `json:"success,omitempty" `

    /*
        model     */
    ModelList  *[]AlibabaAlihealthDrugWesQueryBillcountBillTypeCountDTO `json:"model_list,omitempty" `

    /*
        msgInfo     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        msgCode     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugWesQueryBillcountResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugWesQueryBillcountResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthDrugWesQueryBillcountResultModel) SetModelList(v []AlibabaAlihealthDrugWesQueryBillcountBillTypeCountDTO) *AlibabaAlihealthDrugWesQueryBillcountResultModel {
    s.ModelList = &v
    return s
}
func (s *AlibabaAlihealthDrugWesQueryBillcountResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugWesQueryBillcountResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugWesQueryBillcountResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugWesQueryBillcountResultModel {
    s.MsgCode = &v
    return s
}
