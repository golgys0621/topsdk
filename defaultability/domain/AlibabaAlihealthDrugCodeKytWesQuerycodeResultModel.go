package domain


type AlibabaAlihealthDrugCodeKytWesQuerycodeResultModel struct {
    /*
        内层大对象     */
    Models  *[]AlibabaAlihealthDrugCodeKytWesQuerycodeCodeFullInfoDto `json:"models,omitempty" `

    /*
        消息码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        消息提示内容     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        查询成功失败标记     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeResultModel) SetModels(v []AlibabaAlihealthDrugCodeKytWesQuerycodeCodeFullInfoDto) *AlibabaAlihealthDrugCodeKytWesQuerycodeResultModel {
    s.Models = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugCodeKytWesQuerycodeResultModel {
    s.ResponseSuccess = &v
    return s
}
