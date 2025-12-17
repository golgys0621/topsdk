package domain


type AlibabaAlihealthSynergySyEntResourceSaveResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        任务id     */
    Model  *int64 `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyEntResourceSaveResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyEntResourceSaveResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntResourceSaveResultModel) SetModel(v int64) *AlibabaAlihealthSynergySyEntResourceSaveResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntResourceSaveResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyEntResourceSaveResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntResourceSaveResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyEntResourceSaveResultModel {
    s.MsgCode = &v
    return s
}
