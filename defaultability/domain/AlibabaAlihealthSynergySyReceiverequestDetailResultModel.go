package domain


type AlibabaAlihealthSynergySyReceiverequestDetailResultModel struct {
    /*
        是否响应成功(true |false)     */
    Success  *bool `json:"success,omitempty" `

    /*
        索取详情     */
    Model  *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyReceiverequestDetailResultModel) SetSuccess(v bool) *AlibabaAlihealthSynergySyReceiverequestDetailResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthSynergySyReceiverequestDetailResultModel) SetModel(v AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO) *AlibabaAlihealthSynergySyReceiverequestDetailResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyReceiverequestDetailResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyReceiverequestDetailResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyReceiverequestDetailResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyReceiverequestDetailResultModel {
    s.MsgCode = &v
    return s
}
