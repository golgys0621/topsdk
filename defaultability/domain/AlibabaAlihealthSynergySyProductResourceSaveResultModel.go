package domain


type AlibabaAlihealthSynergySyProductResourceSaveResultModel struct {
    /*
        是否成功     */
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

func (s *AlibabaAlihealthSynergySyProductResourceSaveResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyProductResourceSaveResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductResourceSaveResultModel) SetModel(v int64) *AlibabaAlihealthSynergySyProductResourceSaveResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductResourceSaveResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyProductResourceSaveResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductResourceSaveResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyProductResourceSaveResultModel {
    s.MsgCode = &v
    return s
}
