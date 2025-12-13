package domain


type AlibabaAlihealthDrugLsydQueryBillcountResultModel struct {
    /*
        返回成功失败     */
    Success  *bool `json:"success,omitempty" `

    /*
        model     */
    ModelList  *[]AlibabaAlihealthDrugLsydQueryBillcountBillTypeCountDTO `json:"model_list,omitempty" `

    /*
        msgInfo     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        msgCode     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugLsydQueryBillcountResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugLsydQueryBillcountResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydQueryBillcountResultModel) SetModelList(v []AlibabaAlihealthDrugLsydQueryBillcountBillTypeCountDTO) *AlibabaAlihealthDrugLsydQueryBillcountResultModel {
    s.ModelList = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydQueryBillcountResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugLsydQueryBillcountResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydQueryBillcountResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugLsydQueryBillcountResultModel {
    s.MsgCode = &v
    return s
}
