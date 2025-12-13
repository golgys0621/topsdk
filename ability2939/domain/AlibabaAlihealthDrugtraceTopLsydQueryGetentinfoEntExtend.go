package domain


type AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoEntExtend struct {
    /*
        可替换的entId     */
    ReplaceEntId  *string `json:"replace_ent_id,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoEntExtend) SetReplaceEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoEntExtend {
    s.ReplaceEntId = &v
    return s
}
