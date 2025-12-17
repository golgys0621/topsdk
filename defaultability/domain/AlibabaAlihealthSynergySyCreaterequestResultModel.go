package domain


type AlibabaAlihealthSynergySyCreaterequestResultModel struct {
    /*
        是否响应成功(true |false)     */
    ResSuccess  *bool `json:"res_success,omitempty" `

    /*
        异步任务ID，再调用alibaba.alihealth.synergy.sy.resource.result接口查看任务状态, 及被索取ID:syReceiveId     */
    Model  *int64 `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyCreaterequestResultModel) SetResSuccess(v bool) *AlibabaAlihealthSynergySyCreaterequestResultModel {
    s.ResSuccess = &v
    return s
}
func (s *AlibabaAlihealthSynergySyCreaterequestResultModel) SetModel(v int64) *AlibabaAlihealthSynergySyCreaterequestResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthSynergySyCreaterequestResultModel) SetMsgInfo(v string) *AlibabaAlihealthSynergySyCreaterequestResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyCreaterequestResultModel) SetMsgCode(v string) *AlibabaAlihealthSynergySyCreaterequestResultModel {
    s.MsgCode = &v
    return s
}
