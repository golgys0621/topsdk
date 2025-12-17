package domain


type AlibabaAlihealthSynergySyOutboxDetailResultModel struct {
    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        结果     */
    Model  *AlibabaAlihealthSynergySyOutboxDetailSyOutboxDetailResponse `json:"model,omitempty" `

    /*
        返回值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyOutboxDetailResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyOutboxDetailResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailResultModel) SetModel(v AlibabaAlihealthSynergySyOutboxDetailSyOutboxDetailResponse) *AlibabaAlihealthSynergySyOutboxDetailResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyOutboxDetailResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyOutboxDetailResultModel {
    s.MsgCode = &v
    return s
}
