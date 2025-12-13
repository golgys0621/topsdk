package domain


type AlibabaAlihealthDrugMscQueryBillcountResultModel struct {
    /*
        model     */
    ModelList  *[]AlibabaAlihealthDrugMscQueryBillcountBillTypeCountDTO `json:"model_list,omitempty" `

    /*
        msgInfo     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        msgCode     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

}

func (s *AlibabaAlihealthDrugMscQueryBillcountResultModel) SetModelList(v []AlibabaAlihealthDrugMscQueryBillcountBillTypeCountDTO) *AlibabaAlihealthDrugMscQueryBillcountResultModel {
    s.ModelList = &v
    return s
}
func (s *AlibabaAlihealthDrugMscQueryBillcountResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscQueryBillcountResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscQueryBillcountResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscQueryBillcountResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscQueryBillcountResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugMscQueryBillcountResultModel {
    s.Success = &v
    return s
}
