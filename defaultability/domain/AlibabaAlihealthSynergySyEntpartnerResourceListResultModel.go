package domain


type AlibabaAlihealthSynergySyEntpartnerResourceListResultModel struct {
    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        资料列表     */
    Model  *[]AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyEntpartnerResourceListResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyEntpartnerResourceListResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerResourceListResultModel) SetModel(v []AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel) *AlibabaAlihealthSynergySyEntpartnerResourceListResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerResourceListResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyEntpartnerResourceListResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerResourceListResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyEntpartnerResourceListResultModel {
    s.MsgCode = &v
    return s
}
