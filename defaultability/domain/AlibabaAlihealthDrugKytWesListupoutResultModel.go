package domain


type AlibabaAlihealthDrugKytWesListupoutResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugKytWesListupoutPageInfoDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugKytWesListupoutResultModel) SetModel(v AlibabaAlihealthDrugKytWesListupoutPageInfoDto) *AlibabaAlihealthDrugKytWesListupoutResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugKytWesListupoutResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugKytWesListupoutResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugKytWesListupoutResultModel {
    s.ResponseSuccess = &v
    return s
}
