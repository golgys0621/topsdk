package domain


type AlibabaAlihealthDrugMscListpartByagentEntExtend struct {
    /*
        可替换entId     */
    ReplaceEntId  *string `json:"replace_ent_id,omitempty" `

    /*
        可替换refEntid     */
    ReplaceRefEntId  *string `json:"replace_ref_ent_id,omitempty" `

}

func (s *AlibabaAlihealthDrugMscListpartByagentEntExtend) SetReplaceEntId(v string) *AlibabaAlihealthDrugMscListpartByagentEntExtend {
    s.ReplaceEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentEntExtend) SetReplaceRefEntId(v string) *AlibabaAlihealthDrugMscListpartByagentEntExtend {
    s.ReplaceRefEntId = &v
    return s
}
