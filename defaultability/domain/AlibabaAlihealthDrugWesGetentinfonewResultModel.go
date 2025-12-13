package domain


type AlibabaAlihealthDrugWesGetentinfonewResultModel struct {
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
    Model  *AlibabaAlihealthDrugWesGetentinfonewTopEntInfoRespDto `json:"model,omitempty" `

}

func (s *AlibabaAlihealthDrugWesGetentinfonewResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugWesGetentinfonewResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfonewResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugWesGetentinfonewResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfonewResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugWesGetentinfonewResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfonewResultModel) SetModel(v AlibabaAlihealthDrugWesGetentinfonewTopEntInfoRespDto) *AlibabaAlihealthDrugWesGetentinfonewResultModel {
    s.Model = &v
    return s
}
