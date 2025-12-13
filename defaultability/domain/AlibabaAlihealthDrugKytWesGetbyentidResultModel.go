package domain


type AlibabaAlihealthDrugKytWesGetbyentidResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugKytWesGetbyentidResultModel) SetModel(v AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) *AlibabaAlihealthDrugKytWesGetbyentidResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugKytWesGetbyentidResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugKytWesGetbyentidResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugKytWesGetbyentidResultModel {
    s.ResponseSuccess = &v
    return s
}
