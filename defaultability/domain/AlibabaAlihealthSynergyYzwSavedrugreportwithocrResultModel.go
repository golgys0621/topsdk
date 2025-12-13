package domain


type AlibabaAlihealthSynergyYzwSavedrugreportwithocrResultModel struct {
    /*
        是否响应成功     */
    ResSuccess  *bool `json:"res_success,omitempty" `

    /*
        轮询ID 参考https://open.taobao.com/api.htm?docId=70403&docType=2     */
    Model  *int64 `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwSavedrugreportwithocrResultModel) SetResSuccess(v bool) *AlibabaAlihealthSynergyYzwSavedrugreportwithocrResultModel {
    s.ResSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportwithocrResultModel) SetModel(v int64) *AlibabaAlihealthSynergyYzwSavedrugreportwithocrResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportwithocrResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwSavedrugreportwithocrResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportwithocrResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwSavedrugreportwithocrResultModel {
    s.MsgCode = &v
    return s
}
