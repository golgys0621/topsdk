package domain


type AlibabaAlihealthEntYzwHasuploaddrugreportResultModel struct {
    /*
        true表示上传过，false表示没上传过     */
    Model  *bool `json:"model,omitempty" `

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

func (s *AlibabaAlihealthEntYzwHasuploaddrugreportResultModel) SetModel(v bool) *AlibabaAlihealthEntYzwHasuploaddrugreportResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthEntYzwHasuploaddrugreportResultModel) SetMsgCode(v string) *AlibabaAlihealthEntYzwHasuploaddrugreportResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthEntYzwHasuploaddrugreportResultModel) SetMsgInfo(v string) *AlibabaAlihealthEntYzwHasuploaddrugreportResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthEntYzwHasuploaddrugreportResultModel) SetResponseSuccess(v string) *AlibabaAlihealthEntYzwHasuploaddrugreportResultModel {
    s.ResponseSuccess = &v
    return s
}
