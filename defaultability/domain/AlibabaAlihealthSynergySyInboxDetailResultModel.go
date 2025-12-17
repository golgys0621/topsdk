package domain


type AlibabaAlihealthSynergySyInboxDetailResultModel struct {
    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        结果     */
    Model  *AlibabaAlihealthSynergySyInboxDetailSyInboxDetailResponse `json:"model,omitempty" `

    /*
        返回值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyInboxDetailResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyInboxDetailResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailResultModel) SetModel(v AlibabaAlihealthSynergySyInboxDetailSyInboxDetailResponse) *AlibabaAlihealthSynergySyInboxDetailResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyInboxDetailResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyInboxDetailResultModel {
    s.MsgCode = &v
    return s
}
