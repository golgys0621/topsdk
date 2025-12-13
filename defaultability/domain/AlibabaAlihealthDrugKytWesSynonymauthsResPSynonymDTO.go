package domain


type AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO struct {
    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" `

    /*
        市     */
    CityDesc  *string `json:"city_desc,omitempty" `

    /*
        省     */
    ProvDesc  *string `json:"prov_desc,omitempty" `

    /*
        区     */
    AreaDesc  *string `json:"area_desc,omitempty" `

    /*
        企业主键     */
    EntId  *string `json:"ent_id,omitempty" `

    /*
        区域编码     */
    DictRegionCode  *string `json:"dict_region_code,omitempty" `

    /*
        企业唯一标识     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        货主     */
    SynOwnEntId  *string `json:"syn_own_ent_id,omitempty" `

    /*
        货主标识     */
    UserEntId  *string `json:"user_ent_id,omitempty" `

    /*
        创建日期     */
    CrtDate  *string `json:"crt_date,omitempty" `

    /*
        角色     */
    UserRoleType  *string `json:"user_role_type,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO) SetEntName(v string) *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO) SetCityDesc(v string) *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO {
    s.CityDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO) SetProvDesc(v string) *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO {
    s.ProvDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO) SetAreaDesc(v string) *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO {
    s.AreaDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO) SetEntId(v string) *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO) SetDictRegionCode(v string) *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO {
    s.DictRegionCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO) SetSynOwnEntId(v string) *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO {
    s.SynOwnEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO) SetUserEntId(v string) *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO {
    s.UserEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO) SetCrtDate(v string) *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO {
    s.CrtDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO) SetUserRoleType(v string) *AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO {
    s.UserRoleType = &v
    return s
}
