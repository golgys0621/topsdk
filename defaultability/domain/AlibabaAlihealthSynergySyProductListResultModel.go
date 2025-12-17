package domain


type AlibabaAlihealthSynergySyProductListResultModel struct {
    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        结果     */
    Model  *AlibabaAlihealthSynergySyProductListPage `json:"model,omitempty" `

    /*
        返回值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyProductListResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyProductListResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListResultModel) SetModel(v AlibabaAlihealthSynergySyProductListPage) *AlibabaAlihealthSynergySyProductListResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyProductListResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyProductListResultModel {
    s.MsgCode = &v
    return s
}
