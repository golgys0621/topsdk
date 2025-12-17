package domain


type AlibabaAlihealthSynergySyResourceResultResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        对象列表     */
    Model  *[]AlibabaAlihealthSynergySyResourceResultOnenetTaskDTO `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyResourceResultResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyResourceResultResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyResourceResultResultModel) SetModel(v []AlibabaAlihealthSynergySyResourceResultOnenetTaskDTO) *AlibabaAlihealthSynergySyResourceResultResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyResourceResultResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyResourceResultResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyResourceResultResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyResourceResultResultModel {
    s.MsgCode = &v
    return s
}
