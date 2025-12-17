package domain


type AlibabaAlihealthSynergySySignatureResourceSealResultModel struct {
    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        签章资料id     */
    Model  *int64 `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySySignatureResourceSealResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySySignatureResourceSealResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureResourceSealResultModel) SetModel(v int64) *AlibabaAlihealthSynergySySignatureResourceSealResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureResourceSealResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySySignatureResourceSealResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureResourceSealResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySySignatureResourceSealResultModel {
    s.MsgCode = &v
    return s
}
