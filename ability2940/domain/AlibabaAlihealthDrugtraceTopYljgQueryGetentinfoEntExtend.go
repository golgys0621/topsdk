package domain


type AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoEntExtend struct {
    /*
        可替换的entId     */
    ReplaceEntId  *string `json:"replace_ent_id,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoEntExtend) SetReplaceEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoEntExtend {
    s.ReplaceEntId = &v
    return s
}
