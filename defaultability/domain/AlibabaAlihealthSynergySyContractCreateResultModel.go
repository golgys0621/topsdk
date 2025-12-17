package domain


type AlibabaAlihealthSynergySyContractCreateResultModel struct {
    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        合同Id     */
    Model  *int64 `json:"model,omitempty" `

    /*
        返回值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyContractCreateResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyContractCreateResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractCreateResultModel) SetModel(v int64) *AlibabaAlihealthSynergySyContractCreateResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractCreateResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyContractCreateResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractCreateResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyContractCreateResultModel {
    s.MsgCode = &v
    return s
}
