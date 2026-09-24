package domain


type AlibabaAlihealthDrugtraceTopYljgQueryListpartsEntExtend struct {
    /*
        可替换entid     */
    ReplaceEntId  *string `json:"replace_ent_id,omitempty" `

    /*
        可替换refEntid     */
    ReplaceRefEntId  *string `json:"replace_ref_ent_id,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsEntExtend) SetReplaceEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsEntExtend {
    s.ReplaceEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsEntExtend) SetReplaceRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsEntExtend {
    s.ReplaceRefEntId = &v
    return s
}
