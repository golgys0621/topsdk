package domain


type AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto struct {
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

func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto) SetEntId(v string) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto) SetRefEntId(v string) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto) SetAuditStatus(v int64) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto {
    s.AuditStatus = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto) SetEntName(v string) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto) SetUniqueCode(v string) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto {
    s.UniqueCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto) SetLicTypeName(v string) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto {
    s.LicTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto) SetLicTypeCode(v int64) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto {
    s.LicTypeCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto) SetProvName(v string) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto) SetProvCode(v string) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto {
    s.ProvCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto) SetCityName(v string) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto) SetCityCode(v string) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto {
    s.CityCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto) SetAreaName(v string) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto) SetAreaCode(v string) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto {
    s.AreaCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto) SetSettleStatus(v int64) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto {
    s.SettleStatus = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto) SetRegRegionDetail(v string) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoRespDto {
    s.RegRegionDetail = &v
    return s
}
