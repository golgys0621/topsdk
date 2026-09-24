package domain


type AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistResultModel struct {
    /*
        返回模型     */
    Model  *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistPageInfoDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistResultModel) SetModel(v AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistPageInfoDto) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistResultModel {
    s.ResponseSuccess = &v
    return s
}
