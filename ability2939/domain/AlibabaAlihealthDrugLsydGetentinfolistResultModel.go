package domain


type AlibabaAlihealthDrugLsydGetentinfolistResultModel struct {
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
    Model  *[]AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto `json:"model,omitempty" `

}

func (s *AlibabaAlihealthDrugLsydGetentinfolistResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugLsydGetentinfolistResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugLsydGetentinfolistResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugLsydGetentinfolistResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistResultModel) SetModel(v []AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto) *AlibabaAlihealthDrugLsydGetentinfolistResultModel {
    s.Model = &v
    return s
}
