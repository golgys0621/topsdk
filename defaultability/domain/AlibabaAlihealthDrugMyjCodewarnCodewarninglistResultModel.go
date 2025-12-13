package domain


type AlibabaAlihealthDrugMyjCodewarnCodewarninglistResultModel struct {
    /*
        返回模型     */
    Model  *AlibabaAlihealthDrugMyjCodewarnCodewarninglistPageInfoDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistResultModel) SetModel(v AlibabaAlihealthDrugMyjCodewarnCodewarninglistPageInfoDto) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistResultModel {
    s.ResponseSuccess = &v
    return s
}
