package domain


type AlibabaAlihealthSynergyYzwSavedrugreportbyhttpResultModel struct {
    /*
        任务id  ，调用接口查询状态alibaba.alihealth.synergy.yzw.getdrugreportwithhttpstatus     */
    Model  *int64 `json:"model,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    ResponseSuccess  *string `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpResultModel) SetModel(v int64) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpResultModel) SetResponseSuccess(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpResultModel {
    s.ResponseSuccess = &v
    return s
}
