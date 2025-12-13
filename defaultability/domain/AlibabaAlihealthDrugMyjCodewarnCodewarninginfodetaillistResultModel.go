package domain


type AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistResultModel struct {
    /*
        返回模型     */
    Model  *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistPageInfoDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistResultModel) SetModel(v AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistPageInfoDto) *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistResultModel {
    s.ResponseSuccess = &v
    return s
}
