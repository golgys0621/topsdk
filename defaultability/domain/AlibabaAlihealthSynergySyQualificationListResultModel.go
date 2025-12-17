package domain


type AlibabaAlihealthSynergySyQualificationListResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        返回对象     */
    Model  *[]AlibabaAlihealthSynergySyQualificationListSyQualificationInfoResponse `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态值     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyQualificationListResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyQualificationListResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyQualificationListResultModel) SetModel(v []AlibabaAlihealthSynergySyQualificationListSyQualificationInfoResponse) *AlibabaAlihealthSynergySyQualificationListResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyQualificationListResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyQualificationListResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyQualificationListResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyQualificationListResultModel {
    s.MsgCode = &v
    return s
}
