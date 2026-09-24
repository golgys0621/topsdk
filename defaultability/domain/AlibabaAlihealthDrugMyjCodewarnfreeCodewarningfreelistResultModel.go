package domain


type AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistResultModel struct {
    /*
        返回模型     */
    Model  *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistPageInfoDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistResultModel) SetModel(v AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistPageInfoDto) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistResultModel {
    s.ResponseSuccess = &v
    return s
}
