package domain


type AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoResultModel struct {
    /*
        model     */
    ModelList  *[]AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO `json:"model_list,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        返回信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    Success  *string `json:"success,omitempty" `

}

func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoResultModel) SetModelList(v []AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoResultModel {
    s.ModelList = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoResultModel) SetSuccess(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoResultModel {
    s.Success = &v
    return s
}
