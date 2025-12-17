package domain


type AlibabaAlihealthSynergySyListreceiverequestResultModel struct {
    /*
        是否响应成功(true |false)     */
    ResSuccess  *bool `json:"res_success,omitempty" `

    /*
        结果     */
    Model  *AlibabaAlihealthSynergySyListreceiverequestPage `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyListreceiverequestResultModel) SetResSuccess(v bool) *AlibabaAlihealthSynergySyListreceiverequestResultModel {
    s.ResSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestResultModel) SetModel(v AlibabaAlihealthSynergySyListreceiverequestPage) *AlibabaAlihealthSynergySyListreceiverequestResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyListreceiverequestResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyListreceiverequestResultModel {
    s.MsgCode = &v
    return s
}
