package domain


type AlibabaAlihealthEntYzwGetdrugidbyprodcodeResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        药品id     */
    Model  *string `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthEntYzwGetdrugidbyprodcodeResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthEntYzwGetdrugidbyprodcodeResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthEntYzwGetdrugidbyprodcodeResultModel) SetModel(v string) *AlibabaAlihealthEntYzwGetdrugidbyprodcodeResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthEntYzwGetdrugidbyprodcodeResultModel) SetMsgInfo(v string) *AlibabaAlihealthEntYzwGetdrugidbyprodcodeResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthEntYzwGetdrugidbyprodcodeResultModel) SetMsgCode(v string) *AlibabaAlihealthEntYzwGetdrugidbyprodcodeResultModel {
    s.MsgCode = &v
    return s
}
