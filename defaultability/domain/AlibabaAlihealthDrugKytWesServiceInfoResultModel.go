package domain


type AlibabaAlihealthDrugKytWesServiceInfoResultModel struct {
    /*
        业务结果成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        服务对象     */
    Model  *AlibabaAlihealthDrugKytWesServiceInfoWesServiceInfoDTO `json:"model,omitempty" `

    /*
        返回结果内容     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回结果的code     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesServiceInfoResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugKytWesServiceInfoResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesServiceInfoResultModel) SetModel(v AlibabaAlihealthDrugKytWesServiceInfoWesServiceInfoDTO) *AlibabaAlihealthDrugKytWesServiceInfoResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesServiceInfoResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugKytWesServiceInfoResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesServiceInfoResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugKytWesServiceInfoResultModel {
    s.MsgCode = &v
    return s
}
