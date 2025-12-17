package domain


type AlibabaAlihealthSynergySyEntpartnerPageResult struct {
    /*
        调用成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        结果     */
    Model  *AlibabaAlihealthSynergySyEntpartnerPagePage `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyEntpartnerPageResult) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyEntpartnerPageResult {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageResult) SetModel(v AlibabaAlihealthSynergySyEntpartnerPagePage) *AlibabaAlihealthSynergySyEntpartnerPageResult {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageResult) SetMsgInfo(v string) *AlibabaAlihealthSynergySyEntpartnerPageResult {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageResult) SetMsgCode(v string) *AlibabaAlihealthSynergySyEntpartnerPageResult {
    s.MsgCode = &v
    return s
}
