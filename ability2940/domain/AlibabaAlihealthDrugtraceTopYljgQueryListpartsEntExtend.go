package domain


type AlibabaAlihealthDrugtraceTopYljgQueryListpartsEntExtend struct {
    /*
        可替换entid     */
    ReplaceEntId  *string `json:"replace_ent_id,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsEntExtend) SetReplaceEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsEntExtend {
    s.ReplaceEntId = &v
    return s
}
