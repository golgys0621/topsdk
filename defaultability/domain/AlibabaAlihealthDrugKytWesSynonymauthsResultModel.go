package domain


type AlibabaAlihealthDrugKytWesSynonymauthsResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugKytWesSynonymauthsPage `json:"model,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesSynonymauthsResultModel) SetModel(v AlibabaAlihealthDrugKytWesSynonymauthsPage) *AlibabaAlihealthDrugKytWesSynonymauthsResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSynonymauthsResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugKytWesSynonymauthsResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSynonymauthsResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugKytWesSynonymauthsResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSynonymauthsResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugKytWesSynonymauthsResultModel {
    s.ResponseSuccess = &v
    return s
}
