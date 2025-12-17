package domain


type AlibabaAlihealthSynergySyListrequestdetailResultModel struct {
    /*
        success     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        model     */
    Model  *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO `json:"model,omitempty" `

    /*
        msgInfo     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        msgCode     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyListrequestdetailResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyListrequestdetailResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailResultModel) SetModel(v AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO) *AlibabaAlihealthSynergySyListrequestdetailResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyListrequestdetailResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyListrequestdetailResultModel {
    s.MsgCode = &v
    return s
}
