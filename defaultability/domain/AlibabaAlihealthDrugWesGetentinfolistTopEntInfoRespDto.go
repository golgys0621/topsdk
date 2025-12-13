package domain


type AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto struct {
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

func (s *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto) SetEntId(v string) *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto) SetRefEntId(v string) *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto) SetAuditStatus(v int64) *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto {
    s.AuditStatus = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto) SetEntName(v string) *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto) SetUniqueCode(v string) *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto {
    s.UniqueCode = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto) SetLicTypeName(v string) *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto {
    s.LicTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto) SetLicTypeCode(v int64) *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto {
    s.LicTypeCode = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto) SetProvName(v string) *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto) SetProvCode(v string) *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto {
    s.ProvCode = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto) SetCityName(v string) *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto) SetCityCode(v string) *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto {
    s.CityCode = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto) SetAreaName(v string) *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto) SetAreaCode(v string) *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto {
    s.AreaCode = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto) SetSettleStatus(v int64) *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto {
    s.SettleStatus = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto) SetRegRegionDetail(v string) *AlibabaAlihealthDrugWesGetentinfolistTopEntInfoRespDto {
    s.RegRegionDetail = &v
    return s
}
