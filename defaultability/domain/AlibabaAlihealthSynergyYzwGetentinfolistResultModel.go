package domain


type AlibabaAlihealthSynergyYzwGetentinfolistResultModel struct {
    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        错误编码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        错误信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        响应结果     */
    Model  *[]AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto `json:"model,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwGetentinfolistResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergyYzwGetentinfolistResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwGetentinfolistResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwGetentinfolistResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistResultModel) SetModel(v []AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto) *AlibabaAlihealthSynergyYzwGetentinfolistResultModel {
    s.Model = &v
    return s
}
