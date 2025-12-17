package domain


type AlibabaAlihealthSynergySyEntResourceListResultModel struct {
    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        资料列表     */
    Model  *[]AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyEntResourceListResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyEntResourceListResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntResourceListResultModel) SetModel(v []AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel) *AlibabaAlihealthSynergySyEntResourceListResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntResourceListResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyEntResourceListResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntResourceListResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyEntResourceListResultModel {
    s.MsgCode = &v
    return s
}
