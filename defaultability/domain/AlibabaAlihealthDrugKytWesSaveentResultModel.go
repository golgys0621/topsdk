package domain


type AlibabaAlihealthDrugKytWesSaveentResultModel struct {
    /*
        true：接口调用成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        具体返回值     */
    Model  *AlibabaAlihealthDrugKytWesSaveentModel `json:"model,omitempty" `

    /*
        接口调用失败具体信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        接口调用失败具体code     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesSaveentResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugKytWesSaveentResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSaveentResultModel) SetModel(v AlibabaAlihealthDrugKytWesSaveentModel) *AlibabaAlihealthDrugKytWesSaveentResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSaveentResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugKytWesSaveentResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSaveentResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugKytWesSaveentResultModel {
    s.MsgCode = &v
    return s
}
