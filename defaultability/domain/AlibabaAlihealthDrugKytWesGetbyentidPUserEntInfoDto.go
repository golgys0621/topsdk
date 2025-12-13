package domain


type AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto struct {
    /*
        所在地编码     */
    DictRegionCode  *string `json:"dict_region_code,omitempty" `

    /*
        所在地明细     */
    RegRegionDetail  *string `json:"reg_region_detail,omitempty" `

    /*
        注册地编码     */
    RegRegionCode  *string `json:"reg_region_code,omitempty" `

    /*
        注册地明细     */
    OrgCode  *string `json:"org_code,omitempty" `

    /*
        是否法人     */
    LegalOrgFlag  *string `json:"legal_org_flag,omitempty" `

    /*
        所属管理机构     */
    DirectManage  *string `json:"direct_manage,omitempty" `

    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" `

    /*
        是否入网     */
    IsNetwork  *string `json:"is_network,omitempty" `

    /*
        企业唯一标识     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        企业id     */
    EntId  *string `json:"ent_id,omitempty" `

    /*
        所在地明细     */
    DictRegionDetail  *string `json:"dict_region_detail,omitempty" `

    /*
        状态1.使用中0.已废除     */
    Status  *string `json:"status,omitempty" `

    /*
        拼音缩写     */
    EntCapitalName  *string `json:"ent_capital_name,omitempty" `

    /*
        企业类型     */
    UserRoleType  *string `json:"user_role_type,omitempty" `

    /*
        企业类型编码     */
    UserRoleTypeStr  *string `json:"user_role_type_str,omitempty" `

    /*
        企业机构详细类别     */
    EntOrgType  *string `json:"ent_org_type,omitempty" `

    /*
        省     */
    ProvName  *string `json:"prov_name,omitempty" `

    /*
        区县     */
    AreaName  *string `json:"area_name,omitempty" `

    /*
        市     */
    CityName  *string `json:"city_name,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) SetDictRegionCode(v string) *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto {
    s.DictRegionCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) SetRegRegionDetail(v string) *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto {
    s.RegRegionDetail = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) SetRegRegionCode(v string) *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto {
    s.RegRegionCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) SetOrgCode(v string) *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto {
    s.OrgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) SetLegalOrgFlag(v string) *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto {
    s.LegalOrgFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) SetDirectManage(v string) *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto {
    s.DirectManage = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) SetEntName(v string) *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) SetIsNetwork(v string) *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto {
    s.IsNetwork = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) SetEntId(v string) *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) SetDictRegionDetail(v string) *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto {
    s.DictRegionDetail = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) SetStatus(v string) *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) SetEntCapitalName(v string) *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto {
    s.EntCapitalName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) SetUserRoleType(v string) *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto {
    s.UserRoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) SetUserRoleTypeStr(v string) *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto {
    s.UserRoleTypeStr = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) SetEntOrgType(v string) *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto {
    s.EntOrgType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) SetProvName(v string) *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) SetAreaName(v string) *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto) SetCityName(v string) *AlibabaAlihealthDrugKytWesGetbyentidPUserEntInfoDto {
    s.CityName = &v
    return s
}
