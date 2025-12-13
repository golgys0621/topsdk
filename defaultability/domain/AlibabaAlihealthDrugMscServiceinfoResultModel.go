package domain


type AlibabaAlihealthDrugMscServiceinfoResultModel struct {
    /*
        业务结果成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        服务对象     */
    Model  *AlibabaAlihealthDrugMscServiceinfoMSCServiceInfoDTO `json:"model,omitempty" `

    /*
        返回结果内容     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回结果的code     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugMscServiceinfoResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMscServiceinfoResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugMscServiceinfoResultModel) SetModel(v AlibabaAlihealthDrugMscServiceinfoMSCServiceInfoDTO) *AlibabaAlihealthDrugMscServiceinfoResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMscServiceinfoResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscServiceinfoResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscServiceinfoResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscServiceinfoResultModel {
    s.MsgCode = &v
    return s
}
