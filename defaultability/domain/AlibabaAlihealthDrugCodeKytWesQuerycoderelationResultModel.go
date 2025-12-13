package domain


type AlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel struct {
    /*
        model     */
    ModelList  *[]AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto `json:"model_list,omitempty" `

    /*
        msgCode     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        msgInfo     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel) SetModelList(v []AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel {
    s.ModelList = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel {
    s.ResponseSuccess = &v
    return s
}
