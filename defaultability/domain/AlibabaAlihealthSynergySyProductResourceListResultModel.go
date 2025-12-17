package domain


type AlibabaAlihealthSynergySyProductResourceListResultModel struct {
    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        资料列表     */
    Model  *[]AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel `json:"model,omitempty" `

    /*
        错误值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        错误码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyProductResourceListResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyProductResourceListResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductResourceListResultModel) SetModel(v []AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel) *AlibabaAlihealthSynergySyProductResourceListResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductResourceListResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyProductResourceListResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductResourceListResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyProductResourceListResultModel {
    s.MsgCode = &v
    return s
}
