package domain


type AlibabaAlihealthSynergySyListrequestResultModel struct {
    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        结果     */
    Model  *AlibabaAlihealthSynergySyListrequestPage `json:"model,omitempty" `

    /*
        返回值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyListrequestResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyListrequestResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListrequestResultModel) SetModel(v AlibabaAlihealthSynergySyListrequestPage) *AlibabaAlihealthSynergySyListrequestResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListrequestResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyListrequestResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListrequestResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyListrequestResultModel {
    s.MsgCode = &v
    return s
}
