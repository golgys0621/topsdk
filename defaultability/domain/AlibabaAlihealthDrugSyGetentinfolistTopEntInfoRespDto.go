package domain


type AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto struct {
    /*
        企业ID【ent_id】     */
    EntId  *string `json:"ent_id,omitempty" `

    /*
        企业唯一标识【ref_ent_id】     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        1-审核通过，0-审核中，2-审核不通过     */
    AuditStatus  *int64 `json:"audit_status,omitempty" `

    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" `

    /*
        唯一代码     */
    UniqueCode  *string `json:"unique_code,omitempty" `

    /*
        唯一代码来源的资质名称（非精准）     */
    LicTypeName  *string `json:"lic_type_name,omitempty" `

    /*
        唯一代码来源的资质代码（非精准）     */
    LicTypeCode  *int64 `json:"lic_type_code,omitempty" `

    /*
        企业所在省份名称     */
    ProvName  *string `json:"prov_name,omitempty" `

    /*
        企业所在省份代码     */
    ProvCode  *string `json:"prov_code,omitempty" `

    /*
        企业所在城市名称     */
    CityName  *string `json:"city_name,omitempty" `

    /*
        企业所在城市代码     */
    CityCode  *string `json:"city_code,omitempty" `

    /*
        企业所在区县名称     */
    AreaName  *string `json:"area_name,omitempty" `

    /*
        企业所在区县代码       */
    AreaCode  *string `json:"area_code,omitempty" `

    /*
        是否入驻，1-入驻企业，0-非入驻     */
    SettleStatus  *int64 `json:"settle_status,omitempty" `

    /*
        企业注册详细地址     */
    RegRegionDetail  *string `json:"reg_region_detail,omitempty" `

}

func (s *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto) SetEntId(v string) *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto) SetRefEntId(v string) *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto) SetAuditStatus(v int64) *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto {
    s.AuditStatus = &v
    return s
}
func (s *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto) SetEntName(v string) *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto) SetUniqueCode(v string) *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto {
    s.UniqueCode = &v
    return s
}
func (s *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto) SetLicTypeName(v string) *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto {
    s.LicTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto) SetLicTypeCode(v int64) *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto {
    s.LicTypeCode = &v
    return s
}
func (s *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto) SetProvName(v string) *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto) SetProvCode(v string) *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto {
    s.ProvCode = &v
    return s
}
func (s *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto) SetCityName(v string) *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto) SetCityCode(v string) *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto {
    s.CityCode = &v
    return s
}
func (s *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto) SetAreaName(v string) *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto) SetAreaCode(v string) *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto {
    s.AreaCode = &v
    return s
}
func (s *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto) SetSettleStatus(v int64) *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto {
    s.SettleStatus = &v
    return s
}
func (s *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto) SetRegRegionDetail(v string) *AlibabaAlihealthDrugSyGetentinfolistTopEntInfoRespDto {
    s.RegRegionDetail = &v
    return s
}
