package domain


type AlibabaAlihealthDrugKytWesSearchstatusResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugKytWesSearchstatusPageInfoDto `json:"model,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        响应结果     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesSearchstatusResultModel) SetModel(v AlibabaAlihealthDrugKytWesSearchstatusPageInfoDto) *AlibabaAlihealthDrugKytWesSearchstatusResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugKytWesSearchstatusResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugKytWesSearchstatusResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugKytWesSearchstatusResultModel {
    s.ResponseSuccess = &v
    return s
}
