package domain


type AlibabaAlihealthSynergyYzwListsealsResultModel struct {
    /*
        结果     */
    Model  *[]AlibabaAlihealthSynergyYzwListsealsCASealDTO `json:"model,omitempty" `

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

func (s *AlibabaAlihealthSynergyYzwListsealsResultModel) SetModel(v []AlibabaAlihealthSynergyYzwListsealsCASealDTO) *AlibabaAlihealthSynergyYzwListsealsResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwListsealsResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwListsealsResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwListsealsResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwListsealsResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwListsealsResultModel) SetResponseSuccess(v string) *AlibabaAlihealthSynergyYzwListsealsResultModel {
    s.ResponseSuccess = &v
    return s
}
