package domain


type AlibabaAlihealthEntYzwAuthGetdruginfoResultModel struct {
    /*
        是否响应成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        查询对象     */
    Model  *AlibabaAlihealthEntYzwAuthGetdruginfoModel `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthEntYzwAuthGetdruginfoResultModel) SetSuccess(v bool) *AlibabaAlihealthEntYzwAuthGetdruginfoResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthEntYzwAuthGetdruginfoResultModel) SetModel(v AlibabaAlihealthEntYzwAuthGetdruginfoModel) *AlibabaAlihealthEntYzwAuthGetdruginfoResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthEntYzwAuthGetdruginfoResultModel) SetMsgInfo(v string) *AlibabaAlihealthEntYzwAuthGetdruginfoResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthEntYzwAuthGetdruginfoResultModel) SetMsgCode(v string) *AlibabaAlihealthEntYzwAuthGetdruginfoResultModel {
    s.MsgCode = &v
    return s
}
