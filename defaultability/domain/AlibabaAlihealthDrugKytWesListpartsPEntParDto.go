package domain


type AlibabaAlihealthDrugKytWesListpartsPEntParDto struct {
    /*
        往来单位ID：企业自定义编号     */
    PartnerId  *string `json:"partner_id,omitempty" `

    /*
        往来单位名称     */
    PartnerName  *string `json:"partner_name,omitempty" `

    /*
        企业id：废弃字段     */
    EntId  *string `json:"ent_id,omitempty" `

    /*
        调用企业唯一标识     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        往来单位企业所在省编码     */
    EntProvCode  *string `json:"ent_prov_code,omitempty" `

    /*
        往来单位所在省     */
    ProvName  *string `json:"prov_name,omitempty" `

    /*
        往来单位所在市     */
    AreaName  *string `json:"area_name,omitempty" `

    /*
        往来单位所在县     */
    CityName  *string `json:"city_name,omitempty" `

    /*
        是不是入网企业：1-是；0-不是     */
    IsNetwork  *string `json:"is_network,omitempty" `

    /*
        往来单位拼音缩写     */
    PartnerCapitalName  *string `json:"partner_capital_name,omitempty" `

    /*
        往来单位类型     */
    PartnerType  *string `json:"partner_type,omitempty" `

    /*
        往来单位企业entId     */
    PartnerEntId  *string `json:"partner_ent_id,omitempty" `

    /*
        往来单位最近修改日期     */
    LastModDate  *string `json:"last_mod_date,omitempty" `

    /*
        添加到本企业往来单位列表日期     */
    CrtDate  *string `json:"crt_date,omitempty" `

    /*
        创建IC名称：废弃字段     */
    CrtIcName  *string `json:"crt_ic_name,omitempty" `

    /*
        状态     */
    Status  *string `json:"status,omitempty" `

    /*
        修改IC名称：废弃字段     */
    ModIcName  *string `json:"mod_ic_name,omitempty" `

    /*
        级别：废弃字段     */
    PartnerLevel  *string `json:"partner_level,omitempty" `

    /*
        修改IC码：废弃字段     */
    ModIcCode  *string `json:"mod_ic_code,omitempty" `

    /*
        记录ID     */
    PEntParId  *string `json:"p_ent_par_id,omitempty" `

    /*
        创建IC码：废弃字段     */
    CrtIcCode  *string `json:"crt_ic_code,omitempty" `

    /*
        往来单位企业refEntId     */
    ParRefEntId  *string `json:"par_ref_ent_id,omitempty" `

    /*
        往来单位审核状态：0-审核中；1-审核通过；2-审核不通过     */
    AuditFlag  *string `json:"audit_flag,omitempty" `

    /*
        往来单位企业类型描述     */
    PartnerTypeDesc  *string `json:"partner_type_desc,omitempty" `

    /*
        拓展属性     */
    EntExtend  *AlibabaAlihealthDrugKytWesListpartsEntExtend `json:"ent_extend,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetPartnerId(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.PartnerId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetPartnerName(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.PartnerName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetEntId(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetEntProvCode(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.EntProvCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetProvName(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetAreaName(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetCityName(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetIsNetwork(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.IsNetwork = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetPartnerCapitalName(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.PartnerCapitalName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetPartnerType(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.PartnerType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetPartnerEntId(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.PartnerEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetLastModDate(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.LastModDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetCrtDate(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.CrtDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetCrtIcName(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.CrtIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetStatus(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetModIcName(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.ModIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetPartnerLevel(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.PartnerLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetModIcCode(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.ModIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetPEntParId(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.PEntParId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetCrtIcCode(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.CrtIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetParRefEntId(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.ParRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetAuditFlag(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.AuditFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetPartnerTypeDesc(v string) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.PartnerTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPEntParDto) SetEntExtend(v AlibabaAlihealthDrugKytWesListpartsEntExtend) *AlibabaAlihealthDrugKytWesListpartsPEntParDto {
    s.EntExtend = &v
    return s
}
