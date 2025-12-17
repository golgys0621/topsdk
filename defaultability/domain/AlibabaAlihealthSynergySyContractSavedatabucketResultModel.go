package domain


type AlibabaAlihealthSynergySyContractSavedatabucketResultModel struct {
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

func (s *AlibabaAlihealthSynergySyContractSavedatabucketResultModel) SetModel(v string) *AlibabaAlihealthSynergySyContractSavedatabucketResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractSavedatabucketResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyContractSavedatabucketResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractSavedatabucketResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyContractSavedatabucketResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractSavedatabucketResultModel) SetResponseSuccess(v string) *AlibabaAlihealthSynergySyContractSavedatabucketResultModel {
    s.ResponseSuccess = &v
    return s
}
