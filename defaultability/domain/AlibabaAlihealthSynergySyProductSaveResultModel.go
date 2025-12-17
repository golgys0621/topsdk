package domain


type AlibabaAlihealthSynergySyProductSaveResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        产品id     */
    Model  *int64 `json:"model,omitempty" `

    /*
        错误信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        错误码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyProductSaveResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergySyProductSaveResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveResultModel) SetModel(v int64) *AlibabaAlihealthSynergySyProductSaveResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyProductSaveResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductSaveResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyProductSaveResultModel {
    s.MsgCode = &v
    return s
}
