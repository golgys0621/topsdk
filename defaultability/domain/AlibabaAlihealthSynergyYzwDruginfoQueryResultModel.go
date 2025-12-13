package domain


type AlibabaAlihealthSynergyYzwDruginfoQueryResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        结果     */
    Model  *[]AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDruginfoQueryResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergyYzwDruginfoQueryResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDruginfoQueryResultModel) SetModel(v []AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO) *AlibabaAlihealthSynergyYzwDruginfoQueryResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDruginfoQueryResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwDruginfoQueryResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDruginfoQueryResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwDruginfoQueryResultModel {
    s.MsgCode = &v
    return s
}
