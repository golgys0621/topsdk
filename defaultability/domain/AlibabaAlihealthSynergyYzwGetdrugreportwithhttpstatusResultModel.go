package domain


type AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusResultModel struct {
    /*
        是否响应成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        结果     */
    Model  *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusOnenetTaskDTO `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusResultModel) SetSuccess(v bool) *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusResultModel) SetModel(v AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusOnenetTaskDTO) *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusResultModel {
    s.MsgCode = &v
    return s
}
