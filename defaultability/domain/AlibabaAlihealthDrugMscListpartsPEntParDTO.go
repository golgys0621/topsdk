package domain


type AlibabaAlihealthDrugMscListpartsPEntParDTO struct {
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
    EntExtend  *AlibabaAlihealthDrugMscListpartsEntExtend `json:"ent_extend,omitempty" `

}

func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetPartnerId(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.PartnerId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetPartnerName(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.PartnerName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetEntId(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetRefEntId(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetEntProvCode(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.EntProvCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetProvName(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetAreaName(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetCityName(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetIsNetwork(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.IsNetwork = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetPartnerCapitalName(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.PartnerCapitalName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetPartnerType(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.PartnerType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetPartnerEntId(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.PartnerEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetLastModDate(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.LastModDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetCrtDate(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.CrtDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetCrtIcName(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.CrtIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetStatus(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetModIcName(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.ModIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetPartnerLevel(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.PartnerLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetModIcCode(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.ModIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetPEntParId(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.PEntParId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetCrtIcCode(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.CrtIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetParRefEntId(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.ParRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetAuditFlag(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.AuditFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetPartnerTypeDesc(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.PartnerTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetEntExtend(v AlibabaAlihealthDrugMscListpartsEntExtend) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.EntExtend = &v
    return s
}
