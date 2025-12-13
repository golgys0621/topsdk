package domain


type AlibabaAlihealthDrugKytWesGetbyorgcodePUserEntInfoDto struct {
    /*
        企业id     */
    EntId  *string `json:"ent_id,omitempty" `

    /*
        企业唯一标识     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        200 入驻     */
    NetworkType  *int64 `json:"network_type,omitempty" `

    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" `

    /*
        统一信用代码     */
    OrgCode  *string `json:"org_code,omitempty" `

    /*
        注册地址编码     */
    RegRegionCode  *string `json:"reg_region_code,omitempty" `

    /*
        企业类型编码     */
    UserRoleType  *int64 `json:"user_role_type,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesGetbyorgcodePUserEntInfoDto) SetEntId(v string) *AlibabaAlihealthDrugKytWesGetbyorgcodePUserEntInfoDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyorgcodePUserEntInfoDto) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesGetbyorgcodePUserEntInfoDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyorgcodePUserEntInfoDto) SetNetworkType(v int64) *AlibabaAlihealthDrugKytWesGetbyorgcodePUserEntInfoDto {
    s.NetworkType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyorgcodePUserEntInfoDto) SetEntName(v string) *AlibabaAlihealthDrugKytWesGetbyorgcodePUserEntInfoDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyorgcodePUserEntInfoDto) SetOrgCode(v string) *AlibabaAlihealthDrugKytWesGetbyorgcodePUserEntInfoDto {
    s.OrgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyorgcodePUserEntInfoDto) SetRegRegionCode(v string) *AlibabaAlihealthDrugKytWesGetbyorgcodePUserEntInfoDto {
    s.RegRegionCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetbyorgcodePUserEntInfoDto) SetUserRoleType(v int64) *AlibabaAlihealthDrugKytWesGetbyorgcodePUserEntInfoDto {
    s.UserRoleType = &v
    return s
}
