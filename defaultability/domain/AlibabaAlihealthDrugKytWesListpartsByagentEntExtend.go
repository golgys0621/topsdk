package domain


type AlibabaAlihealthDrugKytWesListpartsByagentEntExtend struct {
    /*
        可替换entId     */
    ReplaceEntId  *string `json:"replace_ent_id,omitempty" `

    /*
        可替换refEntid     */
    ReplaceRefEntId  *string `json:"replace_ref_ent_id,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesListpartsByagentEntExtend) SetReplaceEntId(v string) *AlibabaAlihealthDrugKytWesListpartsByagentEntExtend {
    s.ReplaceEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentEntExtend) SetReplaceRefEntId(v string) *AlibabaAlihealthDrugKytWesListpartsByagentEntExtend {
    s.ReplaceRefEntId = &v
    return s
}
