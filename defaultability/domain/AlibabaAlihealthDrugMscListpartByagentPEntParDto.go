package domain


type AlibabaAlihealthDrugMscListpartByagentPEntParDto struct {
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
        往来单位所在县     */
    AreaName  *string `json:"area_name,omitempty" `

    /*
        往来单位所在市      */
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
    EntExtend  *AlibabaAlihealthDrugMscListpartByagentEntExtend `json:"ent_extend,omitempty" `

    /*
        1（待替换）：经平台数据治理，系统已匹配到正确企业。请自行确认并替换为正确的企业信息，并使用 replace_ref_ent_id 查询清洗后的企业信息。 2（待更新）：企业名称已发生变更。请使用 refentid 获取最新的企业信息。 其余返回值：已废弃，无需处理。     */
    Shared  *string `json:"shared,omitempty" `

    /*
        唯一认证代码     */
    OrgCode  *string `json:"org_code,omitempty" `

}

func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetPartnerId(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.PartnerId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetPartnerName(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.PartnerName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetEntId(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetRefEntId(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetEntProvCode(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.EntProvCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetProvName(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetAreaName(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetCityName(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetIsNetwork(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.IsNetwork = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetPartnerCapitalName(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.PartnerCapitalName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetPartnerType(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.PartnerType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetPartnerEntId(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.PartnerEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetLastModDate(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.LastModDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetCrtDate(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.CrtDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetCrtIcName(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.CrtIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetStatus(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetModIcName(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.ModIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetPartnerLevel(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.PartnerLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetModIcCode(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.ModIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetPEntParId(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.PEntParId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetCrtIcCode(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.CrtIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetParRefEntId(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.ParRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetAuditFlag(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.AuditFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetPartnerTypeDesc(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.PartnerTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetEntExtend(v AlibabaAlihealthDrugMscListpartByagentEntExtend) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.EntExtend = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetShared(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.Shared = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPEntParDto) SetOrgCode(v string) *AlibabaAlihealthDrugMscListpartByagentPEntParDto {
    s.OrgCode = &v
    return s
}
