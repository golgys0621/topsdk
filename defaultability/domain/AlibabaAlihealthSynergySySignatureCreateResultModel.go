package domain


type AlibabaAlihealthSynergySySignatureCreateResultModel struct {
    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        签章id     */
    Model  *int64 `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySySignatureCreateResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySySignatureCreateResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureCreateResultModel) SetModel(v int64) *AlibabaAlihealthSynergySySignatureCreateResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureCreateResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySySignatureCreateResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureCreateResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySySignatureCreateResultModel {
    s.MsgCode = &v
    return s
}
