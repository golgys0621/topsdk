package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel struct {
	/*
	   资料id     */
	Id *int64 `json:"id,omitempty" `

	/*
	   创建时间     */
	GmtCreate *util.LocalTime `json:"gmt_create,omitempty" `

	/*
	   修改时间     */
	GmtModified *util.LocalTime `json:"gmt_modified,omitempty" `

	/*
	   企业id     */
	RefEntId *string `json:"ref_ent_id,omitempty" `

	/*
	   资料名称     */
	ResourceName *string `json:"resource_name,omitempty" `

	/*
	   资料分类     */
	ResourceTypeCode *string `json:"resource_type_code,omitempty" `

	/*
	   资料分类描述     */
	ResourceTypeDesc *string `json:"resource_type_desc,omitempty" `

	/*
	   有效期，长期返回2999-01-01 00:00:00     */
	ExpireDate *util.LocalTime `json:"expire_date,omitempty" `

	/*
	   备注     */
	Note *string `json:"note,omitempty" `

	/*
	   文件地址，20分钟有效期     */
	PdfOssKeyUrl *string `json:"pdf_oss_key_url,omitempty" `

	/*
	   0已过期 1正常     */
	UseStatus *int64 `json:"use_status,omitempty" `

	/*
	   归档状态：0未归档1已归档     */
	ArchiveStatus *int64 `json:"archive_status,omitempty" `
}

func (s *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel) SetId(v int64) *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel {
	s.Id = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel) SetGmtCreate(v util.LocalTime) *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel {
	s.GmtCreate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel) SetGmtModified(v util.LocalTime) *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel {
	s.GmtModified = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel) SetRefEntId(v string) *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel) SetResourceName(v string) *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel {
	s.ResourceName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel) SetResourceTypeCode(v string) *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel {
	s.ResourceTypeCode = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel) SetResourceTypeDesc(v string) *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel {
	s.ResourceTypeDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel) SetExpireDate(v util.LocalTime) *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel {
	s.ExpireDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel) SetNote(v string) *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel {
	s.Note = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel) SetPdfOssKeyUrl(v string) *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel {
	s.PdfOssKeyUrl = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel) SetUseStatus(v int64) *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel {
	s.UseStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel) SetArchiveStatus(v int64) *AlibabaAlihealthSynergySyAgentpersonResourceListSyAgentPersonResourceModel {
	s.ArchiveStatus = &v
	return s
}
