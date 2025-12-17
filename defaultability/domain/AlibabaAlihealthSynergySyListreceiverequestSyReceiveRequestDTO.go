package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO struct {
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
	   明细     */
	DetailList *[]AlibabaAlihealthSynergySyListreceiverequestSyRequestDetailDTO `json:"detail_list,omitempty" `

	/*
	   索取ID（废弃）     */
	SyRequestId *int64 `json:"sy_request_id,omitempty" `

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
}

func (s *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO) SetSendRefEntId(v string) *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO {
	s.SendRefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO) SetSendEntName(v string) *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO {
	s.SendEntName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO) SetRequestNote(v string) *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO {
	s.RequestNote = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO) SetRequestInfo(v string) *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO {
	s.RequestInfo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO) SetType(v int64) *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO {
	s.Type = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO) SetDetailList(v []AlibabaAlihealthSynergySyListreceiverequestSyRequestDetailDTO) *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO {
	s.DetailList = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO) SetSyRequestId(v int64) *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO {
	s.SyRequestId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO) SetRequestPerson(v string) *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO {
	s.RequestPerson = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO) SetRequestStatus(v string) *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO {
	s.RequestStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO) SetRequestStatusDesc(v string) *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO {
	s.RequestStatusDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO) SetGmtCreate(v util.LocalTime) *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO {
	s.GmtCreate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO) SetId(v int64) *AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO {
	s.Id = &v
	return s
}
