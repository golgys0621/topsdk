package domain


type AlibabaAlihealthSynergySySignatureDetailGetResultModel struct {
    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        签章对象     */
    Model  *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetailResponse `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySySignatureDetailGetResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySySignatureDetailGetResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetResultModel) SetModel(v AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetailResponse) *AlibabaAlihealthSynergySySignatureDetailGetResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySySignatureDetailGetResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySySignatureDetailGetResultModel {
    s.MsgCode = &v
    return s
}
