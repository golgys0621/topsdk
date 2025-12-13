package domain


type AlibabaAlihealthDrugKytWesListpartsEntExtend struct {
    /*
        可替换entid     */
    ReplaceEntId  *string `json:"replace_ent_id,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesListpartsEntExtend) SetReplaceEntId(v string) *AlibabaAlihealthDrugKytWesListpartsEntExtend {
    s.ReplaceEntId = &v
    return s
}
