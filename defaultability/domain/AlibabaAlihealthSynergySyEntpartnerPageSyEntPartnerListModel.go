package domain


type AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel struct {
    /*
        合作企业refEntId     */
    PartnerRefEntId  *string `json:"partner_ref_ent_id,omitempty" `

    /*
        往来单位省份名称     */
    ProvName  *string `json:"prov_name,omitempty" `

    /*
        往来单位名称     */
    PartnerName  *string `json:"partner_name,omitempty" `

    /*
        maxWeight     */
    MaxWeight  *int64 `json:"max_weight,omitempty" `

    /*
        合作企业EntId     */
    PartnerEntId  *string `json:"partner_ent_id,omitempty" `

    /*
        0曾用信息1使用中     */
    UseStatus  *int64 `json:"use_status,omitempty" `

    /*
        tagList     */
    TagList  *AlibabaAlihealthSynergySyEntpartnerPageTagDefinitionDTO `json:"tag_list,omitempty" `

    /*
        往来单位详细地址     */
    RegRegionDetail  *string `json:"reg_region_detail,omitempty" `

    /*
        往来单位用户自定义编码     */
    CustomNo  *string `json:"custom_no,omitempty" `

    /*
        往来单位城市名称     */
    CityName  *string `json:"city_name,omitempty" `

    /*
        往来单位区域名称     */
    AreaName  *string `json:"area_name,omitempty" `

    /*
        exchangeStatus     */
    ExchangeStatus  *int64 `json:"exchange_status,omitempty" `

    /*
        是否有待更新，0没有1有     */
    UpdateStatus  *int64 `json:"update_status,omitempty" `

    /*
        归档状态0未归档1已归档     */
    ArchiveStatus  *int64 `json:"archive_status,omitempty" `

    /*
        往来单位大类型     */
    EntTypeDesc  *string `json:"ent_type_desc,omitempty" `

    /*
        过期状态0未过期1已过期     */
    ExpireStatus  *int64 `json:"expire_status,omitempty" `

    /*
        往来单位小类型     */
    EntOrgTypeDesc  *string `json:"ent_org_type_desc,omitempty" `

    /*
        id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        往来单位大类型code     */
    EntTypeCode  *string `json:"ent_type_code,omitempty" `

    /*
        tagIdList     */
    TagIdList  *[]int64 `json:"tag_id_list,omitempty" `

    /*
        往来单位小类型code     */
    EntOrgTypeCode  *string `json:"ent_org_type_code,omitempty" `

}

func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetPartnerRefEntId(v string) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.PartnerRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetProvName(v string) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetPartnerName(v string) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.PartnerName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetMaxWeight(v int64) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.MaxWeight = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetPartnerEntId(v string) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.PartnerEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetUseStatus(v int64) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.UseStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetTagList(v AlibabaAlihealthSynergySyEntpartnerPageTagDefinitionDTO) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.TagList = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetRegRegionDetail(v string) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.RegRegionDetail = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetCustomNo(v string) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.CustomNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetCityName(v string) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetAreaName(v string) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetExchangeStatus(v int64) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.ExchangeStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetUpdateStatus(v int64) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.UpdateStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetArchiveStatus(v int64) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.ArchiveStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetEntTypeDesc(v string) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.EntTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetExpireStatus(v int64) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.ExpireStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetEntOrgTypeDesc(v string) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.EntOrgTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetId(v int64) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.Id = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetEntTypeCode(v string) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.EntTypeCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetTagIdList(v []int64) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.TagIdList = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) SetEntOrgTypeCode(v string) *AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel {
    s.EntOrgTypeCode = &v
    return s
}
