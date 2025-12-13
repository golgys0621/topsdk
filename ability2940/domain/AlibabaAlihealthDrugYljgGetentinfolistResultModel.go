package domain


type AlibabaAlihealthDrugYljgGetentinfolistResultModel struct {
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
    Model  *[]AlibabaAlihealthDrugYljgGetentinfolistTopEntInfoRespDto `json:"model,omitempty" `

}

func (s *AlibabaAlihealthDrugYljgGetentinfolistResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugYljgGetentinfolistResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgGetentinfolistResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugYljgGetentinfolistResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgGetentinfolistResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugYljgGetentinfolistResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgGetentinfolistResultModel) SetModel(v []AlibabaAlihealthDrugYljgGetentinfolistTopEntInfoRespDto) *AlibabaAlihealthDrugYljgGetentinfolistResultModel {
    s.Model = &v
    return s
}
