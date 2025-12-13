package domain


type AlibabaAlihealthDrugYljgSearchbillResultModel struct {
    /*
        返回模型     */
    Model  *AlibabaAlihealthDrugYljgSearchbillPageInfoDTO `json:"model,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        返回信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugYljgSearchbillResultModel) SetModel(v AlibabaAlihealthDrugYljgSearchbillPageInfoDTO) *AlibabaAlihealthDrugYljgSearchbillResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugYljgSearchbillResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugYljgSearchbillResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugYljgSearchbillResultModel {
    s.ResponseSuccess = &v
    return s
}
