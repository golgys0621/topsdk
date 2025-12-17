package domain


type AlibabaAlihealthSynergySyContractSealResultModel struct {
    /*
        是否调用成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        是否盖章成功     */
    Model  *bool `json:"model,omitempty" `

    /*
        返回值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyContractSealResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyContractSealResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractSealResultModel) SetModel(v bool) *AlibabaAlihealthSynergySyContractSealResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractSealResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyContractSealResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractSealResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyContractSealResultModel {
    s.MsgCode = &v
    return s
}
