package domain


type AlibabaAlihealthSynergySyRefuserequestResultModel struct {
    /*
        是否响应成功(true |false)     */
    ResSuccess  *bool `json:"res_success,omitempty" `

    /*
        是否拒绝成功 true | false     */
    Model  *bool `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyRefuserequestResultModel) SetResSuccess(v bool) *AlibabaAlihealthSynergySyRefuserequestResultModel {
    s.ResSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyRefuserequestResultModel) SetModel(v bool) *AlibabaAlihealthSynergySyRefuserequestResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyRefuserequestResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyRefuserequestResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyRefuserequestResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyRefuserequestResultModel {
    s.MsgCode = &v
    return s
}
