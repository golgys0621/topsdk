package domain


type AlibabaAlihealthDrugMscListpartByagentResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugMscListpartByagentPage `json:"model,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        返回值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugMscListpartByagentResultModel) SetModel(v AlibabaAlihealthDrugMscListpartByagentPage) *AlibabaAlihealthDrugMscListpartByagentResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscListpartByagentResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscListpartByagentResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMscListpartByagentResultModel {
    s.ResponseSuccess = &v
    return s
}
