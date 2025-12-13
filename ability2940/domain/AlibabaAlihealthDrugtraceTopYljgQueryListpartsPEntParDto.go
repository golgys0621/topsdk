package domain


type AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto struct {
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
    EntExtend  *AlibabaAlihealthDrugtraceTopYljgQueryListpartsEntExtend `json:"ent_extend,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetPartnerId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.PartnerId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetPartnerName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.PartnerName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetEntProvCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.EntProvCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetProvName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetAreaName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetCityName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetIsNetwork(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.IsNetwork = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetPartnerCapitalName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.PartnerCapitalName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetPartnerType(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.PartnerType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetPartnerEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.PartnerEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetLastModDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.LastModDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetCrtDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.CrtDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetCrtIcName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.CrtIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetStatus(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetModIcName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.ModIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetPartnerLevel(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.PartnerLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetModIcCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.ModIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetPEntParId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.PEntParId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetCrtIcCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.CrtIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetParRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.ParRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetAuditFlag(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.AuditFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetPartnerTypeDesc(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.PartnerTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetEntExtend(v AlibabaAlihealthDrugtraceTopYljgQueryListpartsEntExtend) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.EntExtend = &v
    return s
}
