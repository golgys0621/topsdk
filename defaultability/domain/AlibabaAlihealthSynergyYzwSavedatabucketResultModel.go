package domain


type AlibabaAlihealthSynergyYzwSavedatabucketResultModel struct {
    /*
        文件标识     */
    Model  *string `json:"model,omitempty" `

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

func (s *AlibabaAlihealthSynergyYzwSavedatabucketResultModel) SetModel(v string) *AlibabaAlihealthSynergyYzwSavedatabucketResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedatabucketResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwSavedatabucketResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedatabucketResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwSavedatabucketResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedatabucketResultModel) SetResponseSuccess(v string) *AlibabaAlihealthSynergyYzwSavedatabucketResultModel {
    s.ResponseSuccess = &v
    return s
}
