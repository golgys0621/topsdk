package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyInboxDetailSyInboxModel struct {
	/*
	   收件箱id     */
	Id *int64 `json:"id,omitempty" `

	/*
	   创建时间     */
	GmtCreate *util.LocalTime `json:"gmt_create,omitempty" `

	/*
	   更新时间     */
	GmtModified *util.LocalTime `json:"gmt_modified,omitempty" `

	/*
	   本企业refEntId     */
	RefEntId *string `json:"ref_ent_id,omitempty" `

	/*
	   发送企业RefEntId     */
	SendRefEntId *string `json:"send_ref_ent_id,omitempty" `

	/*
	   发送企业EntId     */
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
	   资料类型，0企业资料1产品资料     */
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
	   拒收的资料数     */
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
	   查收状态 0拒收，1全部查收，2部分查收     */
	CheckStatus *int64 `json:"check_status,omitempty" `

	/*
	   查收状态     */
	CheckStatusDesc *string `json:"check_status_desc,omitempty" `

	/*
	   索取id     */
	RequestId *int64 `json:"request_id,omitempty" `
}

func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetId(v int64) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.Id = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetGmtCreate(v util.LocalTime) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.GmtCreate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetGmtModified(v util.LocalTime) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.GmtModified = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetRefEntId(v string) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetSendRefEntId(v string) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.SendRefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetSendEntId(v string) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.SendEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetSendEntName(v string) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.SendEntName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetSendPerson(v string) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.SendPerson = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetSendPersonId(v string) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.SendPersonId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetType(v int64) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.Type = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetSendTime(v util.LocalTime) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.SendTime = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetExchangeStatus(v int64) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.ExchangeStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetExchangeStatusDesc(v string) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.ExchangeStatusDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetTotalResourceCount(v int64) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.TotalResourceCount = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetReceiveResourceCount(v int64) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.ReceiveResourceCount = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetRejectResourceCount(v int64) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.RejectResourceCount = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetNote(v string) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.Note = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetWithdrawPerson(v string) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.WithdrawPerson = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetWithdrawPersonId(v string) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.WithdrawPersonId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetWithdrawTime(v util.LocalTime) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.WithdrawTime = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetCheckStatus(v int64) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.CheckStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetCheckStatusDesc(v string) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.CheckStatusDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxModel) SetRequestId(v int64) *AlibabaAlihealthSynergySyInboxDetailSyInboxModel {
	s.RequestId = &v
	return s
}
