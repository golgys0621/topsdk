package domain


type AlibabaAlihealthDrugKytWesGetentinfoPUserEntInfoDto struct {
    /*
        企业唯一标识【ref_ent_id】（单据上传时的货主企业ref_user_id就是填这个字段）     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        企业ID【ent_id】（单据上传时的收发货企业id就是填这个字段）     */
    EntId  *string `json:"ent_id,omitempty" `

    /*
        入网标识【0非入网;1开头代表入网企业;200代表入驻码上放心的企业】     */
    NetworkType  *string `json:"network_type,omitempty" `

    /*
        1-审核通过，0-审核中，2-审核不通过     */
    AuditStatus  *string `json:"audit_status,omitempty" `

    /*
        拓展属性     */
    EntExtend  *AlibabaAlihealthDrugKytWesGetentinfoEntExtend `json:"ent_extend,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesGetentinfoPUserEntInfoDto) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesGetentinfoPUserEntInfoDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetentinfoPUserEntInfoDto) SetEntId(v string) *AlibabaAlihealthDrugKytWesGetentinfoPUserEntInfoDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetentinfoPUserEntInfoDto) SetNetworkType(v string) *AlibabaAlihealthDrugKytWesGetentinfoPUserEntInfoDto {
    s.NetworkType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetentinfoPUserEntInfoDto) SetAuditStatus(v string) *AlibabaAlihealthDrugKytWesGetentinfoPUserEntInfoDto {
    s.AuditStatus = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesGetentinfoPUserEntInfoDto) SetEntExtend(v AlibabaAlihealthDrugKytWesGetentinfoEntExtend) *AlibabaAlihealthDrugKytWesGetentinfoPUserEntInfoDto {
    s.EntExtend = &v
    return s
}
