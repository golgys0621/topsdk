package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO struct {
	/*
	   requestCount     */
	RequestCount *int64 `json:"request_count,omitempty" `

	/*
	   gmtModified     */
	GmtModified *util.LocalTime `json:"gmt_modified,omitempty" `

	/*
	   requestPersonId     */
	RequestPersonId *string `json:"request_person_id,omitempty" `

	/*
	   entId     */
	EntId *string `json:"ent_id,omitempty" `

	/*
	   索取状态描述     */
	RequestStatusDesc *string `json:"request_status_desc,omitempty" `

	/*
	   索取留言     */
	RequestNote *string `json:"request_note,omitempty" `

	/*
	   gmtCreate     */
	GmtCreate *util.LocalTime `json:"gmt_create,omitempty" `

	/*
	   0企业资料1产品资料     */
	Type *int64 `json:"type,omitempty" `

	/*
	   索取人     */
	RequestPerson *string `json:"request_person,omitempty" `

	/*
	   被索取的企业ent_id     */
	ReceiveEntId *string `json:"receive_ent_id,omitempty" `

	/*
	   被索取的企业ref_ent_id     */
	ReceiveRefEntId *string `json:"receive_ref_ent_id,omitempty" `

	/*
	   0为删除1已删除     */
	DeleteStatus *int64 `json:"delete_status,omitempty" `

	/*
	   detailList     */
	DetailList *[]AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO `json:"detail_list,omitempty" `

	/*
	   refEntId     */
	RefEntId *string `json:"ref_ent_id,omitempty" `

	/*
	   索取内容     */
	RequestInfo *string `json:"request_info,omitempty" `

	/*
	   id     */
	Id *int64 `json:"id,omitempty" `

	/*
	   receiveEntName     */
	ReceiveEntName *string `json:"receive_ent_name,omitempty" `

	/*
	   索取状态     */
	RequestStatus *string `json:"request_status,omitempty" `
}

func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO) SetRequestCount(v int64) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO {
	s.RequestCount = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO) SetGmtModified(v util.LocalTime) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO {
	s.GmtModified = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO) SetRequestPersonId(v string) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO {
	s.RequestPersonId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO) SetEntId(v string) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO {
	s.EntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO) SetRequestStatusDesc(v string) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO {
	s.RequestStatusDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO) SetRequestNote(v string) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO {
	s.RequestNote = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO) SetGmtCreate(v util.LocalTime) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO {
	s.GmtCreate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO) SetType(v int64) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO {
	s.Type = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO) SetRequestPerson(v string) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO {
	s.RequestPerson = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO) SetReceiveEntId(v string) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO {
	s.ReceiveEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO) SetReceiveRefEntId(v string) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO {
	s.ReceiveRefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO) SetDeleteStatus(v int64) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO {
	s.DeleteStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO) SetDetailList(v []AlibabaAlihealthSynergySyListrequestdetailSyRequestDetailDTO) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO {
	s.DetailList = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO) SetRefEntId(v string) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO) SetRequestInfo(v string) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO {
	s.RequestInfo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO) SetId(v int64) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO {
	s.Id = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO) SetReceiveEntName(v string) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO {
	s.ReceiveEntName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO) SetRequestStatus(v string) *AlibabaAlihealthSynergySyListrequestdetailSyRequestDTO {
	s.RequestStatus = &v
	return s
}
