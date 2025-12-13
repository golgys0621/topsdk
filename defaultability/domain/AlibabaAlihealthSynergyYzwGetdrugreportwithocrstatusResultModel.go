package domain


type AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusResultModel struct {
    /*
        是否响应成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        结果     */
    Model  *AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusOnenetTaskDTO `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusResultModel) SetSuccess(v bool) *AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusResultModel) SetModel(v AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusOnenetTaskDTO) *AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusResultModel {
    s.MsgCode = &v
    return s
}
