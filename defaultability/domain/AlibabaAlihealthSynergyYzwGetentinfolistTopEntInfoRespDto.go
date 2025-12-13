package domain


type AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto struct {
    /*
        企业ID【ent_id】（单据上传时的收发货企业id就是填这个字段）     */
    EntId  *string `json:"ent_id,omitempty" `

    /*
        企业唯一标识【ref_ent_id】（单据上传时的货主企业ref_user_id就是填这个字段）     */
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
        企业所在区县代码     */
    AreaCode  *string `json:"area_code,omitempty" `

    /*
        是否入驻，1-入驻企业，0-非入驻     */
    SettleStatus  *int64 `json:"settle_status,omitempty" `

    /*
        企业注册详细地址     */
    RegRegionDetail  *string `json:"reg_region_detail,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto) SetEntId(v string) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto) SetAuditStatus(v int64) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto {
    s.AuditStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto) SetEntName(v string) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto) SetUniqueCode(v string) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto {
    s.UniqueCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto) SetLicTypeName(v string) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto {
    s.LicTypeName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto) SetLicTypeCode(v int64) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto {
    s.LicTypeCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto) SetProvName(v string) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto) SetProvCode(v string) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto {
    s.ProvCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto) SetCityName(v string) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto) SetCityCode(v string) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto {
    s.CityCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto) SetAreaName(v string) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto) SetAreaCode(v string) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto {
    s.AreaCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto) SetSettleStatus(v int64) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto {
    s.SettleStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto) SetRegRegionDetail(v string) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoRespDto {
    s.RegRegionDetail = &v
    return s
}
