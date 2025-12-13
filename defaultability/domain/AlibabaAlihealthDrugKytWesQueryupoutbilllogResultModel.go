package domain


type AlibabaAlihealthDrugKytWesQueryupoutbilllogResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugKytWesQueryupoutbilllogPageInfoDTO `json:"model,omitempty" `

    /*
        返回值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugKytWesQueryupoutbilllogResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogResultModel) SetModel(v AlibabaAlihealthDrugKytWesQueryupoutbilllogPageInfoDTO) *AlibabaAlihealthDrugKytWesQueryupoutbilllogResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugKytWesQueryupoutbilllogResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugKytWesQueryupoutbilllogResultModel {
    s.MsgCode = &v
    return s
}
