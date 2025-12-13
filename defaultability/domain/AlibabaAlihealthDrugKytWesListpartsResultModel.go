package domain


type AlibabaAlihealthDrugKytWesListpartsResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugKytWesListpartsPage `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugKytWesListpartsResultModel) SetModel(v AlibabaAlihealthDrugKytWesListpartsPage) *AlibabaAlihealthDrugKytWesListpartsResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugKytWesListpartsResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugKytWesListpartsResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugKytWesListpartsResultModel {
    s.ResponseSuccess = &v
    return s
}
