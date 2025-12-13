package domain


type AlibabaAlihealthSynergyYzwQueryreportbycodeResultModel struct {
    /*
        是否响应成功(true |false)     */
    ResSuccess  *bool `json:"res_success,omitempty" `

    /*
        结果     */
    Model  *[]AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeResultModel) SetResSuccess(v bool) *AlibabaAlihealthSynergyYzwQueryreportbycodeResultModel {
    s.ResSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeResultModel) SetModel(v []AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO) *AlibabaAlihealthSynergyYzwQueryreportbycodeResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwQueryreportbycodeResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwQueryreportbycodeResultModel {
    s.MsgCode = &v
    return s
}
