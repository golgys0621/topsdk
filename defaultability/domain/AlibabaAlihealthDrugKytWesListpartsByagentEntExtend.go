package domain


type AlibabaAlihealthDrugKytWesListpartsByagentEntExtend struct {
    /*
        可替换entId     */
    ReplaceEntId  *string `json:"replace_ent_id,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesListpartsByagentEntExtend) SetReplaceEntId(v string) *AlibabaAlihealthDrugKytWesListpartsByagentEntExtend {
    s.ReplaceEntId = &v
    return s
}
