package domain


type AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto struct {
    /*
        企业id     */
    EntId  *string `json:"ent_id,omitempty" `

    /*
        企业唯一标识     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        是否入网     */
    IsNetwork  *string `json:"is_network,omitempty" `

    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" `

    /*
        所属管理机构     */
    DirectManage  *string `json:"direct_manage,omitempty" `

    /*
        是否法人     */
    LegalOrgFlag  *string `json:"legal_org_flag,omitempty" `

    /*
        机构代码     */
    OrgCode  *string `json:"org_code,omitempty" `

    /*
        注册地编码     */
    RegRegionCode  *string `json:"reg_region_code,omitempty" `

    /*
        注册地明细     */
    RegRegionDetail  *string `json:"reg_region_detail,omitempty" `

    /*
        所在地编码     */
    DictRegionCode  *string `json:"dict_region_code,omitempty" `

    /*
        所在地明细     */
    DictRegionDetail  *string `json:"dict_region_detail,omitempty" `

    /*
        状态1.使用中0.已废除     */
    Status  *string `json:"status,omitempty" `

    /*
        企业类型[1,生产企业 2批发企业 3医疗机构 4药店 5物流]     */
    UserRoleTypeStr  *string `json:"user_role_type_str,omitempty" `

    /*
        企业类型编码     */
    UserRoleType  *string `json:"user_role_type,omitempty" `

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

func (s *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto) SetEntId(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto) SetIsNetwork(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto {
    s.IsNetwork = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto) SetEntName(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto) SetDirectManage(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto {
    s.DirectManage = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto) SetLegalOrgFlag(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto {
    s.LegalOrgFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto) SetOrgCode(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto {
    s.OrgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto) SetRegRegionCode(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto {
    s.RegRegionCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto) SetRegRegionDetail(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto {
    s.RegRegionDetail = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto) SetDictRegionCode(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto {
    s.DictRegionCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto) SetDictRegionDetail(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto {
    s.DictRegionDetail = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto) SetStatus(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto) SetUserRoleTypeStr(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto {
    s.UserRoleTypeStr = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto) SetUserRoleType(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto {
    s.UserRoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto) SetEntOrgType(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto {
    s.EntOrgType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto) SetProvName(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto) SetAreaName(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto) SetCityName(v string) *AlibabaAlihealthDrugKytWesGetbyrefentidPUserEntInfoDto {
    s.CityName = &v
    return s
}
