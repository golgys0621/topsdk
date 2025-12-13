package domain


type AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto struct {
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
        扩展属性     */
    EntExtend  *AlibabaAlihealthDrugtraceTopLsydQueryListpartsEntExtend `json:"ent_extend,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetPartnerId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.PartnerId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetPartnerName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.PartnerName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetEntProvCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.EntProvCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetProvName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetAreaName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetCityName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetIsNetwork(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.IsNetwork = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetPartnerCapitalName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.PartnerCapitalName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetPartnerType(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.PartnerType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetPartnerEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.PartnerEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetLastModDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.LastModDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetCrtDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.CrtDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetCrtIcName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.CrtIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetStatus(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetModIcName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.ModIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetPartnerLevel(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.PartnerLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetModIcCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.ModIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetPEntParId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.PEntParId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetCrtIcCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.CrtIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetParRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.ParRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetAuditFlag(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.AuditFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetPartnerTypeDesc(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.PartnerTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetEntExtend(v AlibabaAlihealthDrugtraceTopLsydQueryListpartsEntExtend) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.EntExtend = &v
    return s
}
