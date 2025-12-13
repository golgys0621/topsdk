package domain


type AlibabaAlihealthDrugKytWesGetbyrefentidResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugKytWesGetbyrefentidResultModel) SetModel(v AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto) *AlibabaAlihealthDrugKytWesGetbyrefentidResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugKytWesGetbyrefentidResultModel {
    s.ResponseSuccess = &v
    return s
}
