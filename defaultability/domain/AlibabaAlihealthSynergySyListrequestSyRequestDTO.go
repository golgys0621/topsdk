package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyListrequestSyRequestDTO struct {
	/*
	   主键     */
	Id *int64 `json:"id,omitempty" `

	/*
	   创建时间     */
	GmtCreate *util.LocalTime `json:"gmt_create,omitempty" `

	/*
	   更新时间     */
	GmtModified *util.LocalTime `json:"gmt_modified,omitempty" `

	/*
	   企业ref_ent_id     */
	RefEntId *string `json:"ref_ent_id,omitempty" `

	/*
	   企业ent_id     */
	EntId *string `json:"ent_id,omitempty" `

	/*
	   被索取的企业ref_ent_id     */
	ReceiveRefEntId *string `json:"receive_ref_ent_id,omitempty" `

	/*
	   被索取的企业ent_id     */
	ReceiveEntId *string `json:"receive_ent_id,omitempty" `

	/*
	   被索取的企业名称     */
	ReceiveEntName *string `json:"receive_ent_name,omitempty" `

	/*
	   索取内容     */
	RequestInfo *string `json:"request_info,omitempty" `

	/*
	   索取数量     */
	RequestCount *int64 `json:"request_count,omitempty" `

	/*
	   0企业资料1产品资料     */
	Type *int64 `json:"type,omitempty" `

	/*
	   索取人     */
	RequestPerson *string `json:"request_person,omitempty" `

	/*
	   索取人id     */
	RequestPersonId *string `json:"request_person_id,omitempty" `

	/*
	   索取状态     */
	RequestStatus *string `json:"request_status,omitempty" `

	/*
	   索取状态描述     */
	RequestStatusDesc *string `json:"request_status_desc,omitempty" `

	/*
	   索取留言     */
	RequestNote *string `json:"request_note,omitempty" `
}

func (s *AlibabaAlihealthSynergySyListrequestSyRequestDTO) SetId(v int64) *AlibabaAlihealthSynergySyListrequestSyRequestDTO {
	s.Id = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestSyRequestDTO) SetGmtCreate(v util.LocalTime) *AlibabaAlihealthSynergySyListrequestSyRequestDTO {
	s.GmtCreate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestSyRequestDTO) SetGmtModified(v util.LocalTime) *AlibabaAlihealthSynergySyListrequestSyRequestDTO {
	s.GmtModified = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestSyRequestDTO) SetRefEntId(v string) *AlibabaAlihealthSynergySyListrequestSyRequestDTO {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestSyRequestDTO) SetEntId(v string) *AlibabaAlihealthSynergySyListrequestSyRequestDTO {
	s.EntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestSyRequestDTO) SetReceiveRefEntId(v string) *AlibabaAlihealthSynergySyListrequestSyRequestDTO {
	s.ReceiveRefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestSyRequestDTO) SetReceiveEntId(v string) *AlibabaAlihealthSynergySyListrequestSyRequestDTO {
	s.ReceiveEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestSyRequestDTO) SetReceiveEntName(v string) *AlibabaAlihealthSynergySyListrequestSyRequestDTO {
	s.ReceiveEntName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestSyRequestDTO) SetRequestInfo(v string) *AlibabaAlihealthSynergySyListrequestSyRequestDTO {
	s.RequestInfo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestSyRequestDTO) SetRequestCount(v int64) *AlibabaAlihealthSynergySyListrequestSyRequestDTO {
	s.RequestCount = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestSyRequestDTO) SetType(v int64) *AlibabaAlihealthSynergySyListrequestSyRequestDTO {
	s.Type = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestSyRequestDTO) SetRequestPerson(v string) *AlibabaAlihealthSynergySyListrequestSyRequestDTO {
	s.RequestPerson = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestSyRequestDTO) SetRequestPersonId(v string) *AlibabaAlihealthSynergySyListrequestSyRequestDTO {
	s.RequestPersonId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestSyRequestDTO) SetRequestStatus(v string) *AlibabaAlihealthSynergySyListrequestSyRequestDTO {
	s.RequestStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestSyRequestDTO) SetRequestStatusDesc(v string) *AlibabaAlihealthSynergySyListrequestSyRequestDTO {
	s.RequestStatusDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestSyRequestDTO) SetRequestNote(v string) *AlibabaAlihealthSynergySyListrequestSyRequestDTO {
	s.RequestNote = &v
	return s
}
