package domain


type AlibabaAlihealthDrugKytWesGetentinfoResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugKytWesGetentinfoPUserEntInfoDto `json:"model,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        返回信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesGetentinfoResultModel) SetModel(v AlibabaAlihealthDrugKytWesGetentinfoPUserEntInfoDto) *AlibabaAlihealthDrugKytWesGetentinfoResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetentinfoResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugKytWesGetentinfoResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetentinfoResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugKytWesGetentinfoResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetentinfoResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugKytWesGetentinfoResultModel {
    s.ResponseSuccess = &v
    return s
}
