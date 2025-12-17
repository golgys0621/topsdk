package domain


type AlibabaAlihealthSynergySyOutboxListSyOutboxModel struct {
    /*
        发件箱id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        收件企业名称     */
    ReceiveEntName  *string `json:"receive_ent_name,omitempty" `

    /*
        0企业资料1产品资料     */
    Type  *int64 `json:"type,omitempty" `

    /*
        提交人     */
    SubmitPerson  *string `json:"submit_person,omitempty" `

    /*
        301发送方已撤回,302接收方待查收,303接收方审批中,304接收方全部查收,305接收方部分查收,306接收方全部拒收     */
    ExchangeStatus  *string `json:"exchange_status,omitempty" `

    /*
        全部资料数     */
    TotalResourceCount  *int64 `json:"total_resource_count,omitempty" `

    /*
        对方接收资料数     */
    ReceiveResourceCount  *int64 `json:"receive_resource_count,omitempty" `

    /*
        对方拒收资料数     */
    RejectResourceCount  *int64 `json:"reject_resource_count,omitempty" `

    /*
        发送时间     */
    SendTime  *string `json:"send_time,omitempty" `

    /*
        状态名称     */
    ExchangeStatusDesc  *string `json:"exchange_status_desc,omitempty" `

    /*
        总页数     */
    TotalPage  *int64 `json:"total_page,omitempty" `

    /*
        撤回时间     */
    WithdrawTime  *string `json:"withdraw_time,omitempty" `

    /*
        对方拒收时间     */
    RejectTime  *string `json:"reject_time,omitempty" `

    /*
        对方拒收原因     */
    RejectReason  *string `json:"reject_reason,omitempty" `

    /*
        留言     */
    Note  *string `json:"note,omitempty" `

    /*
        收件企业refentid     */
    ReceiveRefEntId  *string `json:"receive_ref_ent_id,omitempty" `

}

func (s *AlibabaAlihealthSynergySyOutboxListSyOutboxModel) SetId(v int64) *AlibabaAlihealthSynergySyOutboxListSyOutboxModel {
    s.Id = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListSyOutboxModel) SetReceiveEntName(v string) *AlibabaAlihealthSynergySyOutboxListSyOutboxModel {
    s.ReceiveEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListSyOutboxModel) SetType(v int64) *AlibabaAlihealthSynergySyOutboxListSyOutboxModel {
    s.Type = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListSyOutboxModel) SetSubmitPerson(v string) *AlibabaAlihealthSynergySyOutboxListSyOutboxModel {
    s.SubmitPerson = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListSyOutboxModel) SetExchangeStatus(v string) *AlibabaAlihealthSynergySyOutboxListSyOutboxModel {
    s.ExchangeStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListSyOutboxModel) SetTotalResourceCount(v int64) *AlibabaAlihealthSynergySyOutboxListSyOutboxModel {
    s.TotalResourceCount = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListSyOutboxModel) SetReceiveResourceCount(v int64) *AlibabaAlihealthSynergySyOutboxListSyOutboxModel {
    s.ReceiveResourceCount = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListSyOutboxModel) SetRejectResourceCount(v int64) *AlibabaAlihealthSynergySyOutboxListSyOutboxModel {
    s.RejectResourceCount = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListSyOutboxModel) SetSendTime(v string) *AlibabaAlihealthSynergySyOutboxListSyOutboxModel {
    s.SendTime = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListSyOutboxModel) SetExchangeStatusDesc(v string) *AlibabaAlihealthSynergySyOutboxListSyOutboxModel {
    s.ExchangeStatusDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListSyOutboxModel) SetTotalPage(v int64) *AlibabaAlihealthSynergySyOutboxListSyOutboxModel {
    s.TotalPage = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListSyOutboxModel) SetWithdrawTime(v string) *AlibabaAlihealthSynergySyOutboxListSyOutboxModel {
    s.WithdrawTime = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListSyOutboxModel) SetRejectTime(v string) *AlibabaAlihealthSynergySyOutboxListSyOutboxModel {
    s.RejectTime = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListSyOutboxModel) SetRejectReason(v string) *AlibabaAlihealthSynergySyOutboxListSyOutboxModel {
    s.RejectReason = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListSyOutboxModel) SetNote(v string) *AlibabaAlihealthSynergySyOutboxListSyOutboxModel {
    s.Note = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListSyOutboxModel) SetReceiveRefEntId(v string) *AlibabaAlihealthSynergySyOutboxListSyOutboxModel {
    s.ReceiveRefEntId = &v
    return s
}
