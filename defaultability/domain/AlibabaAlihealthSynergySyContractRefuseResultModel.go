package domain


type AlibabaAlihealthSynergySyContractRefuseResultModel struct {
    /*
        是否调用成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        是否拒收成功     */
    Model  *bool `json:"model,omitempty" `

    /*
        返回值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyContractRefuseResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyContractRefuseResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractRefuseResultModel) SetModel(v bool) *AlibabaAlihealthSynergySyContractRefuseResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractRefuseResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyContractRefuseResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractRefuseResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyContractRefuseResultModel {
    s.MsgCode = &v
    return s
}
