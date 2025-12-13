package domain


type AlibabaAlihealthDrugCodeKytWesDelbillinfoResultModel struct {
    /*
        内层对象     */
    Models  *string `json:"models,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        响应结果     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugCodeKytWesDelbillinfoResultModel) SetModels(v string) *AlibabaAlihealthDrugCodeKytWesDelbillinfoResultModel {
    s.Models = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesDelbillinfoResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugCodeKytWesDelbillinfoResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesDelbillinfoResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugCodeKytWesDelbillinfoResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesDelbillinfoResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugCodeKytWesDelbillinfoResultModel {
    s.ResponseSuccess = &v
    return s
}
