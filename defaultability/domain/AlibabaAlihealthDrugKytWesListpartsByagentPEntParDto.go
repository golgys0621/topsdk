package domain


type AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto struct {
    /*
        往来单位ID：企业自定义编号     */
    PartnerId  *string `json:"partner_id,omitempty" `

    /*
        往来单位名称     */
    PartnerName  *string `json:"partner_name,omitempty" `

    /*
        企业id: 废弃字段     */
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
        创建IC名称: 废弃字段     */
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
    EntExtend  *AlibabaAlihealthDrugKytWesListpartsByagentEntExtend `json:"ent_extend,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetPartnerId(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.PartnerId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetPartnerName(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.PartnerName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetEntId(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetEntProvCode(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.EntProvCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetProvName(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetAreaName(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetCityName(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetIsNetwork(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.IsNetwork = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetPartnerCapitalName(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.PartnerCapitalName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetPartnerType(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.PartnerType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetPartnerEntId(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.PartnerEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetLastModDate(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.LastModDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetCrtDate(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.CrtDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetCrtIcName(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.CrtIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetStatus(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetModIcName(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.ModIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetPartnerLevel(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.PartnerLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetModIcCode(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.ModIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetPEntParId(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.PEntParId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetCrtIcCode(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.CrtIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetParRefEntId(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.ParRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetAuditFlag(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.AuditFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetPartnerTypeDesc(v string) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.PartnerTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) SetEntExtend(v AlibabaAlihealthDrugKytWesListpartsByagentEntExtend) *AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto {
    s.EntExtend = &v
    return s
}
