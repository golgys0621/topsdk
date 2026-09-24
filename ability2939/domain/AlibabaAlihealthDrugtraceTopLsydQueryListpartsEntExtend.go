package domain


type AlibabaAlihealthDrugtraceTopLsydQueryListpartsEntExtend struct {
    /*
        可替换的entId     */
    ReplaceEntId  *string `json:"replace_ent_id,omitempty" `

    /*
        可替换refEntid     */
    ReplaceRefEntId  *string `json:"replace_ref_ent_id,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsEntExtend) SetReplaceEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsEntExtend {
    s.ReplaceEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsEntExtend) SetReplaceRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsEntExtend {
    s.ReplaceRefEntId = &v
    return s
}
