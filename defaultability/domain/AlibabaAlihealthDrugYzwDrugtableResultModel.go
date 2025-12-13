package domain


type AlibabaAlihealthDrugYzwDrugtableResultModel struct {
    /*
        是否响应成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugYzwDrugtableModel `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugYzwDrugtableResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugYzwDrugtableResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthDrugYzwDrugtableResultModel) SetModel(v AlibabaAlihealthDrugYzwDrugtableModel) *AlibabaAlihealthDrugYzwDrugtableResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugYzwDrugtableResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugYzwDrugtableResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugYzwDrugtableResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugYzwDrugtableResultModel {
    s.MsgCode = &v
    return s
}
