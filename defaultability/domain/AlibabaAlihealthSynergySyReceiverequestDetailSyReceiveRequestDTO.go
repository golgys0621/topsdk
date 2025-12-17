package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO struct {
	/*
	   索取企业ID     */
	SendRefEntId *string `json:"send_ref_ent_id,omitempty" `

	/*
	   索取企业     */
	SendEntName *string `json:"send_ent_name,omitempty" `

	/*
	   留言     */
	RequestNote *string `json:"request_note,omitempty" `

	/*
	   索取内容     */
	RequestInfo *string `json:"request_info,omitempty" `

	/*
	   0:企业资料 1：产品资料     */
	Type *int64 `json:"type,omitempty" `

	/*
	   索取人     */
	RequestPerson *string `json:"request_person,omitempty" `

	/*
	   索取状态     */
	RequestStatus *string `json:"request_status,omitempty" `

	/*
	   索取状态     */
	RequestStatusDesc *string `json:"request_status_desc,omitempty" `

	/*
	   索取日期     */
	GmtCreate *util.LocalTime `json:"gmt_create,omitempty" `

	/*
	   被索取ID，用于拒绝索取，发送等     */
	Id *int64 `json:"id,omitempty" `

	/*
	   被索取的产品     */
	DetailList *[]AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDetailDTO `json:"detail_list,omitempty" `
}

func (s *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO) SetSendRefEntId(v string) *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO {
	s.SendRefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO) SetSendEntName(v string) *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO {
	s.SendEntName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO) SetRequestNote(v string) *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO {
	s.RequestNote = &v
	return s
}
func (s *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO) SetRequestInfo(v string) *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO {
	s.RequestInfo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO) SetType(v int64) *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO {
	s.Type = &v
	return s
}
func (s *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO) SetRequestPerson(v string) *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO {
	s.RequestPerson = &v
	return s
}
func (s *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO) SetRequestStatus(v string) *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO {
	s.RequestStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO) SetRequestStatusDesc(v string) *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO {
	s.RequestStatusDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO) SetGmtCreate(v util.LocalTime) *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO {
	s.GmtCreate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO) SetId(v int64) *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO {
	s.Id = &v
	return s
}
func (s *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO) SetDetailList(v []AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDetailDTO) *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDTO {
	s.DetailList = &v
	return s
}
