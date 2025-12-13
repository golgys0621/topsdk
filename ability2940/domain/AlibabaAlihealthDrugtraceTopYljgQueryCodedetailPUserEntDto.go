package domain


type AlibabaAlihealthDrugtraceTopYljgQueryCodedetailPUserEntDto struct {
    /*
        企业id     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailPUserEntDto) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailPUserEntDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailPUserEntDto) SetEntName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailPUserEntDto {
    s.EntName = &v
    return s
}
