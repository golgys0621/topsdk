package domain


type AlibabaAlihealthSynergySyContractRevokeResultModel struct {
    /*
        是否调用成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        是否撤回成功     */
    Model  *bool `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyContractRevokeResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyContractRevokeResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractRevokeResultModel) SetModel(v bool) *AlibabaAlihealthSynergySyContractRevokeResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractRevokeResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyContractRevokeResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractRevokeResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyContractRevokeResultModel {
    s.MsgCode = &v
    return s
}
