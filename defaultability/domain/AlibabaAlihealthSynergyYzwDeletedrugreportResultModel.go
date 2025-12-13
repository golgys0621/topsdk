package domain


type AlibabaAlihealthSynergyYzwDeletedrugreportResultModel struct {
    /*
        是否删除成功:true是删除成功,false是删除失败     */
    Model  *bool `json:"model,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDeletedrugreportResultModel) SetModel(v bool) *AlibabaAlihealthSynergyYzwDeletedrugreportResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDeletedrugreportResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwDeletedrugreportResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDeletedrugreportResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwDeletedrugreportResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDeletedrugreportResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthSynergyYzwDeletedrugreportResultModel {
    s.ResponseSuccess = &v
    return s
}
