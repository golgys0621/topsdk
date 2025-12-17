package domain


type AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel struct {
    /*
        资料id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        创建时间     */
    GmtCreate  *string `json:"gmt_create,omitempty" `

    /*
        修改时间     */
    GmtModified  *string `json:"gmt_modified,omitempty" `

    /*
        资料名称     */
    ResourceName  *string `json:"resource_name,omitempty" `

    /*
        资料分类编码     */
    ResourceTypeCode  *string `json:"resource_type_code,omitempty" `

    /*
        资料分类描述     */
    ResourceTypeDesc  *string `json:"resource_type_desc,omitempty" `

    /*
        有效期，长期返回2999-01-01 00:00:00     */
    ExpireDate  *string `json:"expire_date,omitempty" `

    /*
        文件地址，20分钟有效期     */
    PdfOssKeyUrl  *string `json:"pdf_oss_key_url,omitempty" `

    /*
        0已过期 1正常     */
    UseStatus  *int64 `json:"use_status,omitempty" `

    /*
        归档状态：0未归档1已归档     */
    ArchiveStatus  *int64 `json:"archive_status,omitempty" `

    /*
        备注     */
    Note  *string `json:"note,omitempty" `

    /*
        扩展属性，json对象数组，code：资质支持的属性code，value:值     */
    ExtendInfo  *string `json:"extend_info,omitempty" `

}

func (s *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel) SetId(v int64) *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel {
    s.Id = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel) SetGmtCreate(v string) *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel {
    s.GmtCreate = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel) SetGmtModified(v string) *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel {
    s.GmtModified = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel) SetResourceName(v string) *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel {
    s.ResourceName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel) SetResourceTypeCode(v string) *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel {
    s.ResourceTypeCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel) SetResourceTypeDesc(v string) *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel {
    s.ResourceTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel) SetExpireDate(v string) *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel {
    s.ExpireDate = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel) SetPdfOssKeyUrl(v string) *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel {
    s.PdfOssKeyUrl = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel) SetUseStatus(v int64) *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel {
    s.UseStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel) SetArchiveStatus(v int64) *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel {
    s.ArchiveStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel) SetNote(v string) *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel {
    s.Note = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel) SetExtendInfo(v string) *AlibabaAlihealthSynergySyEntResourceListSyEntPartnerResourceDetailModel {
    s.ExtendInfo = &v
    return s
}
