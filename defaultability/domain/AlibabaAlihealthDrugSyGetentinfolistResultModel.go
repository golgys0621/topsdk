package domain


type AlibabaAlihealthDrugSyGetentinfolistResultModel struct {
    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        企业列表     */
    Model  *[]AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto `json:"model,omitempty" `

    /*
        错误信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        错误编码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugSyGetentinfolistResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugSyGetentinfolistResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugSyGetentinfolistResultModel) SetModel(v []AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto) *AlibabaAlihealthDrugSyGetentinfolistResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugSyGetentinfolistResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugSyGetentinfolistResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugSyGetentinfolistResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugSyGetentinfolistResultModel {
    s.MsgCode = &v
    return s
}
