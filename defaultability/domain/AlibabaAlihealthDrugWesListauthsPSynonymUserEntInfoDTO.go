package domain


type AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO struct {
    /*
        企业ID     */
    EntId  *string `json:"ent_id,omitempty" `

    /*
        企业名称拼音     */
    EntCapitalName  *string `json:"ent_capital_name,omitempty" `

    /*
        区     */
    AreaName  *string `json:"area_name,omitempty" `

    /*
        市     */
    CityName  *string `json:"city_name,omitempty" `

    /*
        地市编码     */
    DictRegionCode  *string `json:"dict_region_code,omitempty" `

    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" `

    /*
        省     */
    ProvName  *string `json:"prov_name,omitempty" `

    /*
        注册编码     */
    RegRegionCode  *string `json:"reg_region_code,omitempty" `

    /*
        企业类型     */
    UserRoleType  *string `json:"user_role_type,omitempty" `

    /*
        是否入网     */
    IsNetwork  *string `json:"is_network,omitempty" `

    /*
        企业唯一标识     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

}

func (s *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO) SetEntId(v string) *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO) SetEntCapitalName(v string) *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO {
    s.EntCapitalName = &v
    return s
}
func (s *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO) SetAreaName(v string) *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO) SetCityName(v string) *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO) SetDictRegionCode(v string) *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO {
    s.DictRegionCode = &v
    return s
}
func (s *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO) SetEntName(v string) *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO) SetProvName(v string) *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO) SetRegRegionCode(v string) *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO {
    s.RegRegionCode = &v
    return s
}
func (s *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO) SetUserRoleType(v string) *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO {
    s.UserRoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO) SetIsNetwork(v string) *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO {
    s.IsNetwork = &v
    return s
}
func (s *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO) SetRefEntId(v string) *AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO {
    s.RefEntId = &v
    return s
}
