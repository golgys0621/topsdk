package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel struct {
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

func (s *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel) SetId(v int64) *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel {
	s.Id = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel) SetGmtCreate(v util.LocalTime) *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel {
	s.GmtCreate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel) SetGmtModified(v util.LocalTime) *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel {
	s.GmtModified = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel) SetRefEntId(v string) *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel) SetResourceName(v string) *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel {
	s.ResourceName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel) SetResourceTypeCode(v string) *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel {
	s.ResourceTypeCode = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel) SetResourceTypeDesc(v string) *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel {
	s.ResourceTypeDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel) SetExpireDate(v util.LocalTime) *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel {
	s.ExpireDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel) SetNote(v string) *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel {
	s.Note = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel) SetPdfOssKeyUrl(v string) *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel {
	s.PdfOssKeyUrl = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel) SetUseStatus(v int64) *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel {
	s.UseStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel) SetArchiveStatus(v int64) *AlibabaAlihealthSynergySyProductResourceListSyProductResourceModel {
	s.ArchiveStatus = &v
	return s
}
