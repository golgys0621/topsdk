package domain


type AlibabaAlihealthDrugLsydQueryUpoutbillcountResultModel struct {
    /*
        是否响应成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        model     */
    ModelList  *[]AlibabaAlihealthDrugLsydQueryUpoutbillcountBillTypeCountDTO `json:"model_list,omitempty" `

    /*
        msgInfo     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        msgCode     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugLsydQueryUpoutbillcountResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugLsydQueryUpoutbillcountResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydQueryUpoutbillcountResultModel) SetModelList(v []AlibabaAlihealthDrugLsydQueryUpoutbillcountBillTypeCountDTO) *AlibabaAlihealthDrugLsydQueryUpoutbillcountResultModel {
    s.ModelList = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydQueryUpoutbillcountResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugLsydQueryUpoutbillcountResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydQueryUpoutbillcountResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugLsydQueryUpoutbillcountResultModel {
    s.MsgCode = &v
    return s
}
