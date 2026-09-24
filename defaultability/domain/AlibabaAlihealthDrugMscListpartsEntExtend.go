package domain


type AlibabaAlihealthDrugMscListpartsEntExtend struct {
    /*
        可替换的entId     */
    ReplaceEntId  *string `json:"replace_ent_id,omitempty" `

    /*
        可替换refEntid     */
    ReplaceRefEntId  *string `json:"replace_ref_ent_id,omitempty" `

}

func (s *AlibabaAlihealthDrugMscListpartsEntExtend) SetReplaceEntId(v string) *AlibabaAlihealthDrugMscListpartsEntExtend {
    s.ReplaceEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsEntExtend) SetReplaceRefEntId(v string) *AlibabaAlihealthDrugMscListpartsEntExtend {
    s.ReplaceRefEntId = &v
    return s
}
