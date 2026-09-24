package domain


type AlibabaAlihealthDrugKytWesListpartsEntExtend struct {
    /*
        可替换entid     */
    ReplaceEntId  *string `json:"replace_ent_id,omitempty" `

    /*
        可替换refEntid     */
    ReplaceRefEntId  *string `json:"replace_ref_ent_id,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesListpartsEntExtend) SetReplaceEntId(v string) *AlibabaAlihealthDrugKytWesListpartsEntExtend {
    s.ReplaceEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsEntExtend) SetReplaceRefEntId(v string) *AlibabaAlihealthDrugKytWesListpartsEntExtend {
    s.ReplaceRefEntId = &v
    return s
}
