package domain


type AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel struct {
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

    /*
        合作企业refEntId     */
    PartnerRefEntId  *string `json:"partner_ref_ent_id,omitempty" `

    /*
        合作企业entId     */
    PartnerEntId  *string `json:"partner_ent_id,omitempty" `

    /*
        合作企业名称     */
    PartnerEntName  *string `json:"partner_ent_name,omitempty" `

}

func (s *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel) SetId(v int64) *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel {
    s.Id = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel) SetGmtCreate(v string) *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel {
    s.GmtCreate = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel) SetGmtModified(v string) *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel {
    s.GmtModified = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel) SetResourceName(v string) *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel {
    s.ResourceName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel) SetResourceTypeCode(v string) *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel {
    s.ResourceTypeCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel) SetResourceTypeDesc(v string) *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel {
    s.ResourceTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel) SetExpireDate(v string) *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel {
    s.ExpireDate = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel) SetPdfOssKeyUrl(v string) *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel {
    s.PdfOssKeyUrl = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel) SetUseStatus(v int64) *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel {
    s.UseStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel) SetArchiveStatus(v int64) *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel {
    s.ArchiveStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel) SetNote(v string) *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel {
    s.Note = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel) SetExtendInfo(v string) *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel {
    s.ExtendInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel) SetPartnerRefEntId(v string) *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel {
    s.PartnerRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel) SetPartnerEntId(v string) *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel {
    s.PartnerEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel) SetPartnerEntName(v string) *AlibabaAlihealthSynergySyEntpartnerResourceListSyEntPartnerResourceDetailModel {
    s.PartnerEntName = &v
    return s
}
