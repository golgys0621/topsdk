package domain


type AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO struct {
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
        企业标识     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

}

func (s *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO) SetEntId(v string) *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO) SetEntCapitalName(v string) *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO {
    s.EntCapitalName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO) SetAreaName(v string) *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO) SetCityName(v string) *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO) SetDictRegionCode(v string) *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO {
    s.DictRegionCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO) SetEntName(v string) *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO) SetProvName(v string) *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO) SetRegRegionCode(v string) *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO {
    s.RegRegionCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO) SetUserRoleType(v string) *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO {
    s.UserRoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO) SetIsNetwork(v string) *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO {
    s.IsNetwork = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO) SetRefEntId(v string) *AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO {
    s.RefEntId = &v
    return s
}
