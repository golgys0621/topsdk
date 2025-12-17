package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyInboxListSyInboxModel struct {
	/*
	   id     */
	Id *int64 `json:"id,omitempty" `

	/*
	   创建时间     */
	GmtCreate *util.LocalTime `json:"gmt_create,omitempty" `

	/*
	   更新时间     */
	GmtModified *util.LocalTime `json:"gmt_modified,omitempty" `

	/*
	   发送企业RefEntId     */
	SendRefEntId *string `json:"send_ref_ent_id,omitempty" `

	/*
	   发送企业entId     */
	SendEntId *string `json:"send_ent_id,omitempty" `

	/*
	   发送企业名称     */
	SendEntName *string `json:"send_ent_name,omitempty" `

	/*
	   发件人     */
	SendPerson *string `json:"send_person,omitempty" `

	/*
	   发件人id     */
	SendPersonId *string `json:"send_person_id,omitempty" `

	/*
	   0企业资料1产品资料     */
	Type *int64 `json:"type,omitempty" `

	/*
	   发送时间     */
	SendTime *util.LocalTime `json:"send_time,omitempty" `

	/*
	   301发送方已撤回,302接收方待查收,303接收方审批中,304接收方全部查收,305接收方部分查收,306接收方全部拒收     */
	ExchangeStatus *int64 `json:"exchange_status,omitempty" `

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
	   拒绝的资料数     */
	RejectResourceCount *int64 `json:"reject_resource_count,omitempty" `

	/*
	   留言     */
	Note *string `json:"note,omitempty" `

	/*
	   撤回人     */
	WithdrawPerson *string `json:"withdraw_person,omitempty" `

	/*
	   撤回人id     */
	WithdrawPersonId *string `json:"withdraw_person_id,omitempty" `

	/*
	   撤回时间     */
	WithdrawTime *util.LocalTime `json:"withdraw_time,omitempty" `

	/*
	   拒收原因     */
	RejectReason *string `json:"reject_reason,omitempty" `

	/*
	   拒收时间     */
	RejectTime *util.LocalTime `json:"reject_time,omitempty" `

	/*
	   0拒收，1全部查收，2部分查收     */
	CheckStatus *int64 `json:"check_status,omitempty" `

	/*
	   0拒收，1全部查收，2部分查收     */
	CheckStatusDesc *string `json:"check_status_desc,omitempty" `
}

func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetId(v int64) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.Id = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetGmtCreate(v util.LocalTime) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.GmtCreate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetGmtModified(v util.LocalTime) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.GmtModified = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetSendRefEntId(v string) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.SendRefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetSendEntId(v string) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.SendEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetSendEntName(v string) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.SendEntName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetSendPerson(v string) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.SendPerson = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetSendPersonId(v string) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.SendPersonId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetType(v int64) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.Type = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetSendTime(v util.LocalTime) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.SendTime = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetExchangeStatus(v int64) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.ExchangeStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetExchangeStatusDesc(v string) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.ExchangeStatusDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetTotalResourceCount(v int64) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.TotalResourceCount = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetReceiveResourceCount(v int64) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.ReceiveResourceCount = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetRejectResourceCount(v int64) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.RejectResourceCount = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetNote(v string) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.Note = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetWithdrawPerson(v string) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.WithdrawPerson = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetWithdrawPersonId(v string) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.WithdrawPersonId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetWithdrawTime(v util.LocalTime) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.WithdrawTime = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetRejectReason(v string) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.RejectReason = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetRejectTime(v util.LocalTime) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.RejectTime = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetCheckStatus(v int64) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.CheckStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxListSyInboxModel) SetCheckStatusDesc(v string) *AlibabaAlihealthSynergySyInboxListSyInboxModel {
	s.CheckStatusDesc = &v
	return s
}
