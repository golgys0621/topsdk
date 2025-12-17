package domain


type AlibabaAlihealthSynergySyContractDetailResultModel struct {
    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        结果     */
    Model  *AlibabaAlihealthSynergySyContractDetailSyContractDetail `json:"model,omitempty" `

    /*
        是否调用成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthSynergySyContractDetailResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyContractDetailResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyContractDetailResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailResultModel) SetModel(v AlibabaAlihealthSynergySyContractDetailSyContractDetail) *AlibabaAlihealthSynergySyContractDetailResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyContractDetailResultModel {
    s.ResponseSuccess = &v
    return s
}
