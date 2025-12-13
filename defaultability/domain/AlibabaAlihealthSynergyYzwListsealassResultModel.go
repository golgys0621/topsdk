package domain


type AlibabaAlihealthSynergyYzwListsealassResultModel struct {
    /*
        结果     */
    Model  *[]AlibabaAlihealthSynergyYzwListsealassCASealDTO `json:"model,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    ResponseSuccess  *string `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwListsealassResultModel) SetModel(v []AlibabaAlihealthSynergyYzwListsealassCASealDTO) *AlibabaAlihealthSynergyYzwListsealassResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwListsealassResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwListsealassResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwListsealassResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwListsealassResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwListsealassResultModel) SetResponseSuccess(v string) *AlibabaAlihealthSynergyYzwListsealassResultModel {
    s.ResponseSuccess = &v
    return s
}
