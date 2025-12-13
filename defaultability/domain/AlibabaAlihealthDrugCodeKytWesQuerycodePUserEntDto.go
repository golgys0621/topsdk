package domain


type AlibabaAlihealthDrugCodeKytWesQuerycodePUserEntDto struct {
    /*
        生产厂的refEntId     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        生产厂企业名称     */
    EntName  *string `json:"ent_name,omitempty" `

}

func (s *AlibabaAlihealthDrugCodeKytWesQuerycodePUserEntDto) SetRefEntId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodePUserEntDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodePUserEntDto) SetEntName(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodePUserEntDto {
    s.EntName = &v
    return s
}
