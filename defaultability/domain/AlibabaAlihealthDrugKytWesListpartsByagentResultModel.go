package domain


type AlibabaAlihealthDrugKytWesListpartsByagentResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugKytWesListpartsByagentPage `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugKytWesListpartsByagentResultModel) SetModel(v AlibabaAlihealthDrugKytWesListpartsByagentPage) *AlibabaAlihealthDrugKytWesListpartsByagentResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugKytWesListpartsByagentResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugKytWesListpartsByagentResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugKytWesListpartsByagentResultModel {
    s.ResponseSuccess = &v
    return s
}
