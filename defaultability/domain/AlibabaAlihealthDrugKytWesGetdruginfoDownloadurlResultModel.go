package domain


type AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel struct {
    /*
        true：接口调用成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        返回值     */
    Model  *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlJSONObject `json:"model,omitempty" `

    /*
        接口调用失败具体信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        接口调用失败具体code     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel) SetModel(v AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlJSONObject) *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel {
    s.MsgCode = &v
    return s
}
