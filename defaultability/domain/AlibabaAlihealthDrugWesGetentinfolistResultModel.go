package domain


type AlibabaAlihealthDrugWesGetentinfolistResultModel struct {
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
    Model  *[]AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto `json:"model,omitempty" `

}

func (s *AlibabaAlihealthDrugWesGetentinfolistResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugWesGetentinfolistResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfolistResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugWesGetentinfolistResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfolistResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugWesGetentinfolistResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfolistResultModel) SetModel(v []AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto) *AlibabaAlihealthDrugWesGetentinfolistResultModel {
    s.Model = &v
    return s
}
