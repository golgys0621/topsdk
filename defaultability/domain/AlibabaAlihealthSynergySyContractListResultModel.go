package domain


type AlibabaAlihealthSynergySyContractListResultModel struct {
    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        结果     */
    Model  *AlibabaAlihealthSynergySyContractListPage `json:"model,omitempty" `

    /*
        调用成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthSynergySyContractListResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyContractListResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyContractListResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListResultModel) SetModel(v AlibabaAlihealthSynergySyContractListPage) *AlibabaAlihealthSynergySyContractListResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyContractListResultModel {
    s.ResponseSuccess = &v
    return s
}
