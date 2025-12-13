package domain


type AlibabaAlihealthDrugMscQueryUpoutbillcountResultModel struct {
    /*
        model     */
    ModelList  *[]AlibabaAlihealthDrugMscQueryUpoutbillcountBillTypeCountDTO `json:"model_list,omitempty" `

    /*
        msgInfo     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        msgCode     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        是否响应成功     */
    Success  *bool `json:"success,omitempty" `

}

func (s *AlibabaAlihealthDrugMscQueryUpoutbillcountResultModel) SetModelList(v []AlibabaAlihealthDrugMscQueryUpoutbillcountBillTypeCountDTO) *AlibabaAlihealthDrugMscQueryUpoutbillcountResultModel {
    s.ModelList = &v
    return s
}
func (s *AlibabaAlihealthDrugMscQueryUpoutbillcountResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscQueryUpoutbillcountResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscQueryUpoutbillcountResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscQueryUpoutbillcountResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscQueryUpoutbillcountResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugMscQueryUpoutbillcountResultModel {
    s.Success = &v
    return s
}
