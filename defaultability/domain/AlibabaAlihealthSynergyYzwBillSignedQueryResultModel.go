package domain


type AlibabaAlihealthSynergyYzwBillSignedQueryResultModel struct {
    /*
        是否响应成功(true |false)     */
    Success  *bool `json:"success,omitempty" `

    /*
        结果     */
    Model  *AlibabaAlihealthSynergyYzwBillSignedQueryPage `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwBillSignedQueryResultModel) SetSuccess(v bool) *AlibabaAlihealthSynergyYzwBillSignedQueryResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryResultModel) SetModel(v AlibabaAlihealthSynergyYzwBillSignedQueryPage) *AlibabaAlihealthSynergyYzwBillSignedQueryResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryResultModel {
    s.MsgCode = &v
    return s
}
