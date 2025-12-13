package domain


type AlibabaAlihealthDrugMscGetentinfolistResultModel struct {
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
    Model  *[]AlibabaAlihealthDrugMscGetentinfolistTopEntInfoRespDto `json:"model,omitempty" `

}

func (s *AlibabaAlihealthDrugMscGetentinfolistResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMscGetentinfolistResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfolistResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscGetentinfolistResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfolistResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscGetentinfolistResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfolistResultModel) SetModel(v []AlibabaAlihealthDrugMscGetentinfolistTopEntInfoRespDto) *AlibabaAlihealthDrugMscGetentinfolistResultModel {
    s.Model = &v
    return s
}
