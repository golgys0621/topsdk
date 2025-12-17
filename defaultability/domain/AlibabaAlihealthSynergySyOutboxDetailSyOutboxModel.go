package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel struct {
	/*
	   发件箱id     */
	Id *int64 `json:"id,omitempty" `

	/*
	   创建时间     */
	GmtCreate *util.LocalTime `json:"gmt_create,omitempty" `

	/*
	   更新时间     */
	GmtModified *util.LocalTime `json:"gmt_modified,omitempty" `

	/*
	   本企业refentid     */
	RefEntId *string `json:"ref_ent_id,omitempty" `

	/*
	   接收企业refentid     */
	ReceiveRefEntId *string `json:"receive_ref_ent_id,omitempty" `

	/*
	   接收企业名称     */
	ReceiveEntName *string `json:"receive_ent_name,omitempty" `

	/*
	   发件人     */
	SendPerson *string `json:"send_person,omitempty" `

	/*
	   发件人id     */
	SendPersonId *string `json:"send_person_id,omitempty" `

	/*
	   资料类型，0企业资料1产品资料     */
	Type *int64 `json:"type,omitempty" `

	/*
	   发送时间     */
	SendTime *util.LocalTime `json:"send_time,omitempty" `

	/*
	   301发送方已撤回,302接收方待查收,303接收方审批中,304接收方全部查收,305接收方部分查收,306接收方全部拒收     */
	ExchangeStatus *string `json:"exchange_status,omitempty" `

	/*
	   状态描述     */
	ExchangeStatusDesc *string `json:"exchange_status_desc,omitempty" `

	/*
	   资料数     */
	TotalResourceCount *int64 `json:"total_resource_count,omitempty" `

	/*
	   通过的资料数     */
	ReceiveResourceCount *int64 `json:"receive_resource_count,omitempty" `

	/*
	   拒收的资料数     */
	RejectResourceCount *int64 `json:"reject_resource_count,omitempty" `

	/*
	   留言     */
	Note *string `json:"note,omitempty" `

	/*
	   撤回人     */
	WithdrawPerson *string `json:"withdraw_person,omitempty" `

	/*
	   撤回时间     */
	WithdrawTime *util.LocalTime `json:"withdraw_time,omitempty" `

	/*
	   拒收原因     */
	RejectReason *string `json:"reject_reason,omitempty" `

	/*
	   拒收时间     */
	RejectTime *util.LocalTime `json:"reject_time,omitempty" `
}

func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetId(v int64) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.Id = &v
	return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetGmtCreate(v util.LocalTime) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.GmtCreate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetGmtModified(v util.LocalTime) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.GmtModified = &v
	return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetRefEntId(v string) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetReceiveRefEntId(v string) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.ReceiveRefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetReceiveEntName(v string) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.ReceiveEntName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetSendPerson(v string) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.SendPerson = &v
	return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetSendPersonId(v string) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.SendPersonId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetType(v int64) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.Type = &v
	return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetSendTime(v util.LocalTime) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.SendTime = &v
	return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetExchangeStatus(v string) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.ExchangeStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetExchangeStatusDesc(v string) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.ExchangeStatusDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetTotalResourceCount(v int64) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.TotalResourceCount = &v
	return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetReceiveResourceCount(v int64) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.ReceiveResourceCount = &v
	return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetRejectResourceCount(v int64) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.RejectResourceCount = &v
	return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetNote(v string) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.Note = &v
	return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetWithdrawPerson(v string) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.WithdrawPerson = &v
	return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetWithdrawTime(v util.LocalTime) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.WithdrawTime = &v
	return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetRejectReason(v string) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.RejectReason = &v
	return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) SetRejectTime(v util.LocalTime) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel {
	s.RejectTime = &v
	return s
}
