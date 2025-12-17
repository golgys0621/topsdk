package domain


type AlibabaAlihealthSynergySyListsealsResultModel struct {
    /*
        结果     */
    Model  *[]AlibabaAlihealthSynergySyListsealsCASealDTO `json:"model,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    ResponseSuccess  *string `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthSynergySyListsealsResultModel) SetModel(v []AlibabaAlihealthSynergySyListsealsCASealDTO) *AlibabaAlihealthSynergySyListsealsResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListsealsResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyListsealsResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListsealsResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyListsealsResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListsealsResultModel) SetResponseSuccess(v string) *AlibabaAlihealthSynergySyListsealsResultModel {
    s.ResponseSuccess = &v
    return s
}
